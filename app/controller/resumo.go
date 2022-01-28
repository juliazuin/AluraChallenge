package controller

import (
	"github.com/gin-gonic/gin"
	models "github.com/juliazuin/AluraChallenge/app/model"
	"gorm.io/gorm"
)

type ResumoController struct {
	Db *gorm.DB
}

func NewResumo(db *gorm.DB) *ResumoController {
	return &ResumoController{Db: db}
}

func (resumo *ResumoController) ListResumoByMonth(c *gin.Context) {

	receitaYear := c.Param("year")
	receitaMonth := c.Param("month")

	resumo.getTotalReceitas(receitaMonth, receitaYear)

}

func (resumo *ResumoController) getTotalReceitas(month string, year string) float64 {
	receitas := []models.Receita{}
	result := resumo.Db.Select("valor").Where("data_atual LIKE ?", "%"+month+"/"+year+"%").Find(&receitas)
	var somaValores float64
	if result.RowsAffected > 0 {
		for _, valores := range receitas {
			somaValores += valores.Valor
		}
		print(somaValores)
		return somaValores
	} else {
		return somaValores
	}
}

func getTotalDespesas() {
	/*receitas := []models.Receita{}
	result := resumo.Db.Select("valor").Where("data_atual LIKE ?", "%"+month+"/"+year+"%").Find(&receitas)

	if result.RowsAffected > 0 {
		var somaValores float64
		for _, valores := range receitas {
			somaValores += valores.Valor
		}
		print(somaValores)
	} else {
		print(result)
	}*/
}

func getTotal() {

}

func getTotalByCategory() {

}
