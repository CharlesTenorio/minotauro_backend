package domain

import (
	"errors"
	"strings"
	"time"
)

type Profissional struct {
	IdProfissional string    `json:"id_profissional"`
	Nome           string    `json:"name"`
	IdTipoProf     string    `json:"Tipo do profissional"`
	Sexo           string    `json:"sex"`
	Nascimento     time.Time `json:"birth"`
	Celular        string    `json:"celular"`
	Activado       bool      `json:"activate"`
	Img            string    `json:"img"`
	DataAt         time.Time `json:"data"`
}

func (profissional *Profissional) Prepere() error {
	if erro := profissional.validade(); erro != nil {
		return erro
	}
	profissional.formatSpace()
	return nil
}

func (p *Profissional) validade() error {
	if p.Nome == "" {
		return errors.New("O Nome o brigatorio")
	}

	if p.Sexo == "" {
		return errors.New("O sexo o brigatório")
	}

	if p.Nascimento.IsZero() {
		return errors.New(" Nascimento é o brigatório")
	}

	if p.Celular == "" {
		return errors.New("Celular não informado")
	}

	if p.IdTipoProf == "" {
		return errors.New("O Tipo do profissiona é obrigatório")
	}

	return nil

}

func (profissional *Profissional) formatSpace() {
	profissional.Nome = strings.TrimSpace(profissional.Nome)
}
