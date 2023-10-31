package config

import (
	"log"
	"ugdc11/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
  
func InitDB() *gorm.DB  {
	dsn := "host=localhost user=postgres password=postgresql dbname=ugdc11 port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error Koneksi Database")
	}
	db.AutoMigrate(&entity.User{})
	return db
}