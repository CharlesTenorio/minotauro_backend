package repositorios

import (
	uuid "github.com/satori/go.uuid"
	d "minotauro/app/domain"
	"time"
)

type TipoProfissionalRepository interface {
	GetById(id string) (d.TipoProfissional, error)
	Create(TipoProfissional d.TipoProfissional) (d.TipoProfissional, error)
	Update(TipoProfissional d.TipoProfissional) (d.TipoProfissional, error)
	Delete(id string) (string, error)
	FindAll() ([]d.TipoProfissional, error)
}

func NewTipoProfissional() *d.TipoProfissional {
	tipoProfissional := d.TipoProfissional{
		IdTipoProf: uuid.NewV4().String(),
		DataAt:     time.Now(),
	}

	return &tipoProfissional
}
