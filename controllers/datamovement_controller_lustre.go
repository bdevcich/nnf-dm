package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	kubeflowv1 "github.com/kubeflow/common/pkg/apis/common/v1"
	mpiv2beta1 "github.com/kubeflow/mpi-operator/v2/pkg/apis/kubeflow/v2beta1"

	dmv1alpha1 "github.hpe.com/hpe/hpc-rabsw-nnf-dm/api/v1alpha1"
	nnfv1alpha1 "github.hpe.com/hpe/hpc-rabsw-nnf-sos/api/v1alpha1"

	lustrecsi "github.hpe.com/hpe/hpc-rabsw-lustre-csi-driver/pkg/lustre-driver/service"
	lustrectrl "github.hpe.com/hpe/hpc-rabsw-lustre-fs-operator/controllers"
)

const (
	persistentVolumeSuffix      = "-pv"
	persistentVolumeClaimSuffix = "-pvc"
	mpiJobSuffix                = "-mpi"
	configSuffix                = "-config"
)

const (
	configImage             = "image"             // Image specifies the image used in the MPI launcher & worker containers
	configCommand           = "command"           // Command specifies the command to run. Defaults to "mpirun"
	configSourcePath        = "sourcePath"        // SourcePath is the path of the source file or directory
	configDestinationPath   = "destinationPath"   // DestinationPath is the path of the destination file or directory
	configSourceVolume      = "sourceVolume"      // SourceVolume is the corev1.VolumeSource used as the source volume mount. Defaults to a CSI volume interpreted from the Spec.Source
	configDestinationVolume = "destinationVolume" // DestinationVolume is the corev1.VolumeSource used as the destination volume mount. Defaults to a CSI volume interpreted from the Spec.Destination
)

func (r *DataMovementReconciler) initializeLustreJob(ctx context.Context, dm *dmv1alpha1.DataMovement) (*ctrl.Result, error) {
	log := log.FromContext(ctx, "DataMovement", "Lustre")

	// We need to label all the nodes in the Servers object with a unique label that describes this
	// data movememnt. This label is then used as a selector within the the MPIJob so it correctly
	// targets all the nodes
	result, workerCount, err := r.labelStorageNodes(ctx, dm)
	if err != nil {
		log.Error(err, "Failed to label storage nodes")
		return nil, err
	} else if !result.IsZero() {
		return result, nil
	}

	// TODO: The PV/PVC should only be needed for JobStorageInstance, otherwise the PV/PVC will already
	//       be created, and we only need to reference them in the MPIJob.

	if err := r.createPersistentVolume(ctx, dm); err != nil {
		log.Error(err, "Failed to create persistent volume")
		return nil, err
	}

	if err := r.createPersistentVolumeClaim(ctx, dm); err != nil {
		log.Error(err, "Failed to create persistent volume claim")
		return nil, err
	}

	if err := r.createMpiJob(ctx, dm, workerCount); err != nil {
		log.Error(err, "Failed to create MPI Job")
		return nil, err
	}

	log.Info("Lustre data movement initialized")
	return nil, nil
}

func (r *DataMovementReconciler) labelStorageNodes(ctx context.Context, dm *dmv1alpha1.DataMovement) (*ctrl.Result, int32, error) {
	log := log.FromContext(ctx).WithName("label")

	// List of target node names that are to perform lustre data movement
	targetNodeNames := make([]string, 0)

	switch dm.Spec.Storage.Kind {
	case "NnfStorage":
		storageRef := dm.Spec.Storage
		storage := &nnfv1alpha1.NnfStorage{}
		if err := r.Get(ctx, types.NamespacedName{Name: storageRef.Name, Namespace: storageRef.Namespace}, storage); err != nil {
			return nil, -1, err
		}

		targetAllocationSetIndex := -1
		for allocationSetIndex, allocationSet := range storage.Spec.AllocationSets {
			if allocationSet.TargetType == "OST" {
				targetAllocationSetIndex = allocationSetIndex
			}
		}

		if targetAllocationSetIndex == -1 {
			return nil, -1, fmt.Errorf("OST allocation set not found")
		}

		for _, storageNode := range storage.Spec.AllocationSets[targetAllocationSetIndex].Nodes {
			targetNodeNames = append(targetNodeNames, storageNode.Name)
		}

		break
	default:
		panic(fmt.Sprintf("Unsupported storage type: %s", dm.Spec.Storage.Kind))
	}

	// Retrieve all the NNF Nodes in the cluster - these nodes will be matched against the requested
	// node list and labeled such that the mpijob can target the desired nodes
	nodes := &corev1.NodeList{}
	if err := r.List(ctx, nodes, client.HasLabels{"cray.nnf.node"}); err != nil {
		return nil, -1, err
	}

	label := dm.Name
	log.V(2).Info("Labeling nodes", "label", label, "count", len(targetNodeNames))

	for _, nodeName := range targetNodeNames {

		nodeFound := false
		for _, node := range nodes.Items {

			if node.Name == nodeName {
				if _, found := node.Labels[label]; !found {

					node.Labels[label] = "true"

					log.V(1).Info("Applying label to node", "node", nodeName)
					if err := r.Update(ctx, &node); err != nil {
						if errors.IsConflict(err) {
							return &ctrl.Result{Requeue: true}, -1, nil
						}

						return nil, -1, err
					}
				}

				nodeFound = true
				break
			}
		}

		if !nodeFound {
			log.V(3).Info("Node not found. Check the spelling or status of the node.", "node", nodeName)
		}
	}

	return nil, int32(len(targetNodeNames)), nil
}

func (r *DataMovementReconciler) teardownLustreJob(ctx context.Context, dm *dmv1alpha1.DataMovement) (*ctrl.Result, error) {
	log := log.FromContext(ctx).WithName("unlabel")

	label := dm.Name
	nodes := &corev1.NodeList{}
	if err := r.List(ctx, nodes, client.HasLabels{label}); err != nil {
		return nil, err
	}

	log.V(2).Info("Unlabelling nodes", "count", len(nodes.Items))
	for _, node := range nodes.Items {
		delete(node.Labels, label)
		if err := r.Update(ctx, &node); err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (r *DataMovementReconciler) createPersistentVolume(ctx context.Context, dm *dmv1alpha1.DataMovement) error {
	log := log.FromContext(ctx).WithName("pv")

	if dm.Spec.Storage.Kind != "NnfStorage" {
		panic("Create Persistent Volume requires NNF Storage type")
	}

	storage := &nnfv1alpha1.NnfStorage{}
	if err := r.Get(ctx, types.NamespacedName{Name: dm.Spec.Storage.Name, Namespace: dm.Spec.Storage.Namespace}, storage); err != nil {
		return err
	}

	pv := &corev1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name:      dm.Name + persistentVolumeSuffix,
			Namespace: dm.Namespace,
		},
	}

	result, err := ctrl.CreateOrUpdate(ctx, r.Client, pv, func() error {

		volumeMode := corev1.PersistentVolumeFilesystem
		pv.Spec = corev1.PersistentVolumeSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{
				corev1.ReadWriteMany,
			},
			Capacity: corev1.ResourceList{
				corev1.ResourceStorage: resource.MustParse("1"),
			},
			PersistentVolumeSource: corev1.PersistentVolumeSource{
				CSI: &corev1.CSIPersistentVolumeSource{
					Driver:       lustrecsi.Name,
					FSType:       "lustre",
					VolumeHandle: storage.Status.MgsNode,
				},
			},
			VolumeMode: &volumeMode,
		}

		return ctrl.SetControllerReference(dm, pv, r.Scheme)
	})

	if err != nil {
		log.Error(err, "Failed to create persistent volume")
		return err
	} else if result == controllerutil.OperationResultCreated {
		log.V(2).Info("Created persistent volume", "object", client.ObjectKeyFromObject(pv).String())
	}

	return nil
}

func (r *DataMovementReconciler) createPersistentVolumeClaim(ctx context.Context, dm *dmv1alpha1.DataMovement) error {
	log := log.FromContext(ctx).WithName("pvc")

	pvc := &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      dm.Name + persistentVolumeClaimSuffix,
			Namespace: dm.Namespace,
		},
	}

	result, err := ctrl.CreateOrUpdate(ctx, r.Client, pvc, func() error {
		pvc.Spec = corev1.PersistentVolumeClaimSpec{
			VolumeName: dm.Name + persistentVolumeSuffix,
			AccessModes: []corev1.PersistentVolumeAccessMode{
				corev1.ReadWriteMany,
			},
			Resources: corev1.ResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: resource.MustParse("1"),
				},
			},
		}

		return ctrl.SetControllerReference(dm, pvc, r.Scheme)
	})

	if err != nil {
		log.Error(err, "Failed to create persistent volume claim")
		return err
	} else if result == controllerutil.OperationResultCreated {
		log.V(2).Info("Created persistent volume claim", "object", client.ObjectKeyFromObject(pvc).String())
	}

	return nil
}

func (r *DataMovementReconciler) createMpiJob(ctx context.Context, dm *dmv1alpha1.DataMovement, workerCount int32) error {
	log := log.FromContext(ctx)

	config, err := r.getDataMovementConfigMap(ctx)
	if err != nil {
		log.Error(err, "Failed to read mpi configuration")
		return err
	}

	image := "arti.dev.cray.com/rabsw-docker-master-local/mfu:0.0.1"
	if img, found := config.Data[configImage]; found {
		image = img
	}

	command := []string{"mpirun", "--allow-run-as-root", "dcp", "/mnt/src" + dm.Spec.Source.Path, "/mnt/dest" + dm.Spec.Destination.Path}
	if cmd, found := config.Data[configCommand]; found {
		if strings.HasPrefix(cmd, "/bin/bash -c") {
			command = []string{"/bin/bash", "-c", strings.TrimPrefix(cmd, "/bin/bash -c")}
		} else {
			command = strings.Split(cmd, " ")
		}
		log.V(1).Info("Command override", "command", command)
	}

	launcher := &kubeflowv1.ReplicaSpec{
		Template: corev1.PodTemplateSpec{
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Image:   image,
						Name:    dm.Name,
						Command: command,
					},
				},
			},
		},
	}

	replicas := workerCount
	worker := &kubeflowv1.ReplicaSpec{
		Replicas: &replicas,
		Template: corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					dm.Name: "true",
				},
			},
			Spec: corev1.PodSpec{
				NodeSelector: map[string]string{
					dm.Name: "true",
				},
				Tolerations: []corev1.Toleration{
					{
						Key:      "cray.nnf.node",
						Operator: corev1.TolerationOpEqual,
						Value:    "true",
						Effect:   "NoSchedule",
					},
				},
				Affinity: &corev1.Affinity{
					// Prevent multiple mpi-workers from being scheduled on the same node. That is to say...
					// The pod should _not_ be scheduled (through anti-affinity) onto a node if that node
					// is in the same zone (dm.Name) as a pod having label dm.Name="true".
					PodAntiAffinity: &corev1.PodAntiAffinity{
						PreferredDuringSchedulingIgnoredDuringExecution: []corev1.WeightedPodAffinityTerm{
							{
								Weight: 100,
								PodAffinityTerm: corev1.PodAffinityTerm{
									LabelSelector: &metav1.LabelSelector{
										MatchLabels: map[string]string{
											dm.Name: "true",
										},
									},
									TopologyKey: dm.Name,
								},
							},
						},
					},
				},
				Containers: []corev1.Container{
					{
						Image: config.Data[configImage],
						Name:  dm.Name,
						VolumeMounts: []corev1.VolumeMount{
							{
								Name:      "source",
								MountPath: "/mnt/src",
							},
							{
								Name:      "destination",
								MountPath: "/mnt/dest",
							},
						},
					},
				},
				Volumes: []corev1.Volume{
					{
						Name:         "source",
						VolumeSource: r.getVolumeSource(ctx, dm, config, configSourceVolume, r.getLustreSourcePersistentVolumeClaimName),
					},
					{
						Name:         "destination",
						VolumeSource: r.getVolumeSource(ctx, dm, config, configDestinationVolume, r.getLustreDestinationPersistentVolumeClaimName),
					},
				},
			},
		},
	}

	job := &mpiv2beta1.MPIJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      dm.Name + mpiJobSuffix,
			Namespace: dm.Namespace,
		},
		Spec: mpiv2beta1.MPIJobSpec{
			MPIReplicaSpecs: map[mpiv2beta1.MPIReplicaType]*kubeflowv1.ReplicaSpec{
				mpiv2beta1.MPIReplicaTypeLauncher: launcher,
				mpiv2beta1.MPIReplicaTypeWorker:   worker,
			},
		},
	}

	ctrl.SetControllerReference(dm, job, r.Scheme)

	log.V(2).Info("Creating mpi job", "name", client.ObjectKeyFromObject(job).String())
	return r.Create(ctx, job)
}

func (r *DataMovementReconciler) getVolumeSource(ctx context.Context, dm *dmv1alpha1.DataMovement, config *corev1.ConfigMap, override string, claimFn func(context.Context, *dmv1alpha1.DataMovement) string) corev1.VolumeSource {
	if data, found := config.Data[override]; found {
		source := corev1.VolumeSource{}
		if err := json.Unmarshal([]byte(data), &source); err == nil {
			return source
		} else {
			log.FromContext(ctx).Info("Failed to unmarshal override config " + override)
		}
	}

	return corev1.VolumeSource{
		PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
			ClaimName: claimFn(ctx, dm),
		},
	}
}

func (r *DataMovementReconciler) getLustreSourcePersistentVolumeClaimName(ctx context.Context, dm *dmv1alpha1.DataMovement) string {
	ref := dm.Spec.Source.StorageInstance

	if ref.Kind == "LustreFileSystem" {
		return ref.Name + lustrectrl.PersistentVolumeClaimSuffix
	} else if ref.Kind == "NnfPersistentStorageInstance" {
		return ref.Name + "TODO" // TODO: Should look this up via nnf-sos when creating persistent lustre volume
	} else if ref.Kind == "NnfJobStorageInstance" {
		return dm.Name + persistentVolumeClaimSuffix
	}

	panic("Unsupported Lustre Source PVC: " + ref.Kind)
}

func (r *DataMovementReconciler) getLustreDestinationPersistentVolumeClaimName(ctx context.Context, dm *dmv1alpha1.DataMovement) string {
	ref := dm.Spec.Destination.StorageInstance

	if ref.Kind == "LustreFileSystem" {
		return ref.Name + lustrectrl.PersistentVolumeClaimSuffix
	} else if ref.Kind == "NnfPersistentStorageInstance" {
		return ref.Name + "TODO" // TODO: Should look this up via nnf-sos when creating persistent lustre volume
	} else if ref.Kind == "NnfJobStorageInstance" {
		return dm.Name + persistentVolumeClaimSuffix
	}

	panic("Unsupported Lustre Destination PVC: " + ref.Kind)
}

func (r *DataMovementReconciler) monitorLustreJob(ctx context.Context, dm *dmv1alpha1.DataMovement) (*ctrl.Result, string, string, error) {

	job := &mpiv2beta1.MPIJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      dm.Name + mpiJobSuffix,
			Namespace: dm.Namespace,
		},
	}

	if err := r.Get(ctx, client.ObjectKeyFromObject(job), job); err != nil {
		return nil, dmv1alpha1.DataMovementConditionReasonFailed, "ObjectNotFound", err
	}

	for _, condition := range job.Status.Conditions {
		if condition.Type == kubeflowv1.JobFailed {
			return nil, dmv1alpha1.DataMovementConditionReasonFailed, condition.Message, nil
		} else if condition.Type == kubeflowv1.JobSucceeded {
			return nil, dmv1alpha1.DataMovementConditionReasonSuccess, condition.Message, nil
		}
	}

	// Consider the job still running
	return &ctrl.Result{}, dmv1alpha1.DataMovementConditionTypeRunning, "Running", nil
}
