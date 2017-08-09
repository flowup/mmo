package dockercmd

import "github.com/flowup/mmo/environment"

// These constants are commands that are used to generated code from the
// proto files of services
const (
	ProtocBin    = "protoc "
	ProtoInclude = "-I/usr/local/include -I. -I/googleapis -I${GOPATH}/src "
	ProtoDefName = "proto.proto"

	ImageGo      = "flowup/mmo-go-grpc:" + environment.Ver
	ImageTs      = "flowup/mmo-webrpc:" + environment.Ver
	ImagePy      = "flowup/mmo-py-grpc:" + environment.Ver
	ImageSwagger = "flowup/mmo-py-grpc:" + environment.Ver

	GoGen      = ProtocBin + ProtoInclude + "--gogo_out=plugins=grpc:/out " + ProtoDefName
	GGwGen     = ProtocBin + ProtoInclude + "--grpc-gateway_out=logtostderr=true:/out " + ProtoDefName
	SwaggerGen = ProtocBin + ProtoInclude + "--swagger_out=logtostderr=true:/out " + ProtoDefName
	TsGen      = ProtocBin + ProtoInclude + "--plugin=protoc-gen-ts=/usr/bin/protoc-gen-ts --js_out=import_style=commonjs,binary:/out --ts_out=service=true:/out " + ProtoDefName + " && ts-clean "
	PyGen      = "python -m grpc_tools.protoc " + ProtoInclude + "--python_out=/out --grpc_python_out=/out " + ProtoDefName
)
