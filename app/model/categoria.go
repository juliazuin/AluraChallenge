package model

import "gorm.io/gorm"

//categorias tem mtas despesas
type Categoria struct {
	ID   int    `gorm:"primaryKey"`
	Nome string `json:"nome"`
	//Despesas    []Despesa    `gorm:"foreignKey:UserRefer"`
	//Despesas []Despesa `gorm:"foreignKey:CategoriaId"`
}

type Despesa struct {
	gorm.Model
	// ID        uint    `json:"id" gorm:"primaryKey"`
	Valor       float64 `json:"valor" gorm:"not null"`
	Descricao   string  `json:"descricao" gorm:"unique;not null"`
	DataAtual   string  `json:"data_atual" gorm:"not null"`
	CategoriaId uint    `json:"categoria_id"`
}
