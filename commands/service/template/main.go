package main

import (
	"sync"
	"google.golang.org/grpc"
	"net/http"
	"github.com/flowup/mmo/{{.ProjectName}}/{{.Name}}"
{{if .WebGrpc}}	"github.com/improbable-eng/grpc-web/go/grpcweb"{{end}}
)

func main(){
	s := grpc.NewServer()

	{{if .WebGrpc}}options := []grpcweb.Option{}

	// Wrap grpc server to grpc-web server
	wrappedServer := grpcweb.WrapServer(s, options...)

	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(resp, req)
	}

	httpServer := http.Server{
		Addr:    "",
		Handler: http.HandlerFunc(handler),
	}
	{{end}}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	{{if .WebGrpc}}
	go func() {
		{{.Name}}.Log.Println("starting gRPC Web server on", "")
		if err := httpServer.ListenAndServe(); err != nil {
			{{.Name}}.Log.Fatalf("grpc-web: failed to serve: %v", err)
		}

		wg.Done()
	}(){{end}}

	wg.Wait()
}
