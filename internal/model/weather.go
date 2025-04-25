package model

import (
	"gorm.io/gorm"
	"time"
)

// Weather 天气预报数据模型
type Weather struct {
	ID           uint           `gorm:"primarykey"`
	City         string         `gorm:"type:varchar(50);not null;index:idx_city_date"`
	ForecastDate time.Time      `gorm:"type:date;not null;index:idx_city_date"`
	Temperature  string         `gorm:"type:varchar(10);not null"`
	Humidity     string         `gorm:"type:varchar(10);not null"`
	WindSpeed    string         `gorm:"type:varchar(10);not null"`
	CreatedAt    time.Time      `gorm:"not null"`
	UpdatedAt    time.Time      `gorm:"not null"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

// TableName 指定表名
func (Weather) TableName() string {
	return "na_weather"
}
