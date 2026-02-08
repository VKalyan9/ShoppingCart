package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// "modernc.org/sqlite"
	// "github.com/gin-contrib/cors"
)

var db *gorm.DB

func main() {
	// Initialize database
	var err error
	// db, err := gorm.Open(sqlite.Dialector{
	// 	DriverName: "sqlite",
	// 	DSN: "shopping.db",
	// }, &gorm.Config{})
	db, err = gorm.Open(sqlite.Open("shopping_cart.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate the schema
	db.AutoMigrate(&User{}, &Item{}, &Cart{}, &Order{}, &CartItem{})

	// Create some sample items
	seedItems()

	// Setup Gin router
	router := gin.Default()
	// router.Use(cors.Default())

	// Enable CORS
	router.Use(CORSMiddleware())

	// Routes
	router.POST("/users", createUser)
	router.GET("/users", getUsers)
	router.POST("/users/login", loginUser)
	
	router.POST("/items", createItem)
	router.GET("/items", getItems)
	
	router.POST("/carts", authMiddleware(), addToCart)
	router.GET("/carts", authMiddleware(), getCarts)
	
	router.POST("/orders", authMiddleware(), createOrder)
	router.GET("/orders", authMiddleware(), getOrders)

	// Start server
	router.Run(":8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func seedItems() {
	var count int64
	db.Model(&Item{}).Count(&count)
	if count == 0 {
		items := []Item{
			{Name: "Laptop", Status: "active"},
			{Name: "Mouse", Status: "active"},
			{Name: "Keyboard", Status: "active"},
			{Name: "Monitor", Status: "active"},
			{Name: "Headphones", Status: "active"},
		}
		db.Create(&items)
	}
}
