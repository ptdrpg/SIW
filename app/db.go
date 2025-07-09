package app

import (
	"SI/model"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connexion() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erreur de connexion DB: %v", err)
	}
	
	err = db.AutoMigrate(
		&model.User{},
		&model.Member{},
	)
	if err != nil {
		log.Fatalf("Erreur migration DB: %v", err)
	}

	// check if admin already exist
	var count int64
	db.Model(&model.User{}).Where("email = ?", "admin@gmail.com").Count(&count)

	if count == 0 {
		// Créer uuser admin
		password := "1234admin"
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Erreur hash mot de passe admin: %v", err)
		}

		adminUser := model.User{
			Name:     "admin",
			Email:    "admin@gmail.com",
			Password: string(hashedPass),
		}

		if err := db.Create(&adminUser).Error; err != nil {
			log.Fatalf("Erreur creating admin: %v", err)
		}
	}

	DB = db
}
