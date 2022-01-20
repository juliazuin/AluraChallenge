package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "github.com/juliazuin/AluraChallenge/app/dabatase"
	models "github.com/juliazuin/AluraChallenge/app/model"
	"gorm.io/gorm"
)

type DespesasRepo struct {
	Db *gorm.DB
}

func New() *DespesasRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Despesas{}, &models.Receitas{})
	return &DespesasRepo{Db: db}
}

// postAlbums adds an album from JSON received in the request body.
func (d *DespesasRepo) CreateDespesa(c *gin.Context) {
	var input models.Despesas

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add the new album to the slice.
	despesa := models.Despesas{
		Valor:     input.Valor,
		Descricao: input.Descricao,
		DataAtual: input.DataAtual,
	}
	
	d.Db.Create(&despesa)
	c.JSON(http.StatusOK, gin.H{"data": despesa})
}
