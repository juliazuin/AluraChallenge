package controller

import (
	database "github.com/juliazuin/AluraChallenge/app/dabatase"
	models "github.com/juliazuin/AluraChallenge/app/model"
	"gorm.io/gorm"
)

type CategoriaController struct {
	Db *gorm.DB
}

func NewCategoria() *CategoriaController {
	db := database.InitDb()
	db.AutoMigrate(&models.Categorias{})
	return &CategoriaController{Db: db}
}

func (c *CategoriaController) SeedCategorias() {
	c.Db.Model(&models.Categorias{}).Create([]map[string]interface{}{
		{"Name": "Alimentação"},
		{"Name": "Saúde"},
		{"Name": "Moradia"},
		{"Name": "Transporte"},
		{"Name": "Educação"},
		{"Name": "Lazer"},
		{"Name": "Imprevistos"},
		{"Name": "Outras"},
	})
	return
}
