package main

import (
	"context"
	"flag"
	_ "fmt"
	"log"
	"na_novaai_server/conf"
	"na_novaai_server/internal/api"
	"na_novaai_server/internal/database"
	nai "na_novaai_server/internal/na_interface"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcPort = ":50051"
	httpPort = ":8080"
)

func main() {
	// 解析命令行标志
	configFile := flag.String("c", "conf/server.yaml", "default conf/server.yaml")
	flag.Parse()

	// 初始化配置
	if err := conf.ConfigInit(*configFile); err != nil {
		log.Fatalf("配置初始化失败: %v", err)
	}
	// 初始化数据库连接
	err := db.RegisterDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 创建gRPC服务器
	grpcServer := grpc.NewServer()
	weatherServer := api.NewWeatherServer(db.GetDB())
	nai.RegisterNovaAIServiceServer(grpcServer, weatherServer)

	// 启动gRPC服务器
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	go func() {
		log.Printf("Starting gRPC server on port%s", grpcPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// 创建HTTP服务器（gRPC-Gateway）
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// 注册HTTP处理程序
	err = nai.RegisterNovaAIServiceHandlerFromEndpoint(
		ctx,
		mux,
		"localhost"+grpcPort,
		opts,
	)
	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}

	// 启动HTTP服务器
	log.Printf("Starting HTTP server on port%s", httpPort)
	if err := http.ListenAndServe(httpPort, mux); err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}
}
