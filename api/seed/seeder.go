package seed

import (
	"DemoProject/api/models"
	"github.com/jinzhu/gorm"
	"log"
)

func Load(db *gorm.DB) {

	err := db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.Resource{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

}

