// cmd/server/main.go
package main

import (
	"log"

	"book-crud/internal/book"
	"book-crud/internal/db"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	database.AutoMigrate(&book.Book{})

	r := gin.Default()
	r.RedirectTrailingSlash = false

	// ✅ CORS middleware MUST come before routes
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://book-ui.onrender.com",
		"https://book-ui-n456.onrender.com",
		"https://bookai045.netlify.app},
		// use frontend URL in prod
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: false,
	}))

	book.RegisterRoutes(r.Group("/books"), database)

	// ✅ Run LAST
	r.Run(":8080")

}


