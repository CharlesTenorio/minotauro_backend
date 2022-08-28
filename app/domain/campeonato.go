package domain

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Campeonato struct {
	IdCamponato      string    `json:"id_campeonato"`
	Nome             string    `json:"Nome"`
	DataInicial      time.Time `json:"data inicial"`
	DataFinal        time.Time `json:"data final"`
	ValPrimeiroLugar float64   `json:"valor primeiro lugar"`
	ValSegundoLugar  float64   `json:"valor segundo lugar"`
	ValTerceiroLugar float64   `json:"valor terceiro lugar"`
	ValQuartoLugar   float64   `json:"valor terceiro lugar"`
	Cartaz           string    `json:"Cataz"`
	Ativo            bool      `json:"ativo"`
	DataAt           time.Time `json:"data criacao"`
	SoftDelete       bool      `json:"soft_delete"`
}

func NovoCampeonato() *Campeonato {
	camp := Campeonato{
		IdCamponato: uuid.NewV4().String(),
		DataAt:      time.Now(),
		Ativo:       true,
		SoftDelete:  true,
	}
	return &camp
}
