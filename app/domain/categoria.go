package domain

import (
	"errors"
	"strings"
	"time"
)

type Categoria struct {
	IdCategoria string    `json:"id_categoria"`
	IdParque    string    `json:"id_parque"`
	Nome        string    `json:"nome"`
	Valor       float64   `json:"valor"`
	Ativado     bool      `json:"activate"`
	DataAt      time.Time `json:"data"`
}

func (categoria *Categoria) Prepere() error {
	if erro := categoria.validade(); erro != nil {
		return erro
	}
	categoria.formatSpace()
	return nil
}

func (c *Categoria) validade() error {
	if c.Nome == "" {
		return errors.New("O Nome é o brigatório")
	}

	if c.Valor > 0 {
		return errors.New("O valor e prigatório")
	}

	return nil

}

func (categoria *Categoria) formatSpace() {
	categoria.Nome = strings.TrimSpace(categoria.Nome)

}
