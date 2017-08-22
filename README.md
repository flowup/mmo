# mmo

`mmo` is a cli tool for monorepo microservice orchestration. It helps to easily create, build, test and deploy services within the same repo.

> Tested on Go > 1.8

## Installing from source

MMO supports only installation from source, as there is no support for vendor locked `go get`.

```
git clone https://github.com/flowup/mmo.git # clone the repository to your tree
cd mmo
glide i # you need to have glide installed curl https://glide.sh/get | sh
go install . # install actual mmo binary
```

## Commands

### Help

To see all commands available by the `mmo` CLI, run:
```
mmo help
```

### Project creation

Creating new project can be done using `mmo init` command. This will automatically create a new mmo project. There are multiple flags that can be overriden when using the `mmo init` command. See `mmo init -h` to list them.

```
mmo init projectname
```

Every project is by default created with Wercker CI configuration, Kubernetes deployment configs and Glide vendoring. In case you want to override those, see command flags.

### Service creation

Microservice creation is handled by `mmo add service` command. This automatically creates a new folder for the service, registers it to the project and creates proto files, main, infrastructure and deployment configs.

```
mmo add service myservice
```
### Plugins
MMO supports plugin system. Each plugin is single purpose. There are only generation plugins for now and they can be run using command `mmo gen`. Used plugins can be defined in mmo.yaml in plugins section:
```
...
plugins:
 - some-plugin:latest
 - another-plugin:0.1
...
```

#### Supported plugins

* `flowup/mmo-gen-go-grpc` - Plugin generates gRPC Go API stub and server from protofile `/service/protobuf/proto.proto`, output is saved to /service/proto.pb.go
* `flowup/mmo-gen-grpc-gateway` - Plugin generates rest API gateway from protofile `/service/protobuf/proto.proto`, output is saved to /service/proto.pb.gw.go
* `flowup/mmo-gen-swagger` - Plugin generates swagger definition from protofile `/service/protobuf/proto.proto`, output is saved to /service/swagger.json
* `flowup/mmo-plugin-godeler`

## Example usage
* Init project using `mmo init myproject`, then `cd myproject`
* Create service with grpc gateway using `mmo add service --gateway exampleservice`
* Add plugins `flowup/mmo-gen-go-grpc` and `flowup/mmo-gen-grpc-gateway` to `mmo.yaml` manifest
* Implement `exampleservice/protobuf/proto.proto`
* Init glide using `glide up` - later use only `glide install` and `glide get`
* Run `mmo gen` to generate API stubs, servers and gateway
* Implement `exampleservice/service.go` according to new generated API server