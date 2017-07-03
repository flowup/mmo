#!/bin/sh

sed -i'' '/com_gogo_protobuf_gogoproto_gogo_pb/d' /out/proto_pb.d.ts
sed -i'' '/com_gogo_protobuf_gogoproto_gogo_pb/d' /out/proto_pb_service.ts
sed -i'' '/com_gogo_protobuf_gogoproto_gogo_pb/d' /out/proto_pb.js
