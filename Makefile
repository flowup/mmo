test:
	goconvey -timeout 30s -host 0.0.0.0

bindata:
	go-bindata -pkg project -o commands/project/bindata.go -prefix commands/project/template/ -ignore bindata commands/project/template/...
	go-bindata -pkg service -o commands/service/bindata.go -prefix commands/service/template/ -ignore bindata commands/service/template/...

build:
	go build