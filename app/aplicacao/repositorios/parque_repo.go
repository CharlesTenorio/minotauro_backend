package repositorios

import (
	uuid "github.com/satori/go.uuid"
	d "minotauro/app/domain"
	"time"
)

type ParqueRepository interface {
	GetById(id string) (d.Parque, error)
	Create(Parque d.Parque) (d.Parque, error)
	Update(Parque d.Parque) (d.Parque, error)
	Delete(id string) (string, error)
	FindAll() ([]d.Parque, error)
}

func NewParque() *d.Parque {
	parque := d.Parque{
		IdParque: uuid.NewV4().String(),
		DataAt:   time.Now(),
		Ativado:  true,
	}

	return &parque
}
