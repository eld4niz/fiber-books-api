package model

import (
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type Book struct {
    gorm.Model
    ID         uuid.UUID `gorm:"type:uuid"`
	BookName   string    `json:"book_name"`
	AuthorName string    `json:"author_name"`
	Category   string    `json:"category"`
	Price	  float64   `json:"price"`

}
