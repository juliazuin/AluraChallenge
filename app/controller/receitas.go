package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "github.com/juliazuin/AluraChallenge/app/dabatase"
	models "github.com/juliazuin/AluraChallenge/app/model"
	"gorm.io/gorm"
)

type ReceitasRepo struct {
	Db *gorm.DB
}

func New() *ReceitasRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Receitas{})
	return &ReceitasRepo{Db: db}
}

// postAlbums adds an album from JSON received in the request body.
func (d *ReceitasRepo) CreateReceita(c *gin.Context) {
	//var validate = validator.New()

	var input models.Receitas

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Valor != 0 || input.Descricao != "" || input.DataAtual != "" {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Faltando valores"})
	}

	Receita := models.Receitas{
		Valor:     input.Valor,
		Descricao: input.Descricao,
		DataAtual: input.DataAtual,
	}

	retorno := d.Db.Create(&Receita)

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

func (*ReceitasRepo) UpdateReceita(c *gin.Context) {

}

func (*ReceitasRepo) DeleteReceita(c *gin.Context) {

}

func (*ReceitasRepo) ReceitaById(c *gin.Context) {

}
