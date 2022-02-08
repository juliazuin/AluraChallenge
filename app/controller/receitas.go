package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/juliazuin/AluraChallenge/app/model"
	"gorm.io/gorm"
)

type ReceitasController struct {
	Db *gorm.DB
}

func NewReceita(db *gorm.DB) *ReceitasController {
	return &ReceitasController{Db: db}
}

func (r *ReceitasController) CreateReceita(c *gin.Context) {

	var input models.Receita

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Valor == 0 || input.Descricao == "" || input.DataAtual == "" {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Faltando valores"})
	}

	Receita := models.Receita{
		Valor:     input.Valor,
		Descricao: input.Descricao,
		DataAtual: input.DataAtual,
	}

	retorno := r.Db.Create(&Receita)

	if retorno.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": retorno.Error.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": Receita})
	}
}

func (r *ReceitasController) UpdateReceita(c *gin.Context) {
	var input models.Receita

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	receitaId := c.Param("id")
	receita := models.Receita{}

	r.Db.First(&receita, receitaId)

	receita.Descricao = input.Descricao
	receita.Valor = input.Valor
	receita.DataAtual = input.DataAtual

	r.Db.Save(&receita)

	c.JSON(http.StatusOK, gin.H{"data": &receita})
}

func (r *ReceitasController) DeleteReceita(c *gin.Context) {
	receitaId := c.Param("id")
	result := r.Db.Delete(&models.Receita{}, receitaId)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"data": "Regstro deletado com sucesso"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "Nenhum registro com esse ID"})
	}
}

func (r *ReceitasController) ReceitaById(c *gin.Context) {
	receitaId := c.Param("id")
	var receitas []models.Receita
	result := r.Db.First(&receitas, receitaId)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"data": &receitas})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "Nenhum registro com esse ID"})
	}
}

func (r *ReceitasController) ListReceitas(c *gin.Context) {
	var receitas []models.Receita

	var result *gorm.DB

	if queryParams := c.Request.URL.Query().Get("descricao"); queryParams != "" {
		result = r.Db.Where("descricao LIKE ?", "%"+queryParams+"%").Find(&receitas)
	} else {
		result = r.Db.Find(&receitas)
	}
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"data": &receitas})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "Nenhum registro"})
	}
}

func (r *ReceitasController) ListReceitaByMonth(c *gin.Context) {
	var receitas []models.Receita

	receitaYear := c.Param("year")
	receitaMonth := c.Param("month")

	result := r.Db.Where("data_atual LIKE ?", "%"+receitaMonth+"/"+receitaYear+"%").Find(&receitas)

	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"data": &receitas})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "Nenhum registro"})
	}
}
