package repositorios

import (
	uuid "github.com/satori/go.uuid"
	d "minotauro/app/domain"
	"time"
)

type CorridaRepository interface {
	GetById(id string) (d.Corrida, error)
	Create(Corrida d.Corrida) (d.Corrida, error)
	Update(Corrida d.Corrida) (d.Corrida, error)
	Delete(id string) (string, error)
	FindAll() ([]d.Corrida, error)
}

func NewCyclist() *d.Corrida {
	corrida := d.Corrida{
		IdCorrida: uuid.NewV4().String(),
		CreateAt:  time.Now(),
		Active:    true,
	}

	return &corrida
}
