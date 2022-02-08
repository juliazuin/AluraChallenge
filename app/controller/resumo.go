package controller

import (
	"net/http"

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
	resumos := models.Resumo{}

	receitaYear := c.Param("year")
	receitaMonth := c.Param("month")

	resumoReceitas := resumo.getTotalReceitas(receitaMonth, receitaYear)
	resumoDespesas := resumo.getTotalDespesas(receitaMonth, receitaYear)

	totalMonth := resumoReceitas - resumoDespesas

	despesaByCategory := resumo.getTotalByCategory()

	resumos.TotalReceitasMes = resumoReceitas
	resumos.TotalDespesasMes = resumoReceitas
	resumos.SaldoFinalMes = totalMonth
	resumos.DespesaByCategory = despesaByCategory

	c.JSON(http.StatusOK, gin.H{"data": &resumos})
}

func (resumo *ResumoController) getTotalReceitas(month string, year string) float64 {
	var receitas []models.Receita
	result := resumo.Db.Select("valor").Where("data_atual LIKE ?", "%"+month+"/"+year+"%").Find(&receitas)
	var somaValores float64
	if result.RowsAffected > 0 {
		for _, valores := range receitas {
			somaValores += valores.Valor
		}
		return somaValores
	} else {
		return somaValores
	}
}

func (resumo *ResumoController) getTotalDespesas(month string, year string) float64 {
	var despesas []models.Receita
	result := resumo.Db.Select("valor").Where("data_atual LIKE ?", "%"+month+"/"+year+"%").Find(&despesas)
	var somaValores float64
	if result.RowsAffected > 0 {
		for _, valores := range despesas {
			somaValores += valores.Valor
		}
		return somaValores
	} else {
		return somaValores
	}
}

func (resumo *ResumoController) getTotalByCategory() *[]models.Despesa {
	var despesasByCategory []models.Despesa

	resumo.Db.Find(&despesasByCategory)
	print(&despesasByCategory)

	return &despesasByCategory
}
