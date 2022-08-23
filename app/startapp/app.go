package startapp

import (
	"github.com/gin-gonic/gin"
	"minotauro/app/framework/api"
	"net/http"
)

func Start() {
	g := gin.Default()
	g.POST("/criar_campeonato", api.CriarCampeonato)
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	g.Run()
}
