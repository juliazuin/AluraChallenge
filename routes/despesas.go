package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/juliazuin/AluraChallenge/app/controller"
)

type DespesaRoute struct {
	despesaController *controller.DespesasController
}

func (r *DespesaRoute) Setup() {
	routes := gin.Default()

	routes.POST("/despesas", r.despesaController.CreateDespesa)
	routes.GET("/despesas", r.despesaController.ListDespesas)
	routes.PUT("/despesas", r.despesaController.UpdateDespesa)
	routes.POST("/despesas", r.despesaController.CreateDespesa)
}
