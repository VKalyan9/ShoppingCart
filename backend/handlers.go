package main

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"github.com/gin-gonic/gin"
)

// Generate random token
func generateToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// Create User
func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already exists
	var existingUser User
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = "" // Don't send password back
	c.JSON(http.StatusCreated, user)
}

// Get all users
func getUsers(c *gin.Context) {
	var users []User
	db.Find(&users)
	
	// Remove passwords from response
	for i := range users {
		users[i].Password = ""
		users[i].Token = ""
	}
	
	c.JSON(http.StatusOK, users)
}

// Login User
func loginUser(c *gin.Context) {
	var loginData struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user User
	if err := db.Where("username = ? AND password = ?", loginData.Username, loginData.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate new token
	token := generateToken()
	db.Model(&user).Update("token", token)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user_id": user.ID,
		"username": user.Username,
	})
}

// Create Item
func createItem(c *gin.Context) {
	var item Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// Get all items
func getItems(c *gin.Context) {
	var items []Item
	db.Find(&items)
	c.JSON(http.StatusOK, items)
}

// Add to Cart
func addToCart(c *gin.Context) {
	userID := c.GetUint("user_id")
	
	var requestData struct {
		ItemID uint `json:"item_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find or create cart for user
	var cart Cart
	err := db.Where("user_id = ? AND status = ?", userID, "active").First(&cart).Error
	if err != nil {
		// Create new cart
		cart = Cart{
			UserID: userID,
			Name:   "Shopping Cart",
			Status: "active",
		}
		if err := db.Create(&cart).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Update user's cart_id
		db.Model(&User{}).Where("id = ?", userID).Update("cart_id", cart.ID)
	}

	// Check if item exists
	var item Item
	if err := db.First(&item, requestData.ItemID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	// Add item to cart
	db.Model(&cart).Association("Items").Append(&item)

	// Load cart with items
	db.Preload("Items").First(&cart, cart.ID)

	c.JSON(http.StatusOK, cart)
}

// Get all carts
func getCarts(c *gin.Context) {
	userID := c.GetUint("user_id")
	
	var carts []Cart
	db.Preload("Items").Where("user_id = ?", userID).Find(&carts)
	c.JSON(http.StatusOK, carts)
}

// Create Order
func createOrder(c *gin.Context) {
	userID := c.GetUint("user_id")

	// Find active cart for user
	var cart Cart
	if err := db.Where("user_id = ? AND status = ?", userID, "active").First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No active cart found"})
		return
	}

	// Create order
	order := Order{
		CartID: cart.ID,
		UserID: userID,
	}

	if err := db.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update cart status to ordered
	db.Model(&cart).Update("status", "ordered")

	// Clear user's cart_id
	db.Model(&User{}).Where("id = ?", userID).Update("cart_id", nil)

	c.JSON(http.StatusCreated, order)
}

// Get all orders
func getOrders(c *gin.Context) {
	userID := c.GetUint("user_id")
	
	var orders []Order
	db.Where("user_id = ?", userID).Find(&orders)
	c.JSON(http.StatusOK, orders)
}
