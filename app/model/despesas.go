package model

import (
	"gorm.io/gorm"
)

type Despesa struct {
	gorm.Model
	// ID        uint    `json:"id" gorm:"primaryKey"`
	Valor       float64 `json:"valor" gorm:"not null"`
	Descricao   string  `json:"descricao" gorm:"unique;not null"`
	DataAtual   string  `json:"data_atual" gorm:"not null"`
	CategoriaId uint
}
