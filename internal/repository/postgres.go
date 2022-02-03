package repository


import (
	"fmt"
	"github.com/red-life/simple-authentication-with-jwt/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgresConnection(cfg *config.PostgresConfig) *gorm.DB{
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Tehran", cfg.Host,cfg.Username, cfg.Password, cfg.DB, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}
	return db
}