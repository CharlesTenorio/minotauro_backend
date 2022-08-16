package domain

import (
	"errors"
	"time"
)

type TipoProfissional struct {
	IdTipoProf string    `json:"id_tipo_prof"`
	Nome       string    `json:"name"`
	DataAt     time.Time `json:"data"`
}

func (t *TipoProfissional) validade() error {
	if t.Nome == "" {
		return errors.New("O Nome o brigatório")
	}

	return nil

}
