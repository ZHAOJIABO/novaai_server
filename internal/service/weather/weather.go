package weather

import (
	"context"
	"fmt"
	"na_novaai_server/internal/model"
	nai "na_novaai_server/internal/na_interface"
	"time"

	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewWeatherService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) GetTomorrowWeather(ctx context.Context, req *nai.WeatherRequest) (*nai.WeatherResponse, error) {
	// 计算明天的日期
	tomorrow := time.Now().AddDate(0, 0, 1)
	// 将时间设置为当天的0点0分0秒
	tomorrow = time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, tomorrow.Location())

	// 查询天气数据
	var weather model.Weather
	result := s.db.Where("city = ? AND forecast_date = ?", req.City, tomorrow).First(&weather)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("未找到城市 %s 在 %s 的天气数据", req.City, tomorrow.Format("2006-01-02"))
		}
		return nil, fmt.Errorf("查询天气数据失败: %v", result.Error)
	}

	// 转换为响应格式
	return &nai.WeatherResponse{
		Temperature: weather.Temperature,
		Humidity:    weather.Humidity,
		WindSpeed:   weather.WindSpeed,
	}, nil
}
