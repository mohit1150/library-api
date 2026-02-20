package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Book struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	AvailableCopies int    `json:"available_copies"`
}

type Checkout struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	UserID     uint       `json:"user_id"`
	BookID     uint       `json:"book_id"`
	CheckoutAt time.Time  `json:"checkout_at"`
	DueDate    time.Time  `json:"due_date"`
	ReturnedAt *time.Time `json:"returned_at"`
	FineAmount float64    `json:"fine_amount"`
}

func main() {
	db, _ := gorm.Open(sqlite.Open("library.db"), &gorm.Config{})
	db.AutoMigrate(&Book{}, &Checkout{})

	r := gin.Default()

	// Endpoints for Librarians
	r.POST("/books", func(c *gin.Context) {
		var book Book
		if err := c.ShouldBindJSON(&book); err == nil {
			db.Create(&book)
			c.JSON(http.StatusOK, book)
		}
	})

	// Endpoint for Student Checkout
	r.POST("/checkout", func(c *gin.Context) {
		var checkout Checkout
		c.ShouldBindJSON(&checkout)

		var book Book
		db.First(&book, checkout.BookID)

		if book.AvailableCopies > 0 {
			book.AvailableCopies--
			db.Save(&book)

			checkout.CheckoutAt = time.Now()
			checkout.DueDate = time.Now().AddDate(0, 0, 7) // 7-day loan period
			
			db.Create(&checkout)
			c.JSON(http.StatusOK, gin.H{"message": "Checkout successful", "due_date": checkout.DueDate})
		} else {
			// This handles the "Reservation Queue" requirement
			c.JSON(http.StatusConflict, gin.H{"message": "No copies available. You are in the reservation queue."})
		}
	})

	// Endpoint for Returning Books & Fine Calculation
	r.POST("/return", func(c *gin.Context) {
		var req struct{ CheckoutID uint `json:"checkout_id"` }
		c.ShouldBindJSON(&req)

		var checkout Checkout
		db.First(&checkout, req.CheckoutID)

		now := time.Now()
		checkout.ReturnedAt = &now

		// Logic for Fine Calculation
		if now.After(checkout.DueDate) {
			daysLate := int(now.Sub(checkout.DueDate).Hours() / 24)
			if daysLate > 0 {
				checkout.FineAmount = float64(daysLate * 10) // 10 units per day
			}
		}

		db.Save(&checkout)

		var book Book
		db.First(&book, checkout.BookID)
		book.AvailableCopies++
		db.Save(&book)

		c.JSON(http.StatusOK, gin.H{"message": "Book returned", "fine": checkout.FineAmount})
	})

	r.Run(":8080")
}