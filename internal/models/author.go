package models

import "time"

type Autor struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	Nome         string    `gorm:"type:varchar(255);not null"`
	Nacionalidade string    `gorm:"type:varchar(100)"`
	// Relacionamento: Um Autor tem muitos Livros
	Livros       []Livro   `gorm:"foreignKey:AutorID"` 
	CreatedAt    time.Time
	UpdatedAt    time.Time
}