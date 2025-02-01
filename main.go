package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{}) //Auto Create Table จะลบไม่ได้
	fmt.Print("Migrate Successful")

	// product := Product{
	// 	Name:        "Tuschy_4",
	// 	Description: "MAMA3",
	// 	Price:       1003,
	// }

	// createProduct(db, &product)

	// currenProduct := getProduct(db, 2)
	// fmt.Println(currenProduct)

	// currenProduct.Name = "newtuschy"
	// currenProduct.Description = "NEW"
	// currenProduct.Price = 2000

	// updateProduct(db, currenProduct)

	// deleteProduct(db, 2)

	cur := getProducts(db)
	fmt.Println(cur)

}
