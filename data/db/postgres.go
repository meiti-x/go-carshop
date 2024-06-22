package db

import (
	"fmt"
	"log"
	"time"

	"github.com/meiti-x/golang-web-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDb(cfg *config.Config) error {
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran", cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DbName, cfg.Postgres.SSLMode)

	dbClient, err := gorm.Open(postgres.Open(cnn), &gorm.Config{})

	if err != nil {
		return err
	}
	sqldb, _ := dbClient.DB()

	err = sqldb.Ping()

	if err != nil {
		return err
	}
	sqldb.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqldb.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqldb.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)

	log.Println("DB connection established")
	return nil
}

func GetDbClient() *gorm.DB {
	return dbClient
}

func CloseDbClient() {
	conn, _ := dbClient.DB()
	conn.Close()
}
