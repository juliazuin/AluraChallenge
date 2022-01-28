package routes

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(receitaRoute *ReceitaRoute, despesaRoute *DespesaRoute) *Routes {
	return &Routes{
		receitaRoute,
		despesaRoute,
	}
}

func (r *Routes) Setup() {
	for _, route := range *r {
		route.Setup()
	}
}
