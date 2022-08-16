package domain

import (
	"time"
)

type Campeonato struct {
	IdCamponato      string    `json:"id_campeonato"`
	Nome             string    `json:"nome"`
	DataInicial      time.Time `json:"data inicial"`
	DataFinal        time.Time `json:"data final"`
	ValPrimeiroLugar float64   `json:"valor primeiro lugar"`
	ValSegundoLugar  float64   `json:"valor segundo lugar"`
	ValTerceiroLugar float64   `json:"valor terceiro lugar"`
	Cartaz           string    `json:"Cataz"`
	Ativo            bool      `json:"ativo"`
	DataAt           time.Time `json:"data criacao"`
}
