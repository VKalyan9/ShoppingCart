package main

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Password  string    `gorm:"not null" json:"password,omitempty"`
	Token     string    `json:"token,omitempty"`
	CartID    *uint     `json:"cart_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Item struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type Cart struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UserID    uint       `json:"user_id"`
	Name      string     `json:"name"`
	Status    string     `json:"status"`
	Items     []Item     `gorm:"many2many:cart_items;" json:"items,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
}

type Order struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CartID    uint      `json:"cart_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CartItem struct {
	CartID uint `json:"cart_id"`
	ItemID uint `json:"item_id"`
}
