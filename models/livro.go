package models

import (
	"time"

	"gorm.io/gorm"
)

type Livro struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	Nome       string         `json:"nome"`
	Descricao  string         `json:"descricao"`
	PrecoMedio float32        `json:"preco_medio"`
	Autor      string         `json:"autor"`
	ImageURL   string         `json:"img_url"`
	CreatedAt  time.Time      `json:"created"`
	UpdatedAt  time.Time      `json:"updated"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted"`
}
