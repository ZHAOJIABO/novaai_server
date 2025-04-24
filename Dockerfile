FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o weather-service ./cmd/main.go

EXPOSE 8080 50051

CMD ["./weather-service"]