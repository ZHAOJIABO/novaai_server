package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"na_novaai_server/api/weather"
	"na_novaai_server/conf"
	nai "na_novaai_server/internal/na_interface"
	"net"
	"net/http"
)

func main() {
	// 加载配置
	config, err := conf.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// 启动 gRPC 服务器
	go func() {
		lis, err := net.Listen("tcp", config.Server.GrpcAddress)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		weatherServer := weather.NewWeatherServer()
		nai.RegisterWeatherServiceServer(s, weatherServer)
		log.Printf("gRPC server listening at %v", lis.Addr())

		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// 启动 HTTP 服务器
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = nai.RegisterWeatherServiceHandlerFromEndpoint(
		ctx,
		mux,
		config.Server.GrpcAddress,
		opts,
	)
	if err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}

	log.Printf("HTTP server listening at %v", config.Server.HttpAddress)
	if err := http.ListenAndServe(config.Server.HttpAddress, mux); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
