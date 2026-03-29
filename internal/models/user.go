package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Nome      string         `gorm:"type:varchar(255);not null" json:"nome"`
	Email     string         `gorm:"type:varchar(255);unique;not null" json:"email"`
	// Para não retornar a senha no JSON de resposta, 
	// o ideal no futuro é usar um DTO, mas por ora isso resolve o cadastro.
	Senha     string         `gorm:"type:varchar(255);not null" json:"senha"` 
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}