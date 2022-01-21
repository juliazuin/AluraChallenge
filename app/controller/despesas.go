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
	db.AutoMigrate(&models.Despesas{})
	return &DespesasRepo{Db: db}
}

// postAlbums adds an album from JSON received in the request body.
func (d *DespesasRepo) CreateDespesa(c *gin.Context) {
	//var validate = validator.New()

	var input models.Despesas

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Valor != 0 || input.Descricao != "" || input.DataAtual != "" {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Faltando valores"})
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

func (*DespesasRepo) UpdateDespesa(c *gin.Context) {

}

func (*DespesasRepo) DeleteDespesa(c *gin.Context) {

}

func (*DespesasRepo) DespesaById(c *gin.Context) {

}
