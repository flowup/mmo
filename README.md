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

Microservice creation is handled by `mmo add service` command. This automatically creates a new folder for the service, registers it to the project and creates proto files, main, infrastructure and deployment configs. Each service also has `docs` folder which is supported by `Godeler - tool for ER -> code transformations`

```
mmo add service myservice
```

### Adding plugins

mmo has native support for so called `plugins`. Plugin is an image with deployment configuration that should be deployed every time with the service (a.k.a. dependency). See the [list of available plugins](https://github.com/flowup/mmo/wiki/Plugins) on what you can install.

```
mmo add plugin postgres
```

## Development

Each generator command is using `go-bindata` tool to embed the templates within the code. To install it, run:
```
go get -u github.com/jteeuwen/go-bindata/...
```
