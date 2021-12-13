import crank_pb2
import crank_pb2_grpc


class Crank(crank_pb2_grpc.CrankServicer):
    def CreateCrank(self, create_crank_request, context):
        crank = crank_pb2.BikeCrank(
            id='urn:bicycle:crank:1',
            frame_id=create_crank_request.frame_id,
            manufacturer=create_crank_request.manufacturer,
            arm_length=create_crank_request.arm_length,
            model=create_crank_request.model
        )
        return crank_pb2.CreateCrankResponse(
            crank=crank
        )
