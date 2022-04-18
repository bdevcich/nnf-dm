# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import rsyncdatamovement_pb2 as rsyncdatamovement__pb2


class RsyncDataMoverStub(object):
    """RsyncDataMover service definition describes the API for perform data movement
    for NNF storage. 
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Create = channel.unary_unary(
                '/datamovement.RsyncDataMover/Create',
                request_serializer=rsyncdatamovement__pb2.RsyncDataMovementCreateRequest.SerializeToString,
                response_deserializer=rsyncdatamovement__pb2.RsyncDataMovementCreateResponse.FromString,
                )
        self.Status = channel.unary_unary(
                '/datamovement.RsyncDataMover/Status',
                request_serializer=rsyncdatamovement__pb2.RsyncDataMovementStatusRequest.SerializeToString,
                response_deserializer=rsyncdatamovement__pb2.RsyncDataMovementStatusResponse.FromString,
                )
        self.Delete = channel.unary_unary(
                '/datamovement.RsyncDataMover/Delete',
                request_serializer=rsyncdatamovement__pb2.RsyncDataMovementDeleteRequest.SerializeToString,
                response_deserializer=rsyncdatamovement__pb2.RsyncDataMovementDeleteResponse.FromString,
                )


class RsyncDataMoverServicer(object):
    """RsyncDataMover service definition describes the API for perform data movement
    for NNF storage. 
    """

    def Create(self, request, context):
        """Create sends a new data movement request identified by source and destination fields. It returns
        a response containing a unique identifier which can be used to query the status of the request.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Status(self, request, context):
        """Status requests the status of a previously submitted data movement request. It accepts a unique
        identifier that identifies the request and returns a status message.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Delete(self, request, context):
        """Delete will attempt to delete a completed data movement request. It accepts a unique identifer
        that identifies the request and returns the status of the delete operation.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_RsyncDataMoverServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Create': grpc.unary_unary_rpc_method_handler(
                    servicer.Create,
                    request_deserializer=rsyncdatamovement__pb2.RsyncDataMovementCreateRequest.FromString,
                    response_serializer=rsyncdatamovement__pb2.RsyncDataMovementCreateResponse.SerializeToString,
            ),
            'Status': grpc.unary_unary_rpc_method_handler(
                    servicer.Status,
                    request_deserializer=rsyncdatamovement__pb2.RsyncDataMovementStatusRequest.FromString,
                    response_serializer=rsyncdatamovement__pb2.RsyncDataMovementStatusResponse.SerializeToString,
            ),
            'Delete': grpc.unary_unary_rpc_method_handler(
                    servicer.Delete,
                    request_deserializer=rsyncdatamovement__pb2.RsyncDataMovementDeleteRequest.FromString,
                    response_serializer=rsyncdatamovement__pb2.RsyncDataMovementDeleteResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'datamovement.RsyncDataMover', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class RsyncDataMover(object):
    """RsyncDataMover service definition describes the API for perform data movement
    for NNF storage. 
    """

    @staticmethod
    def Create(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/datamovement.RsyncDataMover/Create',
            rsyncdatamovement__pb2.RsyncDataMovementCreateRequest.SerializeToString,
            rsyncdatamovement__pb2.RsyncDataMovementCreateResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Status(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/datamovement.RsyncDataMover/Status',
            rsyncdatamovement__pb2.RsyncDataMovementStatusRequest.SerializeToString,
            rsyncdatamovement__pb2.RsyncDataMovementStatusResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Delete(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/datamovement.RsyncDataMover/Delete',
            rsyncdatamovement__pb2.RsyncDataMovementDeleteRequest.SerializeToString,
            rsyncdatamovement__pb2.RsyncDataMovementDeleteResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)