#!/bin/sh

for service in "$@"
do
    echo "Generating TypeScript stubs from protofile for service $service"
    out=/source/$service/sdk
    mkdir $out >/dev/null 2>&1 || true
    cd /source/$service/protobuf
    protoc -I/usr/local/include -I. -I/googleapis -I${GOPATH}/src --plugin=protoc-gen-ts=/usr/bin/protoc-gen-ts --js_out=import_style=commonjs,binary:$out --ts_out=service=true:$out proto.proto
    sed -i'' '/com_gogo_protobuf_gogoproto_gogo_pb/d' $out/proto_pb.d.ts
    sed -i'' '/com_gogo_protobuf_gogoproto_gogo_pb/d' $out/proto_pb_service.ts
    sed -i'' '/com_gogo_protobuf_gogoproto_gogo_pb/d' $out/proto_pb.js

    cp /templates/index.ts $out/index.ts
done