package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"na_novaai_server/conf"
	"na_novaai_server/internal/api"
	db "na_novaai_server/internal/database"
	nai "na_novaai_server/internal/na_interface"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type App struct {
	httpServer *http.Server
	grpcServer *grpc.Server
	grpcLis    net.Listener
	//serviceProvider bootstrap.ServiceInterface
	gwmux *runtime.ServeMux
}

func NewApp() (*App, error) {
	app := &App{}

	// 初始化 gRPC 服务器
	app.grpcServer = grpc.NewServer()
	weatherService := api.NewWeatherServer()
	nai.RegisterWeatherServiceServer(app.grpcServer, weatherService)

	// 创建 gRPC 监听器，使用固定地址而不是自动分配
	grpcAddr := conf.GlobalConfig.Server.GrpcAddress
	if grpcAddr == "" {
		grpcAddr = ":50051"
	}
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %v", err)
	}
	app.grpcLis = lis

	// 初始化 HTTP 服务器，使用固定地址
	httpAddr := conf.GlobalConfig.Server.HttpAddress
	if httpAddr == "" {
		httpAddr = ":8080"
	}
	app.httpServer = &http.Server{
		Addr: httpAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Server is running"))
		}),
	}

	return app, nil
}

func (a *App) Run() error {
	// 启动 gRPC 服务器
	go func() {
		log.Printf("gRPC server listening at %v", a.grpcLis.Addr())
		if err := a.grpcServer.Serve(a.grpcLis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	// 启动 HTTP 服务器
	go func() {
		log.Printf("HTTP server listening at %v", a.httpServer.Addr)
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to serve HTTP: %v", err)
		}
	}()

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 关闭服务器
	a.grpcServer.GracefulStop()
	if err := a.httpServer.Shutdown(context.Background()); err != nil {
		return fmt.Errorf("HTTP server shutdown failed: %v", err)
	}

	return nil
}

func main() {
	// 解析命令行标志
	configFile := flag.String("c", "conf/server.yaml", "default conf/server.yaml")
	flag.Parse()

	// 初始化配置
	if err := conf.ConfigInit(*configFile); err != nil {
		log.Fatalf("配置初始化失败: %v", err)
	}

	// 初始化数据库连接
	log.Println("正在初始化数据库...")
	if err := db.InitMysql(); err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}

	hostName, _ := os.Hostname()
	log.Printf("na_novaai_server v%s : server %s - %s\n", api.NovaServerVersion, conf.GlobalConfig.ServerName, hostName)

	// 创建并运行应用程序
	app, err := NewApp()
	if err != nil {
		log.Fatalf("创建应用程序失败: %v", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("运行应用程序失败: %v", err)
	}
}
