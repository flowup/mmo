test:
	goconvey -timeout 30s -host 0.0.0.0

bindata:
	go-bindata -pkg project -o commands/project/bindata.go commands/project/template
