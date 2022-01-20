package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controllers "github.com/juliazuin/AluraChallenge/app/controller"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	userRepo := controllers.New()
	r.POST("/despesas", userRepo.PostDespesa)

	return r
}
