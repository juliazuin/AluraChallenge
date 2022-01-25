package main

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/juliazuin/AluraChallenge/app/controller"
)

func main() {
	setupRouter().Run(":8080")
	createCategories()
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// init categories and seed table
	controllers.NewCategoria().SeedCategorias()
	
	//Despesas rouutes
	despesaController := controllers.NewDespesa()
	r.POST("/despesas", despesaController.CreateDespesa)
	r.PUT("/despesas/:id", despesaController.UpdateDespesa)
	r.DELETE("/despesas/:id", despesaController.DeleteDespesa)
	r.GET("/despesas/:id", despesaController.DespesaById)
	r.GET("/despesas", despesaController.ListDespesas)

	//Receitas Routes
	receitaController := controllers.NewReceita()
	r.POST("/receitas", receitaController.CreateReceita)
	r.PUT("/receitas/:id", receitaController.UpdateReceita)
	r.DELETE("/receitas/:id", receitaController.DeleteReceita)
	r.GET("/receitas/:id", receitaController.ReceitaById)
	r.GET("/receitas", receitaController.ListReceitas)

	return r
}

func createCategories() {

}
