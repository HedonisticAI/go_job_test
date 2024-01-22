package models

import (
	"log/slog"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func form_url() string {
	dbURL := "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@localhost:" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME")
	slog.Info(dbURL)
	return dbURL
}

func Init_Database() *gorm.DB {
	dsn := form_url()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error(err.Error())
		return nil
	}
	err = db.AutoMigrate(&Human{})
	if err != nil {
		slog.Error("DB migration failed" + err.Error())
		return nil
	}
	return db
}
