package main

import (
	"log"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `json: name`
	Description string `json: description`
	Price       int    `json: price`
}

func createProduct(db *gorm.DB, product *Product) error {
	result := db.Create(&product)

	if result.Error != nil {
		return result.Error
	}

	// fmt.Printf("Result Error : %v", result.Error)
	return nil
}

func updateProduct(db *gorm.DB, product *Product) error {
	result := db.Save(&product)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func getProduct(db *gorm.DB, id uint) *Product {
	var product Product
	result := db.First(&product, id)

	if result.Error != nil {
		// log.Fatalf("Error get product: %v", result.Error)
		return nil
	}

	return &product
	// fmt.Printf("Get Product Successful %v", result.RowsAffected)
}

func getProducts(db *gorm.DB) []Product {
	var product []Product
	result := db.Find(&product)

	if result.Error != nil {
		log.Fatalf("Error get product: %v", result.Error)
	}

	// fmt.Println("Get Products Successful", result.RowsAffected)
	return product
}

func deleteProduct(db *gorm.DB, id uint) error {
	var product Product
	result := db.Delete(&product, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
