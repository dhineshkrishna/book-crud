// internal/book/handler.go
package book

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	r.POST("/", createBook(db))
	r.GET("/", getBooks(db))
}

func createBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var b Book
		if err := c.ShouldBindJSON(&b); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&b)
		c.JSON(http.StatusCreated, b)
	}
}

func getBooks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var books []Book
		db.Find(&books)
		c.JSON(http.StatusOK, books)
	}
}
