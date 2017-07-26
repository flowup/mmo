package dockercmd

// These constants are commands that are used to generated code from the
// proto files of services
const (
	ProtocBin    = "protoc "
	ProtoInclude = "-I/usr/local/include -I. -I/googleapis -I${GOPATH}/src "
	ProtoDefName = "proto.proto"

	ImageGo = "flowup/mmo-go-grpc"
	ImageTs = "flowup/mmo-webrpc"
	ImagePy = "flowup/mmo-py-grpc"

	GoGen  = ProtocBin + ProtoInclude + "--gogo_out=plugins=grpc:/out " + ProtoDefName
	GGwGen = ProtocBin + ProtoInclude + "--grpc-gateway_out=logtostderr=true:/out " + ProtoDefName
	TsGen  = ProtocBin + ProtoInclude + "--plugin=protoc-gen-ts=/usr/bin/protoc-gen-ts --js_out=import_style=commonjs,binary:/out --ts_out=service=true:/out " + ProtoDefName + " && ts-clean "
	PyGen  = "python -m grpc_tools.protoc " + ProtoInclude + "--python_out=/out --grpc_python_out=/out " + ProtoDefName
)
