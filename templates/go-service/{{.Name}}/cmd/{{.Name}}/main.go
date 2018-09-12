package main

import (
    "os"
	"net"
	"sync"
	"google.golang.org/grpc"
    "github.com/spf13/viper"
    "{{ .Package }}/{{ .Name }}"
	"google.golang.org/grpc/reflection"	{{if .WebRPC}}
	"github.com/improbable-eng/grpc-web/go/grpcweb"
    "net/http"{{end}}
    log "github.com/sirupsen/logrus"{{if .Sentry}}
    "github.com/evalphobia/logrus_sentry"{{end}}{{if .Profiler}}
	"cloud.google.com/go/profiler"{{end}}{{ if .Tracing }}
	"go.opencensus.io/trace"
	"contrib.go.opencensus.io/exporter/stackdriver" {{end}}
	"net/http"
    "context"
    "github.com/grpc-ecosystem/grpc-gateway/runtime"
)

func init() {
    log.SetFormatter(&log.JSONFormatter{})
    log.SetOutput(os.Stdout)
    log.SetLevel(log.DebugLevel)
}

func main(){

    viper.SetDefault("db.conn", "host=db-svc user=goo dbname=goo sslmode=disable password=goo")
	viper.SetDefault("server.binds.grpc", ":50051"){{if .WebRPC}}
	viper.SetDefault("server.binds.webrpc", ":50060"){{end}}
	viper.SetDefault("server.binds.gw", ":80"){{if .Profiler}}
	viper.SetDefault("profiler", false)
	viper.BindEnv("profiler", "PROFILER"){{end}}{{if .Tracing}}
	viper.SetDefault("tracing", false)
	viper.BindEnv("tracing", "tracing"){{end}}
	viper.SetDefault("gcp.projectId", "default")
	viper.BindEnv("gcp.projectId", "GCP_PROJECTID")




	lis, err := net.Listen("tcp", viper.GetString("server.binds.grpc"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

    // create the grpc server
	s := grpc.NewServer()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// register the service
	{{.Name}}.Register{{.Name | Title}}ServiceServer(s, {{.Name}}.NewService())


	// Register reflection service on gRPC server.
	reflection.Register(s)
	{{ if .Tracing }}
	if viper.GetBool("tracing") {
		exporter, err := stackdriver.NewExporter(stackdriver.Options{
			ProjectID: viper.GetString("gcp.projectId"),
			OnError: func(err error) {
				log.Warnln(errors.Wrap(err, "failed to export trace"))
			},
			/*TraceClientOptions: []option.ClientOption{
				option.WithCredentialsFile("service-account.json"),
			},*/
		})
		if err != nil {
			log.Warn(err)
		} else {
			trace.RegisterExporter(exporter)
			trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
			opts = append(opts, grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
			log.Infoln("Tracing to GCP Stackdriver Trace enabled")
		}
	}
	}{{end}}
    // if any service fails, whole app should fail
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		log.Infoln("Starting gRPC server on", viper.GetString("server.binds.grpc"))
		if err := s.Serve(lis); err != nil {
			log.Fatalf("grpc: failed to serve: %v", err)
		}

		wg.Done()
	}(){{if .WebRPC}}

	options := []grpcweb.Option{}
    // Wrap grpc server to grpc-web server
    wrappedServer := grpcweb.WrapServer(s, options...)

    handler := func(resp http.ResponseWriter, req *http.Request) {
        wrappedServer.ServeHTTP(resp, req)
    }

    httpServer := http.Server{
        Addr:    viper.GetString("server.binds.webrpc"),
        Handler: http.HandlerFunc(handler),
    }

	go func() {
		log.Infoln("starting gRPC Web server on", viper.GetString("server.binds.webrpc"))
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatalf("grpc-web: failed to serve: %v", err)
		}

		wg.Done()
	}()	{{end}}
	{{if .Profiler }}
	if viper.GetBool("profiler") {
		if err := profiler.Start(profiler.Config{Service: "{{ .Name }}"}); err != nil {
			log.Warn(err)
		} else {
			log.Infoln("GCloud profiler enabled")
		}
	}{{end}}

	go func() {
        ctx := context.Background()
        ctx, cancel := context.WithCancel(ctx)
        defer cancel()

        mux := runtime.NewServeMux()
        err := {{ .Name }}.Register{{ .Name | Title }}ServiceHandlerFromEndpoint(ctx, mux, viper.GetString("server.binds.grpc"), opts)
        if err != nil {
            log.Fatalf("gw: failed to register: %v", err)
        }

        log.Infoln("Starting gateway server on", viper.GetString("server.binds.gw"))
        log.Fatalf("gw: failed to server: %v", http.ListenAndServe(viper.GetString("server.binds.gw"), mux))
    }()

	wg.Wait()
}
