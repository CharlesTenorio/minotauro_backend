package startapp

import (
	"github.com/gin-gonic/gin"

	"minotauro/app/framework/api"
)

func Start() {
	g := gin.Default()

	g.POST("/criar_campeonato", api.CriarCampeonato)
	g.GET("/pegar_campeonanto/:id", api.PegarCampeonato)
	g.GET("/listar_campeonanto", api.ListarCampeonato)

	g.Run()
}
