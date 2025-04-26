# 构建阶段
#FROM golang:1.21-alpine AS builder
FROM golang:1.23 AS builder
WORKDIR /app
# 设置 Go 代理为国内镜像（例如，阿里云）
ENV GO111MODULE=on
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

# 复制 go.mod 和 go.sum 优化缓存
COPY go.mod go.sum ./
RUN go mod download

# 再复制源码
COPY . .

# 构建 Go 可执行文件
RUN CGO_ENABLED=0 GOOS=linux go build -o novaai-server cmd/main.go

# 运行阶段（更小更安全的镜像）
FROM alpine:latest

WORKDIR /app

# 拷贝可执行文件
COPY --from=builder /app/novaai-server ./

# 拷贝配置文件（注意路径）
COPY conf/server.yaml ./conf/

# 设置环境变量（如需要，可选）
# ENV GIN_MODE=release

# 开放端口（供K8s使用）
EXPOSE 8080 50051

# 启动命令
CMD ["./novaai-server"]


# 使用官方 Golang 镜像作为构建环境
#FROM golang:1.21-alpine AS builder
#
#WORKDIR /app
#COPY . .
#
## 构建你的 Go 程序
#RUN go build -o app .
#
## 使用更小的运行时镜像
#FROM alpine:latest
#
#WORKDIR /root/
#COPY --from=builder /app/app .
#
## 指定容器启动时执行的命令
#CMD ["./app"]
