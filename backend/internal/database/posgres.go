package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() (*gorm.DB, error) {

	dsn := "host=localhost " +
		"user=cyber " +
		"password=manul " +
		"dbname=school " +
		"port=5432 " +
		"sslmode=disable " +
		"TimeZone=Asia/Yekaterinburg"

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
