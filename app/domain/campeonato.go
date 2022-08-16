package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Campeonato struct {
	IdCamponato      string    `json:"id_campeonato"`
	Nome             string    `json:"nome"`
	DataInicial      time.Time `json:"data inicial"`
	DataFinal        time.Time `json:"data final"`
	ValPrimeiroLugar float64   `json:"valor primeiro lugar"`
	ValSegundoLugar  float64   `json:"valor segundo lugar"`
	ValTerceiroLugar float64   `json:"valor terceiro lugar"`
	Cartaz           string    `json:"Catraz"`
	Ativo            bool      `json:"ativo"`
	DataAt           time.Time `json:"data criacao"`
}

type CampeonatoRepository interface {
	GetById(id string) (Campeonato, error)
	Create(Campeonato) (Campeonato, error)
	Update(Campeonato) (Campeonato, error)
	Delete(id string) (string, error)
	FindAll() ([]Campeonato, error)
}

func NewCampeonato() *Campeonato {
	campeonato := Campeonato{
		IdCamponato: uuid.NewV4().String(),
		DataAt:      time.Now(),
		Ativo:       true,
	}

	return &campeonato
}
