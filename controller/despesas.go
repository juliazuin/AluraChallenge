package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "github.com/juliazuin/AluraChallenge/dabatase"
	models "github.com/juliazuin/AluraChallenge/model"
	"gorm.io/gorm"
)

type DespesasRepo struct {
	Db *gorm.DB
}

func New() *DespesasRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Despesas{})
	return &DespesasRepo{Db: db}
}

// postAlbums adds an album from JSON received in the request body.
func (d *DespesasRepo) PostDespesa(c *gin.Context) {
	var input models.Despesas

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add the new album to the slice.
	teste := models.Despesas{Valor: input.Valor, Descricao: input.Descricao, DataAtual: input.DataAtual}
	d.Db.Create(&teste)
	c.JSON(http.StatusOK, gin.H{"data": teste})
}
