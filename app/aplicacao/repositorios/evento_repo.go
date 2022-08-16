package repositorios

import (
	uuid "github.com/satori/go.uuid"
	d "minotauro/app/domain"
	"time"
)

type EventoRepository interface {
	GetById(id string) (d.Evento, error)
	Create(Evento d.Evento) (d.Evento, error)
	Update(Evento d.Evento) (d.Evento, error)
	Delete(id string) (string, error)
	FindAll() ([]d.Evento, error)
}

func NewEvent() *d.Evento {
	evento := d.Evento{
		IdEvento: uuid.NewV4().String(),
		DataAt:   time.Now(),
		Ativado:  true,
	}

	return &evento
}
