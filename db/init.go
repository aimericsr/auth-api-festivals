package db

import (
	"log"

	"github.com/aimericsr/auth-api-festivals/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dbSource string) {
	db, err := gorm.Open(postgres.Open(dbSource), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	sqlStr := "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"
	tx := db.Exec(sqlStr)
	if tx.Error != nil {
		log.Fatal("Error installing extension on database")
	}

	db.AutoMigrate(&models.User{})

	//db.Create(&models.User{Name: "jedan", Email: "address", Address: "nidce", PreferdName: 23})

	db.Create(&models.User{Name: "jedan", Email: "address", Address: "nidce"})

	//db.Model(&models.User).Update("Price", 200)

	//db.Delete(&models.User, "8566749d-b301-4969-b0a0-39d025583237")
}
