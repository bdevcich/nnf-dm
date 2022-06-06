#!/bin/bash

# Copyright 2021, 2022 Hewlett Packard Enterprise Development LP
# Other additional copyright holders may be indicated within.
#
# The entirety of this work is licensed under the Apache License,
# Version 2.0 (the "License"); you may not use this file except
# in compliance with the License.
#
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


if [ $# -eq 0 ]; then
  echo "Override Config requires one of 'lustre' or 'xfs'"
  exit 1
fi

CMD=$1


if [[ "$CMD" == lustre ]]; then
  echo "$(tput bold)Creating override config $(tput sgr 0)"
  kubectl delete configmap/data-movement-config &> /dev/null
  cat <<-EOF | kubectl create -f -
  apiVersion: v1
  kind: ConfigMap
  metadata:
    name: data-movement-config
    namespace: nnf-dm-system
  data:
    image: arti.dev.cray.com/rabsw-docker-master-local/mfu:0.0.1
    sourceVolume:      '{ "hostPath": { "path": "/nnf", "type": "Directory" } }'
    destinationVolume: '{ "hostPath": { "path": "/nnf", "type": "Directory" } }'
EOF
fi

if [[ "$CMD" == xfs ]]; then
  echo "$(tput bold)Creating override config $(tput sgr 0)"
  kubectl delete configmap/data-movement-config &> /dev/null
  cat <<-EOF | kubectl create -f -
  apiVersion: v1
  kind: ConfigMap
  metadata:
    name: data-movement-config
    namespace: nnf-dm-system
  data:
    sourcePath: "/mnt/file.in"
    destinationPath: "/mnt/file.out"
EOF
fi
