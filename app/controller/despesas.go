package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "github.com/juliazuin/AluraChallenge/app/dabatase"
	models "github.com/juliazuin/AluraChallenge/app/model"
	"gorm.io/gorm"
)

type DespesasController struct {
	Db *gorm.DB
}

func NewDespesa() *DespesasController {
	db := database.InitDb()
	db.AutoMigrate(&models.Despesas{})
	return &DespesasController{Db: db}
}

// postAlbums adds an album from JSON received in the request body.
func (d *DespesasController) CreateDespesa(c *gin.Context) {
	//var validate = validator.New()

	var input models.Despesas

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Valor != 0 || input.Descricao != "" || input.DataAtual != "" {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Faltando valores"})
		return
	}

	despesa := models.Despesas{
		Valor:     input.Valor,
		Descricao: input.Descricao,
		DataAtual: input.DataAtual,
	}

	retorno := d.Db.Create(&despesa)

	if retorno.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": retorno.Error.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": despesa})
	}
	/*errs := validate.Struct(despesa)
	  print(errs)
	  if len(errs.(validator.ValidationErrors)) > 0 {
	  	fmt.Printf("Errs:\n%+v\n", errs)
	  	c.JSON(http.StatusBadRequest, gin.H{"data": "Faltando valores"})

	  } else {
	  	d.Db.Create(&despesa)
	  	c.JSON(http.StatusOK, gin.H{"data": despesa})
	  }*/
}

func (d *DespesasController) UpdateDespesa(c *gin.Context) {
	var input models.Despesas

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	despesaId := c.Param("id")
	despesa := models.Despesas{}

	d.Db.First(&despesa, despesaId)

	despesa.Descricao = input.Descricao
	despesa.Valor = input.Valor
	despesa.DataAtual = input.DataAtual

	d.Db.Save(&despesa)

	c.JSON(http.StatusOK, gin.H{"data": &despesa})
}

func (d *DespesasController) DeleteDespesa(c *gin.Context) {
	despesaId := c.Param("id")
	result := d.Db.Delete(&models.Despesas{}, despesaId)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"data": "Regstro deletado com sucesso"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "Nenhum registro com esse ID"})
	}

}

func (d *DespesasController) DespesaById(c *gin.Context) {
	despesaId := c.Param("id")
	despesas := []models.Despesas{}
	result := d.Db.First(&despesas, despesaId)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"data": &despesas})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "Nenhum registro com esse ID"})
	}
}

func (d *DespesasController) ListDespesas(c *gin.Context) {
	despesas := []models.Despesas{}
	result := d.Db.Find(&despesas)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"data": &despesas})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "Nenhum registro"})
	}
}
