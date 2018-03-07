package server

import (
	"context"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/flowup/mmo/api"
	"github.com/flowup/mmo/config"
	"github.com/golang/protobuf/proto"
	"github.com/google/go-github/github"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	grpcPort = ":50051"
	httpPort = ":50080"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func Serve() {

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	mmoConfig, err := config.LoadConfig(config.FilenameConfig)
	if err != nil {
		log.Fatalf("Failed to load mmo config")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	githubClient := github.NewClient(tc)

	// create the grpc server
	s := grpc.NewServer()

	// register the service
	api.RegisterApiServiceServer(s, api.NewAPIService(mmoConfig, githubClient))

	// Register reflection service on gRPC server.
	reflection.Register(s)

	// if any service fails, whole app should fail
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		log.Infoln("Starting gRPC server on", grpcPort)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("grpc: failed to serve: %v", err)
		}

		wg.Done()
	}()

	go func() {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		mux := runtime.NewServeMux(runtime.WithForwardResponseOption(corsFilter))
		opts := []grpc.DialOption{grpc.WithInsecure()}
		err := api.RegisterApiServiceHandlerFromEndpoint(ctx, mux, grpcPort, opts)
		if err != nil {
			log.Fatalf("gw: failed to register: %v", err)
		}

		log.Infoln("Starting gateway server on", httpPort)
		log.Fatalf("gw: failed to server: %v", http.ListenAndServe(httpPort, mux))
	}()

	wg.Wait()
}

func corsFilter(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return nil
}
