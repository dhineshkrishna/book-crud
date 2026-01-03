// internal/book/model.go
package book

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  float64 `json:"price"`
}
