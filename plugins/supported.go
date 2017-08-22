package plugins

import "github.com/flowup/mmo/environment"

const (
	Grpc_go      = "flowup/mmo-gen-go-grpc:" + environment.Ver
	Grpc_gw      = "flowup/mmo-gen-grpc-gateway:" + environment.Ver
	Grpc_swagger = "flowup/mmo-gen-swagger:" + environment.Ver
	Grpc_ts      = "flowup/mmo-gen-ts-grpc:" + environment.Ver
)
