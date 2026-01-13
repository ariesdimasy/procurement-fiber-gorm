package config

import (
	"fmt"
	"log"
	"os"
	"procurement-fiber-gorm/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func DatabaseConnection() {
	dsn := "root:@tcp(127.0.0.1:3306)/procurement_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database")
		log.Fatal("Failed to Conenct to the database! \n ", err.Error())
		os.Exit(2)
		panic("failed to connect database")

	}

	fmt.Println("RUN")

	log.Println("Connected to database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations...")

	db.AutoMigrate(&models.User{}, &models.Item{}, &models.Supplier{}, &models.Purchasing{}, &models.PurchasingDetail{})

	DB = db
}
