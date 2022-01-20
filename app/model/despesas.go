package model

import "gorm.io/gorm"

type Despesas struct {
	gorm.Model
	ID        uint    `json:"id" gorm:"primaryKey"`
	Valor     float64 `json:"valor"`
	Descricao string  `json:"descricao" gorm:"unique"`
	DataAtual string  `json:"data_atual"`
}
