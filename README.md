# mmo

`mmo` is a cli tool for monorepo microservice orchestration. It helps to easily create, build, test and deploy services within the same repo.

> Tested on Go > 1.9

## Installing from source

MMO supports only installation from source, as there is no support for vendor locked `go get`.

```
go get -u github.com/flowup/mmo # clone the repository to your tree
cd $GOPATH/src/github.com/flowup/mmo
dep ensure
go install
```

## Templates
Command for project and service are using templates for generation. Default templates can be shown using
```bash
$ mmo init --help
$ mmo add service --help
```

Both of these two commands take parameter -t to change template and parameter -x to pass additional options to template. List of passable options can be listed using:
```bash
$ mmo template [name]
```

## Commands

### Help

To see all commands available by the `mmo` CLI, run:
```
$ mmo help
```

### Project creation

Creating new project can be done using `mmo init` command. This will automatically create a new mmo project. There are multiple flags that can be overriden when using the `mmo init` command. See `mmo init -h` to list them.

```
$ mmo init projectname
```

Every project is by default created with Wercker CI configuration, Kubernetes deployment configs and Dep vendoring. In case you want to override those, see command flags.

### Service creation

Microservice creation is handled by `mmo add service` command. This automatically creates a new folder for the service, registers it to the project and creates proto files, main, infrastructure and deployment configs.

```
$ mmo add service myservice
```
### Plugins
MMO supports plugin system. Each plugin is single purpose. There are only generation plugins for now and they can be run using command `mmo gen`. Used plugins can be defined in mmo.yaml in plugins section (project plugins), or in plugins section in service (service plugins):
```
...
plugins:
 - some-plugin:latest
 - another-plugin:0.1
...
```

#### Supported plugins

Service:

* `flowup/mmo-gen-go-grpc-desc:latest` - Plugin generates gRPC Go API stub and server from protofile `/service/protobuf/proto.proto`, output is saved to /service/proto.pb.go, service descriptor for API endpoints is also generated
* `flowup/mmo-gen-grpc-gateway:0.2` - Plugin generates rest API gateway from protofile `/service/protobuf/proto.proto`, output is saved to /service/proto.pb.gw.go

Project:

* `mmo-gen-swagger:apiendpoints` - Plugin generates swagger definition from protofile `/service/protobuf/proto.proto`, output is saved to /service/swagger.json
* `mmo-gen-angular-client:1.1`

## Example usage
* Init project using `mmo init myproject`, then `cd myproject`
* Create service with grpc gateway using `mmo add service -x Gateway=true myservice`
    * All template options are available using `mmo template [template name]`
* Add plugins `flowup/mmo-gen-go-grpc-desc:latest` and `flowup/mmo-gen-grpc-gateway:0.2` to `myservice` plugins scope of the `mmo.yaml` manifest
* Implement `exampleservice/protobuf/proto.proto`
* Lock and install dependencies using `dep ensure`
* Run `mmo gen` to generate API stubs, servers and gateway
* Implement `exampleservice/service.go` according to new generated API server
