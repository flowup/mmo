#!/bin/sh

for service in "$@"
do
    echo "Generating gRPC server and stubs from protofile for service $service"
    cd /source/$service/protobuf
    protoc -I/usr/local/include -I. -I/googleapis -I${GOPATH}/src --grpc-gateway_out=logtostderr=true:/source/$service proto.proto
done