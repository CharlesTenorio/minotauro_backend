package repositorios

import (
	uuid "github.com/satori/go.uuid"
	d "minotauro/app/domain"
	"time"
)

type ProfissionalRepository interface {
	GetById(id string) (d.Profissional, error)
	Create(Profissional d.Profissional) (d.Profissional, error)
	Update(Profissional d.Profissional) (d.Profissional, error)
	Delete(id string) (string, error)
	FindAll() ([]d.Profissional, error)
}

func NewProfissional() *d.Profissional {
	profissional := d.Profissional{
		IdProfissional: uuid.NewV4().String(),
		DataAt:         time.Now(),
		Activado:       true,
	}

	return &profissional
}
