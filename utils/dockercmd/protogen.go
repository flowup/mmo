package dockercmd

var (
	ProtocBin    = "protoc "
	ProtoInclude = "-I/usr/local/include -I. -I/googleapis -I${GOPATH}/src "
	ProtoDefName = "proto.proto"
	GoGen        = ProtocBin + ProtoInclude + "--gogo_out=plugins=grpc:/out " + ProtoDefName
	TsGen        = ProtocBin + ProtoInclude + "--plugin=protoc-gen-ts=/usr/bin/protoc-gen-ts --js_out=import_style=commonjs,binary:/out --ts_out=service=true:/out " + ProtoDefName + " && removeGogoImports "
)
