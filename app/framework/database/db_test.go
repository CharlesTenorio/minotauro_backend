package database_test

import (
	"minotauro/app/framework/database"
	"testing"
)

func TestConexao(t *testing.T) {
	db, err := database.Conn()
	if err != nil {
		t.Error("ERRO AO CONECTAR", err.Error())
	}

	defer db.Close()
}
