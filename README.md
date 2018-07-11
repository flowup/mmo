# MMO

Monorepo Microservice Orchestration is a CLI tool that helps to easily create, build, test and deploy services within the same repository.

## Installation

MMO supports installation from source, as there is no support for vendor locked `go get` at the moment.

> Go > 1.9

```
go get -u github.com/flowup/mmo # clone the repository to your tree (this may fail to install, just continue)
cd $GOPATH/src/github.com/flowup/mmo
dep ensure
go install
```

## Language support

- [x] Golang
- [ ] Python (templates and plugins will be available until RC release)

## Project creation

> Default initialization template can be seen within `mmo init --help`

To create a new project called `myproject`, optional parameter `-p` to change path to project, run:
```bash
$ mmo init myproject
```

Use parameter `-t` to change template and parameter `-x` to pass additional options to template. List of options can be seen using:
```bash
$ mmo template [name]
```

## Scaffolding
MMO provides simple boilerplate generator to easily add components such as microservices, models and plugins via its `mmo add` command.

```
mmo add --help
```

To create a new microservice called `book` within your project, simply run:
```
mmo add service book
```

## Plugin ecosystem

MMO is fully pluggable via Docker images. Officialy supported plugins can be found in this repo https://github.com/flowup/mmo-plugins

- [x] gRPC service plugin (protobuf)
- [x] gRPC HTTP gateway plugin
- [x] Swagger plugin
- [x] Angular API client plugin
- [x] ...

All plugins can be launched via:
```
mmo gen
```

## Launching plugins for a group of services

This functionality is supported to avoid long plugin running times for projects with a large number of services.

This command will run all plugins only for services `book` and `auth` within your project. All other services will be ignored.
```
mmo context book auth
mmo gen

mmo context --reset
```

## Example usage
* Create project using `mmo init myproject`, then `cd myproject`
* Create service with grpc gateway using `mmo add service -x Gateway=true myservice`
    * All template options are available using `mmo template [template name]`
* Add plugins `flowup/mmo-gen-go-grpc-desc:latest` and `flowup/mmo-gen-grpc-gateway:0.2` to `myservice` plugins scope of the `mmo.yaml` manifest
* Implement `exampleservice/protobuf/proto.proto`
* Lock and install dependencies using `dep ensure`
* Run `mmo gen` to generate API stubs, servers and gateway
* Implement `exampleservice/service.go` according to new generated API server
