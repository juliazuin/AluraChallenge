package routes

type Routes []Route

type Route interface {
	Setup()
}

func (r *Routes) Setup() {
	for _, route := range *r {
		route.Setup()
	}
}
