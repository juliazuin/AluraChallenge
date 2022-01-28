package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/juliazuin/AluraChallenge/app/controller"
)

type ReceitaRoute struct {
	receitaController *controller.ReceitasController
}

func (r *ReceitaRoute) Setup() {
	routes := gin.Default()

	routes.POST("/receitas", r.receitaController.CreateReceita)
	routes.PUT("/'receitas/:id'", r.receitaController.UpdateReceita)
	routes.DELETE("/receitas/:id", r.receitaController.DeleteReceita)
	routes.GET("/receitas/:id", r.receitaController.ReceitaById)
	routes.GET("/receitas", r.receitaController.ListReceitas)
}

func NewReceitaRoute(receitaController *controller.ReceitasController) *ReceitaRoute {
	return &ReceitaRoute{
		receitaController: receitaController,
	}
}
