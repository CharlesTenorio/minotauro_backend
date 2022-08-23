package domain

import (
	"errors"
	"strings"
	"time"
)

type TipoProfissional struct {
	IdTipoProf string    `json:"id_tipo_prof"`
	Nome       string    `json:"nome"`
	DataAt     time.Time `json:"data"`
}

func (t *TipoProfissional) validade() error {
	if t.Nome == "" {
		return errors.New("O Nome o brigatório")
	}

	return nil

}

func (tipoProfissional *TipoProfissional) Prepere() error {
	if erro := tipoProfissional.validade(); erro != nil {
		return erro
	}
	tipoProfissional.formatSpace()
	return nil
}

func (tipoProfissional *TipoProfissional) formatSpace() {
	tipoProfissional.Nome = strings.TrimSpace(tipoProfissional.Nome)
}
