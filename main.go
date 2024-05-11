package main

import (
	"net/http"
	"ratings/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupApp()

	r.Run(":8080")
}

func setupApp() *gin.Engine {
	r := gin.Default()

	r.GET("/randomMovie", func(c *gin.Context) {
		c.JSON(http.StatusOK, controllers.RandomMovie())
	})

	return r
}
