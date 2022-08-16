package repositorios

import (
	uuid "github.com/satori/go.uuid"
	d "minotauro/app/domain"
	"time"
)

type RodizioRepository interface {
	GetById(id string) (d.Rodizio, error)
	Create(Rodizio d.Rodizio) (d.Rodizio, error)
	Update(Rodizio d.Rodizio) (d.Rodizio, error)
	Delete(id string) (string, error)
	FindAll() ([]d.Rodizio, error)
}

func NewRodizio() *d.Rodizio {
	rodizio := d.Rodizio{
		IdRodizio: uuid.NewV4().String(),
		CreateAt:  time.Now(),
		Active:    true,
	}

	return &rodizio
}
