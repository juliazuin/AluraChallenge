package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "github.com/juliazuin/AluraChallenge/app/dabatase"
	models "github.com/juliazuin/AluraChallenge/app/model"
	"gorm.io/gorm"
)

type ReceitasController struct {
	Db *gorm.DB
}

func NewReceita() *ReceitasController {
	db := database.InitDb()
	db.AutoMigrate(&models.Receitas{})
	return &ReceitasController{Db: db}
}

// postAlbums adds an album from JSON received in the request body.
func (r *ReceitasController) CreateReceita(c *gin.Context) {
	//var validate = validator.New()

	var input models.Receitas

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Valor == 0 || input.Descricao == "" || input.DataAtual == "" {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Faltando valores"})
	}

	Receita := models.Receitas{
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
	/*errs := validate.Struct(Receita)
	  print(errs)
	  if len(errs.(validator.ValidationErrors)) > 0 {
	  	fmt.Printf("Errs:\n%+v\n", errs)
	  	c.JSON(http.StatusBadRequest, gin.H{"data": "Faltando valores"})

	  } else {
	  	d.Db.Create(&Receita)
	  	c.JSON(http.StatusOK, gin.H{"data": Receita})
	  }*/
}

func (r *ReceitasController) UpdateReceita(c *gin.Context) {
	var input models.Receitas

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	receitaId := c.Param("id")
	receita := models.Receitas{}

	r.Db.First(&receita, receitaId)

	receita.Descricao = input.Descricao
	receita.Valor = input.Valor
	receita.DataAtual = input.DataAtual

	r.Db.Save(&receita)

	c.JSON(http.StatusOK, gin.H{"data": &receita})
}

func (r *ReceitasController) DeleteReceita(c *gin.Context) {
	receitaId := c.Param("id")
	result := r.Db.Delete(&models.Receitas{}, receitaId)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"data": "Regstro deletado com sucesso"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "Nenhum registro com esse ID"})
	}
}

func (r *ReceitasController) ReceitaById(c *gin.Context) {
	receitaId := c.Param("id")
	receitas := []models.Receitas{}
	result := r.Db.First(&receitas, receitaId)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"data": &receitas})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "Nenhum registro com esse ID"})
	}
}

func (r *ReceitasController) ListReceitas(c *gin.Context) {
	receitas := []models.Receitas{}
	result := r.Db.Find(&receitas)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"data": &receitas})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "Nenhum registro"})
	}
}
