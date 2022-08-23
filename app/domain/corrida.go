package domain

import (
	"errors"
	"time"
)

type Corrida struct {
	IdCorrida    string    `json:"id_corrida"`
	IdRodizio    string    `json:"codiog do Rodizio"`
	IdVaqueiro   string    `json:"codigo do vaqueiro"`
	IdCategoria  string    `json:"celular"`
	PrimeiroBoi  string    `json:"1 Boi"`
	SegundoBoi   string    `json:"2 Boi"`
	TerceirooBoi string    `json:"3 Boi"`
	QuartoBoi    string    `json:"4 Boi"`
	QuintoBoi    string    `json:"5 Boi"`
	SextoBoi     string    `json:"6 Boi"`
	SetimoBoi    string    `json:"7 Boi"`
	OitavoBoi    string    `json:"8 Boi"`
	Active       bool      `json:"activate"`
	Img          string    `json:"imgagem"`
	CreateAt     time.Time `json:"data"`
}

func (c *Corrida) validade() error {
	if c.IdCategoria == "" {
		return errors.New("O Nome o brigatório")
	}

	if c.IdRodizio == "" {
		return errors.New("O sexo o brigatório")
	}

	if c.IdVaqueiro == "" {
		return errors.New("Celular não informado")
	}
	return nil
}
