package main

import (
	"net"
	"sync"
	"google.golang.org/grpc"
	"net/http"
	"github.com/flowup/mmo/{{.ProjectName}}/{{.Name}}"
	"google.golang.org/grpc/reflection"
{{if .WebGrpc}}	"github.com/improbable-eng/grpc-web/go/grpcweb"{{end}}
)

func main(){
	lis, err := net.Listen("tcp", "")
	if err != nil {
		{{.Name}}.Log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// Register reflection service on gRPC server.
	reflection.Register(s)

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

	go func() {
		{{.Name}}.Log.Println("Starting gRPC server on", "")
		if err := s.Serve(lis); err != nil {
			{{.Name}}.Log.Fatalf("grpc: failed to serve: %v", err)
		}

		wg.Done()
	}()

	{{if .WebGrpc}}go func() {
		{{.Name}}.Log.Println("starting gRPC Web server on", "")
		if err := httpServer.ListenAndServe(); err != nil {
			{{.Name}}.Log.Fatalf("grpc-web: failed to serve: %v", err)
		}

		wg.Done()
	}(){{end}}

	wg.Wait()
}
