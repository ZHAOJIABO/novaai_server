package api

import (
	"context"
	"gorm.io/gorm"
	nai "na_novaai_server/internal/na_interface"
	"na_novaai_server/internal/service/weather"
)

type WeatherServer struct {
	nai.UnimplementedWeatherServiceServer
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
