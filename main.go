package main

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/juliazuin/AluraChallenge/app/controller"
	database "github.com/juliazuin/AluraChallenge/app/dabatase"
	"gorm.io/gorm"
)

func main() {
	db := setupDatabase()
	setupRouter(db).Run(":8080")
}

func setupDatabase() *gorm.DB {
	db := database.NewDB()
	return db.Db
}

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// init categories and seed table
	//controllers.NewCategoria(db).SeedCategorias()

	//Despesas rouutes
	despesaController := controllers.NewDespesa(db)
	apiDespesas := r.Group("/despesas")
	{
		apiDespesas.POST("", despesaController.CreateDespesa)
		apiDespesas.PUT("/:id", despesaController.UpdateDespesa)
		apiDespesas.DELETE("/:id", despesaController.DeleteDespesa)
		apiDespesas.GET("/:id", despesaController.DespesaById)
		apiDespesas.GET("", despesaController.ListDespesas)
	}

	r.GET("despesa/:year/:month", despesaController.ListDespesaByMonth)

	//Receitas Routes
	receitaController := controllers.NewReceita(db)
	apiReceitas := r.Group("/receitas")
	{
		apiReceitas.POST("", receitaController.CreateReceita)
		apiReceitas.PUT("/:id", receitaController.UpdateReceita)
		apiReceitas.DELETE("/:id", receitaController.DeleteReceita)
		apiReceitas.GET("/:id", receitaController.ReceitaById)
		apiReceitas.GET("", receitaController.ListReceitas)
	}

	r.GET("/receita/:year/:month", receitaController.ListReceitaByMonth)

	// resumo
	resumoController := controllers.NewResumo(db)
	r.GET("/resumo/:year/:month", resumoController.ListResumoByMonth)

	return r
}
