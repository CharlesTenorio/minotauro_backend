package domain

import (
	"errors"

	"time"
)

type Rodizio struct {
	IdRodizio string    `json:"id_rodizio"`
	Numero    uint64    `json:"numero"`
	IdEvento  string    `json:"codigo do evento"`
	Active    bool      `json:"activate"`
	CreateAt  time.Time `json:"data"`
}

func (rodizio *Rodizio) Prepere() error {
	if erro := rodizio.validade(); erro != nil {
		return erro
	}

	return nil
}

func (r *Rodizio) validade() error {
	if r.Numero == 0 {
		return errors.New("O Numero o brigatório")
	}

	if r.IdEvento == "" {
		return errors.New("o código do evento é o brigatório")
	}

	return nil

}
