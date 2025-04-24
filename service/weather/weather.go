package weather

import (
	"context"
	nai "na_novaai_server/internal/na_interface"
)

type Service struct {}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetTomorrowWeather(ctx context.Context, req *nai.WeatherRequest) (*nai.WeatherResponse, error) {
	// 这里实现具体的业务逻辑
	// 可以调用第三方天气API或者数据库等
	return &nai.WeatherResponse{
		Temperature: "25°C",
		Condition:   "晴天",
		Humidity:    "65%",
		WindSpeed:   "3m/s",
	}, nil
}