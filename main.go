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
	controllers.NewCategoria(db).SeedCategorias()

	//Despesas rouutes
	despesaController := controllers.NewDespesa(db)
	r.POST("/despesas", despesaController.CreateDespesa)
	r.PUT("/despesas/:id", despesaController.UpdateDespesa)
	r.DELETE("/despesas/:id", despesaController.DeleteDespesa)
	r.GET("/despesas/:id", despesaController.DespesaById)
	r.GET("/despesas", despesaController.ListDespesas)

	//Receitas Routes
	receitaController := controllers.NewReceita(db)
	r.POST("/receitas", receitaController.CreateReceita)
	r.PUT("/receitas/:id", receitaController.UpdateReceita)
	r.DELETE("/receitas/:id", receitaController.DeleteReceita)
	r.GET("/receitas/:id", receitaController.ReceitaById)
	r.GET("/receitas", receitaController.ListReceitas)

	return r
}
