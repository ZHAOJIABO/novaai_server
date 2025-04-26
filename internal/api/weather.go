package api

import (
	"context"
	nai "na_novaai_server/internal/na_interface"

	"na_novaai_server/internal/service/weather"
	"time"

	"gorm.io/gorm"
)

type WeatherServer struct {
	nai.UnimplementedNovaAIServiceServer
	weatherService *weather.Service
}

func NewWeatherServer(DB *gorm.DB) *WeatherServer {
	return &WeatherServer{
		weatherService: weather.NewWeatherService(DB),
	}
}

func (s *WeatherServer) GetTomorrowWeather(ctx context.Context, req *nai.WeatherRequest) (*nai.WeatherResponse, error) {
	return s.weatherService.GetTomorrowWeather(ctx, req)
}

var startTime = time.Now()

func (s *WeatherServer) HealthCheck(ctx context.Context, req *nai.HealthCheckRequest) (*nai.HealthCheckResponse, error) {

	return &nai.HealthCheckResponse{
		Status:  "ok",
		Version: NovaServerVersion,
		Uptime:  time.Since(startTime).Milliseconds(),
	}, nil
}
