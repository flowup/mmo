#!/bin/sh

for service in "$@"
do
    echo "Generating access from draw.io xml for service $service"
    godeler --parse $service/proto.pb.go --template /go/src/github.com/flowup/godeler/templates_store/go/daoMutators.tmpl --pkg $service -o $service/access.gen.go
done