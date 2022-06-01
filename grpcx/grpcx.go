package grpcx

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
)

func Init() {
	server := grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	listen, err := net.Listen("tcp", ":19900")
	if err != nil {
		return
	}

	go func() {
		err = server.Serve(listen)
		if err != nil {
			return
		}
	}()

	//registry.Get("local").Discovery()

}
