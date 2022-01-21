package model

import "gorm.io/gorm"

type Receitas struct {
	//gorm.Model includes ID, CratedAt, UpdatedAt and DeletedAt fields out of the box
	gorm.Model
	// ID        uint    `json:"id" gorm:"primaryKey"`
	Valor     float64 `json:"valor" gorm:"not null" validate:"required"`
	Descricao string  `json:"descricao" gorm:"unique;not null" validate:"required"`
	DataAtual string  `json:"data_atual" gorm:"not null" validate:"required"`
}
