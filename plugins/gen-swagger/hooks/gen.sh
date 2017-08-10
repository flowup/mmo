#!/bin/sh

cp /plugin/templates/swagger.yaml /source/swagger.yaml

for service in "$@"
do
    echo "Generating swagger definition for service $service"
    cd /source/$service/protobuf
    mkdir /tmp/$service
    protoc -I/usr/local/include -I. -I/googleapis -I${GOPATH}/src --swagger_out=logtostderr=true:/source/$service proto.proto
    #swagger-merger -i /tmp/$service.swagger.json -o /source/swagger.yaml
done