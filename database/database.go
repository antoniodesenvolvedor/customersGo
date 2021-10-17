package database

import (
	"customer-crud/configs"
	"customer-crud/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func GetSession() *gorm.DB {
	db, err := gorm.Open(postgres.Open(configs.GetConfig().DbConnectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func Migrate(session *gorm.DB) {
	err := session.AutoMigrate(&models.Customer{})
	if err != nil {
		log.Fatalf("Não foi possível realizar a migração de modelos")
	}
}
