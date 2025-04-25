package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"na_novaai_server/conf"
	"na_novaai_server/internal/model"
)

var mysqldb *gorm.DB

func GetDB() *gorm.DB {
	return mysqldb
}

func InitMysql() error {
	connectionInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&autocommit=true&parseTime=true&loc=%s&multiStatements=true",
		conf.GlobalConfig.Mysql.User,
		conf.GlobalConfig.Mysql.Password,
		conf.GlobalConfig.Mysql.Addr,
		conf.GlobalConfig.Mysql.Port,
		conf.GlobalConfig.Mysql.Db,
		"Asia%2FShanghai")
	log.Println(connectionInfo)
	var err error
	mysqldb, err = gorm.Open(mysql.Open(connectionInfo), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "na_",
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}
	sqldb, err := mysqldb.DB()
	if err != nil {
		return err
	}
	sqldb.SetMaxIdleConns(3)
	sqldb.SetMaxOpenConns(30)
	sqldb.SetConnMaxLifetime(time.Hour)
	mysqldb.Logger.LogMode(logger.Error)
	if !conf.IsProd() {
		mysqldb = mysqldb.Debug()
	}

	//if err := MigrateDB(sqldb, conf.GlobalConfig.Mysql.Db); err != nil {
	//	fmt.Printf("migrate db failed: %v", err)
	//	return err
	//}

	if err := mysqldb.AutoMigrate(
		&model.Weather{},
	); err != nil {
		return err
	}

	log.Println("DB Init Success")
	return err
}
