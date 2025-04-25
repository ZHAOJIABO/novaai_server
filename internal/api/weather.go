package api

import (
	"context"
	nai "na_novaai_server/internal/na_interface"
	"na_novaai_server/internal/service/weather"
)

type WeatherServer struct {
	nai.UnimplementedWeatherServiceServer
	weatherService *weather.Service
}

func NewWeatherServer() *WeatherServer {
	return &WeatherServer{
		weatherService: weather.NewService(),
	}
}

func (s *WeatherServer) GetTomorrowWeather(ctx context.Context, req *nai.WeatherRequest) (*nai.WeatherResponse, error) {
	return s.weatherService.GetTomorrowWeather(ctx, req)
}
