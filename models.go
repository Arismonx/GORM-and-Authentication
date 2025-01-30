package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       int
}

func createProduct(db *gorm.DB, product *Product) {
	result := db.Create(&product)

	if result.Error != nil {
		log.Fatalf("Error creating book: %v", result.Error)
	}

	fmt.Println("Create Successful")
}
