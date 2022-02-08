package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/juliazuin/AluraChallenge/app/model"
	"gorm.io/gorm"
)

type DespesasController struct {
	Db *gorm.DB
}

func NewDespesa(db *gorm.DB) *DespesasController {
	return &DespesasController{Db: db}
}

func (d *DespesasController) CreateDespesa(c *gin.Context) {

	var input models.Despesa

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if any value is empty return bad request
	if input.Valor == 0 || input.Descricao == "" || input.DataAtual == "" {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Faltando valores"})
		return
	}

	// if categoria_id is empty return categoria_id=8 (outras)
	if input.CategoriaId == 0 {
		input.CategoriaId = 8
	}

	despesa := models.Despesa{
		Valor:       input.Valor,
		Descricao:   input.Descricao,
		DataAtual:   input.DataAtual,
		CategoriaId: input.CategoriaId,
	}

	retorno := d.Db.Create(&despesa)

	if retorno.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": retorno.Error.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": despesa})
	}
}

func (d *DespesasController) UpdateDespesa(c *gin.Context) {
	var input models.Despesa

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	despesaId := c.Param("id")
	despesa := models.Despesa{}

	d.Db.First(&despesa, despesaId)

	despesa.Descricao = input.Descricao
	despesa.Valor = input.Valor
	despesa.DataAtual = input.DataAtual

	d.Db.Save(&despesa)

	c.JSON(http.StatusOK, gin.H{"data": &despesa})
}

func (d *DespesasController) DeleteDespesa(c *gin.Context) {
	despesaId := c.Param("id")
	result := d.Db.Delete(&models.Despesa{}, despesaId)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"data": "Regstro deletado com sucesso"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "Nenhum registro com esse ID"})
	}

}

func (d *DespesasController) DespesaById(c *gin.Context) {
	despesaId := c.Param("id")
	var despesas []models.Despesa
	result := d.Db.First(&despesas, despesaId)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"data": &despesas})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "Nenhum registro com esse ID"})
	}
}

func (d *DespesasController) ListDespesas(c *gin.Context) {
	var despesas []models.Despesa
	var result *gorm.DB

	if queryParams := c.Request.URL.Query().Get("descricao"); queryParams != "" {
		result = d.Db.Where("descricao LIKE ?", "%"+queryParams+"%").Find(&despesas)
	} else {
		result = d.Db.Find(&despesas)
	}
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"data": &despesas})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "Nenhum registro"})
	}
}

func (d *DespesasController) ListDespesaByMonth(c *gin.Context) {
	var despesas []models.Despesa

	despesaYear := c.Param("year")
	despesaMonth := c.Param("month")

	result := d.Db.Where("data_atual LIKE ?", "%"+despesaYear+"/"+despesaMonth+"%").Find(&despesas)

	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"data": &despesas})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "Nenhum registro"})
	}
}
