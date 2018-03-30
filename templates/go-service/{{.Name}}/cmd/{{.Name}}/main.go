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
    "github.com/evalphobia/logrus_sentry"{{end}}{{if .Gateway}}
    "net/http"
    "context"
    "github.com/grpc-ecosystem/grpc-gateway/runtime"{{end}}
)

func init() {
    log.SetFormatter(&log.JSONFormatter{})

    log.SetOutput(os.Stdout)

    log.SetLevel(log.DebugLevel){{if .Sentry}}

    // Init Logrus with Sentry
    sentryDsn, ok := os.LookupEnv("SENTRY_DSN")
    if ok {
        hook, err := logrus_sentry.NewSentryHook(sentryDsn, []log.Level{
            log.PanicLevel,
            log.FatalLevel,
            log.ErrorLevel,
        })

        if err == nil {
            hook.Timeout = 20 * time.Second
            hook.StacktraceConfiguration.Enable = true
            log.AddHook(hook)
        }
    } else {
        log.Warnln("No SENTRY_DSN was found, sentry reporting won't work")
    }{{end}}
}

func main(){

    viper.SetDefault("db.conn", "host=db-svc user=goo dbname=goo sslmode=disable password=goo")
	viper.SetDefault("server.binds.grpc", ":50051"){{if .WebRPC}}
	viper.SetDefault("server.binds.webrpc", ":50060"){{end}}{{if .Gateway}}
	viper.SetDefault("server.binds.gw", ":80"){{end}}

	lis, err := net.Listen("tcp", viper.GetString("server.binds.grpc"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

    // create the grpc server
	s := grpc.NewServer()

	// register the service
	{{.Name}}.Register{{.Name | Title}}ServiceServer(s, {{.Name}}.NewService())


	// Register reflection service on gRPC server.
	reflection.Register(s)

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
	}()	{{end}}	{{if .Gateway}}

	go func() {
        ctx := context.Background()
        ctx, cancel := context.WithCancel(ctx)
        defer cancel()

        mux := runtime.NewServeMux()
        opts := []grpc.DialOption{grpc.WithInsecure()}
        err := {{ .Name }}.Register{{ .Name | Title }}ServiceHandlerFromEndpoint(ctx, mux, viper.GetString("server.binds.grpc"), opts)
        if err != nil {
            log.Fatalf("gw: failed to register: %v", err)
        }

        log.Infoln("Starting gateway server on", viper.GetString("server.binds.gw"))
        log.Fatalf("gw: failed to server: %v", http.ListenAndServe(viper.GetString("server.binds.gw"), mux))
    }(){{end}}

	wg.Wait()
}