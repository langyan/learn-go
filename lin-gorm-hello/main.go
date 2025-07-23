package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID    uint
	Name  string
	Email string
}

func main() {
	// Connect to the database
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	// Migrate the schema
	// This automatically creates the 'users' table based on the User struct
	db.AutoMigrate(&User{})
	// Create a new user
	db.Create(&User{Name: "John Doe", Email: "john.doe@example.com"})
	// Query users
	var user User
	db.First(&user, 1) // Find user with ID 1
	fmt.Println("User found:", user.Name)
	// Update user's email
	db.Model(&user).Update("Email", "john.newemail@example.com")
	// Delete user
	db.Delete(&user)
}
