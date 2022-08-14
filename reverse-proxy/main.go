package main

import (
	jdconfig "GRPC-SERVE/config"
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/golang/glog"
	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	home "GRPC-SERVE/app/home"
)

// var (
// 	// command-line options:
// 	// gRPC server endpoint
// 	grpcServerEndpoint = flag.String("grpc-server-endpoint", ":6066", "gRPC server endpoint")
// )

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}
func run() error {
	httpPort := jdconfig.Cfg.Server.Port
	httpIP := jdconfig.Cfg.Server.Host
	port := jdconfig.Cfg.Ports.RpcPort
	ip := jdconfig.Cfg.Ports.RpcIp
	tcp := ip + ":" + port
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	home_err := home.RegisterLoanServiceHandlerFromEndpoint(ctx, mux, tcp, opts)
	commmon_err := errors.New("Internal Error!")

	if home_err != nil {
		fmt.Println(commmon_err)
		return commmon_err
	}

	loggedRouter := handlers.LoggingHandler(os.Stdout, mux)

	return http.ListenAndServe(httpIP+":"+httpPort, loggedRouter)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
