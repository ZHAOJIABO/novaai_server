package weather

import (
	"context"
	"database/sql"
	"fmt"
	db2 "na_novaai_server/internal/database"
	nai "na_novaai_server/internal/na_interface"
	"time"
)

type Service struct {
	//db *sql.DB
}

func NewService() *Service {
	return &Service{
		//	db: db,
	}
}

func (s *Service) GetTomorrowWeather(ctx context.Context, req *nai.WeatherRequest) (*nai.WeatherResponse, error) {
	db := db2.GetDB()
	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

	query := `
		SELECT temperature, condition, humidity, wind_speed 
		FROM weather_forecasts 
		WHERE city = ? AND forecast_date = ?
		LIMIT 1
	`

	var response nai.WeatherResponse
	err := db.Raw(query, req.City, tomorrow).Row().Scan(
		&response.Temperature,
		&response.Condition,
		&response.Humidity,
		&response.WindSpeed,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("no weather data found for city %s on %s", req.City, tomorrow)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query weather data: %v", err)
	}

	return &response, nil
}
