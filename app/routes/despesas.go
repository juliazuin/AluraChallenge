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
	routes.PUT("/despesas/:id", r.despesaController.UpdateDespesa)
	routes.DELETE("/despesas/:id", r.despesaController.DeleteDespesa)
	routes.GET("/despesas/:id", r.despesaController.DespesaById)
	routes.GET("/despesas", r.despesaController.ListDespesas)
}

func NewDespesaRoute(despesaController *controller.DespesasController) *DespesaRoute {
	return &DespesaRoute{
		despesaController: despesaController,
	}
}
