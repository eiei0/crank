from concurrent import futures
import logging
import grpc
import os

import crank_pb2_grpc
import server as crank_server


def serve():
    host = os.getenv('GRPC_SERVER_HOST', '0.0.0.0')
    port = os.getenv('GRPC_SERVER_PORT', 8092)

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    crank_pb2_grpc.add_CrankServicer_to_server(crank_server.Crank(), server)

    server.add_insecure_port("%s:%s" % (host, port))
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
