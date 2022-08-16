package repositorios

import (
	uuid "github.com/satori/go.uuid"
	d "minotauro/app/domain"
	"time"
)

type CampeonatoRepository interface {
	GetById(id string) (d.Campeonato, error)
	Create(d.Campeonato) (d.Campeonato, error)
	Update(d.Campeonato) (d.Campeonato, error)
	Delete(id string) (string, error)
	FindAll() ([]d.Campeonato, error)
}

func NewCampeonato() *d.Campeonato {
	campeonato := d.Campeonato{
		IdCamponato: uuid.NewV4().String(),
		DataAt:      time.Now(),
		Ativo:       true,
	}

	return &campeonato
}
