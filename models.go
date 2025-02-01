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
		log.Fatalf("Error creating product: %v", result.Error)
	}

	fmt.Println("Create Successful")
}

func updateProduct(db *gorm.DB, product *Product) {
	result := db.Save(&product)

	if result.Error != nil {
		log.Fatalf("Error creating product: %v", result.Error)
	}

	fmt.Println("Updete Successful")
}

func getProduct(db *gorm.DB, id uint) *Product {
	var product Product
	result := db.First(&product, id)

	if result.Error != nil {
		log.Fatalf("Error get product: %v", result.Error)
	}

	return &product
	// fmt.Printf("Get Product Successful %v", result.RowsAffected)
}

func getProducts(db *gorm.DB) *Product {
	var product Product
	result := db.Find(&product)

	if result.Error != nil {
		log.Fatalf("Error get product: %v", result.Error)
	}

	// fmt.Println("Get Products Successful", result.RowsAffected)
	return &product
}

func deleteProduct(db *gorm.DB, id uint) {
	var product Product
	result := db.Delete(&product, id)

	if result.Error != nil {
		log.Fatalf("Error delete product: %v", result.Error)
	}

	fmt.Println("Delete Successful")
}
