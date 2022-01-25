package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/juliazuin/AluraChallenge/app/controller"
)

type ReceitaRoute struct {
	ReceiraController *controller.DespesasController
}

func (r *ReceitaRoute) Setup() {
	routes := gin.Default()

	routes.POST("/receitas", r.ReceiraController.CreateDespesa)
	routes.PUT("/'receitas'", r.ReceiraController.UpdateDespesa)
	routes.GET("/receitas", r.ReceiraController.DespesaById)
	routes.POST("/receitas", r.ReceiraController.DeleteDespesa)
}
