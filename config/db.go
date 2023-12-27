package config

import (
	"awesomeProject/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB сущность базы данных
var DB *gorm.DB

// Инициализация базы данных PgSQL
func InitPgSQL() error {
	var db *gorm.DB

	dsn := "host=localhost user=postgres password=1234 dbname=Uchoba port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	DB = db

	err = db.Table("students").AutoMigrate(domain.Student{})
	if err != nil {
		return err
	}

	err = db.Table("teachers").AutoMigrate(&domain.Teachers{})
	if err != nil {
		return err
	}

	return nil
}
