package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
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
	// Create Table
	db.AutoMigrate(&Product{}, &User{}) //Auto Create Table จะลบไม่ได้
	fmt.Print("Migrate Successful")

	// Product API
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Get("/product", func(c *fiber.Ctx) error {
		return c.JSON(getProducts(db))
	})

	app.Get("/product/:id", func(c *fiber.Ctx) error {
		uint_id, err := strconv.Atoi(c.Params("id"))
		id := uint(uint_id)

		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		cur := getProduct(db, id)
		if cur == nil {
			return c.Status(fiber.StatusNotFound).SendString("Not Found")
		}

		return c.JSON(cur)
	})

	app.Post("/product", func(c *fiber.Ctx) error {
		product := new(Product)

		if err := c.BodyParser(product); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		err = createProduct(db, product)

		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.JSON(fiber.Map{
			"message": "Create Product Successful",
		})
	})

	app.Put("/product/:id", func(c *fiber.Ctx) error {
		uint_id, err := strconv.Atoi(c.Params("id"))
		id := uint(uint_id)

		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		product := new(Product)

		if err := c.BodyParser(product); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		product.ID = id
		err = updateProduct(db, product)

		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.JSON(fiber.Map{
			"message": "Update Product Successful",
		})
	})

	app.Delete("/product/:id", func(c *fiber.Ctx) error {
		uint_id, err := strconv.Atoi(c.Params("id"))
		id := uint(uint_id)

		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		err = deleteProduct(db, id)

		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.JSON(fiber.Map{
			"message": "Delete Product Successful",
		})
	})
	// env

	app.Get("/api/config", func(c *fiber.Ctx) error {
		secretKey := os.Getenv("SECRET_KEY")
		if secretKey == "" {
			secretKey = "defaultSecret"
		}

		return c.JSON(fiber.Map{
			"secret_key": secretKey,
		})
	})

	// User API

	app.Post("/register", func(c *fiber.Ctx) error {
		user := new(User)

		if err := c.BodyParser(user); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		err = createUser(db, user)

		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.JSON(fiber.Map{
			"message": "Create User Successful",
		})
	})

	app.Listen(":8080")
}
