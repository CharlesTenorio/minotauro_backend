package domain

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type TipoProfissional struct {
	IdTipoProf string    `json:"id_tipo_prof"`
	Nome       string    `json:"name"`
	DataAt     time.Time `json:"data"`
}

type TipoProfissionalRepository interface {
	GetById(id string) (TipoProfissional, error)
	Create(TipoProfissional TipoProfissional) (TipoProfissional, error)
	Update(TipoProfissional TipoProfissional) (TipoProfissional, error)
	Delete(id string) (string, error)
	FindAll() ([]TipoProfissional, error)
}

func NewTipoProfissional() *TipoProfissional {
	tipoProfissional := TipoProfissional{
		IdTipoProf: uuid.NewV4().String(),
		DataAt:     time.Now(),
	}

	return &tipoProfissional
}

func (t *TipoProfissional) validade() error {
	if t.Nome == "" {
		return errors.New("O Nome o brigatório")
	}

	return nil

}
