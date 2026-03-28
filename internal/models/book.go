package models

import (
	"time"
	"gorm.io/gorm"
)

type Livro struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Titulo    string    `gorm:"type:varchar(255);not null"`
	Descricao string    `gorm:"type:text"`
	
	// Chave Estrangeira: Aponta para o ID do Autor
	AutorID   uint      `gorm:"not null"`
	// Propriedade de navegação: Permite ao GORM carregar os dados do Autor automaticamente
	Autor     Author     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	
	Publicado bool      `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// Soft Delete
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Exemplo de busca para caso queira ver os itens que sofreram soft delete
// r.db.Unscoped().Find(&livros)