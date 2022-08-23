package domain

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"strings"
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
	ValQuartoLugar   float64   `json:"valor terceiro lugar"`
	Cartaz           string    `json:"Cataz"`
	Ativo            bool      `json:"ativo"`
	DataAt           time.Time `json:"data criacao"`
}

func (campeonato *Campeonato) Prepere() error {
	if erro := campeonato.validade(); erro != nil {
		return erro
	}
	campeonato.formatSpace()
	return nil
}

func (c *Campeonato) validade() error {
	if c.Nome == "" {
		return errors.New("O Nome o brigatório")
	}

	if c.DataInicial.IsZero() {
		return errors.New(" Inicio é o brigatório")
	}
	if c.ValPrimeiroLugar > 0 {
		return errors.New(" Valor do  1 lugar requerido")
	}

	if c.ValSegundoLugar > 0 {
		return errors.New(" Valor do 2 lugar requerido")
	}

	if c.ValTerceiroLugar > 0 {
		return errors.New(" Valor do 3 lugar requerido")
	}

	if c.ValPrimeiroLugar <= c.ValSegundoLugar {
		return errors.New("Valor do primeiro lugar não pdoe ser menor ou igual ao do segundo lugar")
	}

	if c.ValPrimeiroLugar <= c.ValTerceiroLugar {
		return errors.New("Valor do primeiro lugar não pdoe ser menor ou igual do Terceiro")
	}

	if c.ValSegundoLugar <= c.ValTerceiroLugar {
		return errors.New("Valor do segundo lugar não pdoe ser menor ou igual ao do Terceiro")
	}

	return nil

}

func (campeonato *Campeonato) formatSpace() {
	campeonato.Nome = strings.TrimSpace(campeonato.Nome)
}

func NovoCampeonato() *Campeonato {
	camp := Campeonato{
		IdCamponato: uuid.NewV4().String(),
		DataAt:      time.Now(),
		Ativo:       true,
	}

	return &camp
}
