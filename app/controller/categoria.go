package controller

import (
	models "github.com/juliazuin/AluraChallenge/app/model"
	"gorm.io/gorm"
)

type CategoriaController struct {
	Db *gorm.DB
}

func NewCategoria(db *gorm.DB) *CategoriaController {
	return &CategoriaController{Db: db}
}

func (c *CategoriaController) SeedCategorias() {
	c.Db.Model(&models.Categoria{}).Create([]map[string]interface{}{
		{"nome": "Alimentação"},
		{"nome": "Saúde"},
		{"nome": "Moradia"},
		{"nome": "Transporte"},
		{"nome": "Educação"},
		{"nome": "Lazer"},
		{"nome": "Imprevistos"},
		{"nome": "Outras"},
	})
}
