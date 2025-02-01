package main

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

func createUser(db *gorm.DB, user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	result := db.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func loginUser(db *gorm.DB, user *User) (string, error) {
	//get email from user
	selectUser := new(User)
	result := db.Where("email = ?", user.Email).First(selectUser)

	if result.Error != nil {
		return "", result.Error
	}

	//compare password
	err := bcrypt.CompareHashAndPassword([]byte(selectUser.Password), []byte(user.Password))

	if err != nil {
		return "", err
	}

	// Load environment variables from .env file
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// pass = return jwt
	jwtSecretKey := os.Getenv("SECRET_KEY")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = selectUser.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}
