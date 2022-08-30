package api

import (
	"github.com/gin-gonic/gin"
	"minotauro/app/aplicacao/repositorios"
	"minotauro/app/framework/database"
	"minotauro/app/services"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func CriarCampeonato(c *gin.Context) {
	db, err := database.Conn()
	defer db.Close()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	repositorySql := repositorios.NewCampeonatoDb(db)
	serv := services.NewCamponatoServece(repositorySql)
	nome := c.PostForm("nome")
	datIni := c.PostForm("dataInical")
	datFim := c.PostForm("dataFim")
	valPrimeioPremioStr := c.PostForm("valor_primeioro_lugar")
	valSegundoLugarStr := c.PostForm("Valor_segundo_lugar")
	valoTerceiroLugarStr := c.PostForm("valor_terceiro_lugar")
	valorQuartoLugarStr := c.PostForm("valor_quatrto_lugar")
	cartaz := "img"

	vp, err := strconv.ParseFloat(valPrimeioPremioStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vs, err := strconv.ParseFloat(valSegundoLugarStr, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vt, err := strconv.ParseFloat(valoTerceiroLugarStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	vq, err := strconv.ParseFloat(valorQuartoLugarStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data_ini, err := time.Parse("2006-01-02", datIni)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data_fim, err := time.Parse("2006-01-02", datFim)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	nome = strings.ToUpper(nome)
	cp, err := serv.Create(nome, data_ini, data_fim, vp, vs, vt, vq, cartaz)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"compeonato": cp})
	return
}

func PegarCampeonato(c *gin.Context) {
	id := c.Param("id")
	db, err := database.Conn()
	defer db.Close()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	repositorySql := repositorios.NewCampeonatoDb(db)
	serv := services.NewCamponatoServece(repositorySql)

	cp, err := serv.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"compeonato": cp})
	return

}

func ListarCampeonato(c *gin.Context) {

	db, err := database.Conn()
	defer db.Close()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	repositorySql := repositorios.NewCampeonatoDb(db)
	serv := services.NewCamponatoServece(repositorySql)
	cp, err := serv.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"compeonatos": cp})
	return

}
