FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o novaai-server cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/novaai-server .
COPY conf/server.yaml /app/conf/
EXPOSE 8080 50051
CMD ["./novaai-server"]