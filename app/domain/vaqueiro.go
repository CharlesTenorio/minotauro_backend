package domain

import (
	"errors"
	"strings"
	"time"
)

type Vaqueiro struct {
	IdVaqueiro string    `json:"id_vaquerio"`
	Nome       string    `json:"name"`
	Sexo       string    `json:"sex"`
	Nascimento time.Time `json:"birth"`
	Celular    string    `json:"celular"`
	Ativo      bool      `json:"activate"`
	Img        string    `json:"imgagem"`
	DataAt     time.Time `json:"data"`
}

func (vaqueiro *Vaqueiro) Prepere() error {
	if erro := vaqueiro.validade(); erro != nil {
		return erro
	}
	vaqueiro.formatSpace()
	return nil
}

func (c *Vaqueiro) validade() error {
	if c.Nome == "" {
		return errors.New("O Nome o brigatório")
	}

	if c.Sexo == "" {
		return errors.New("O sexo o brigatório")
	}

	if c.Nascimento.IsZero() {
		return errors.New(" Nascimento é o brigatório")
	}

	if c.Celular == "" {
		return errors.New("Celular não informado")
	}
	return nil

}

func (vaqueiro *Vaqueiro) formatSpace() {
	vaqueiro.Nome = strings.TrimSpace(vaqueiro.Nome)
}
