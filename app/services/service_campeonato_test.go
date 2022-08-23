package services_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"minotauro/app/aplicacao/repositorios"
	"minotauro/app/framework/database"
	"minotauro/app/services"
	"testing"
	"time"
)

func TestCampeonatoService_Create(t *testing.T) {
	db, err := database.Conn()
	defer db.Close()
	repositorySql := repositorios.NewCampeonatoDb(db)
	serv := services.NewCamponatoServece(repositorySql)
	agora := time.Now()
	nome := "Pega Garanhus"
	datIni := agora.AddDate(0, 1, 0)
	datFim := agora.AddDate(0, 7, 0)
	vp := 50.000
	vs := 30.000
	vt := 20.000
	vq := 10.000
	cartaz := "img"
	c, err := serv.Create(nome, datIni, datFim, vp, vs, vt, vq, cartaz)
	if err != nil {
		fmt.Println("erro de inserte", err)
	}
	fmt.Println(c.IdCamponato)
	require.Nil(t, err)
	require.Nil(t, err, nil)
}
