package model

import (
	"gorm.io/gorm"
)

type Despesas struct {
	gorm.Model
	// ID        uint    `json:"id" gorm:"primaryKey"`
	Valor     float64 `json:"valor" gorm:"not null" validate:"required"`
	Descricao string  `json:"descricao" gorm:"unique;not null" validate:"required"`
	DataAtual string  `json:"data_atual" gorm:"not null" validate:"required"`
}
