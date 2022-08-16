package repositorios

import (
	uuid "github.com/satori/go.uuid"
	d "minotauro/app/domain"
	"time"
)

type VaqueiroRepository interface {
	GetById(id string) (d.Vaqueiro, error)
	Create(Vaqueiro d.Vaqueiro) (d.Vaqueiro, error)
	Update(Vaqueiro d.Vaqueiro) (d.Vaqueiro, error)
	Delete(id string) (string, error)
	FindAll() ([]d.Vaqueiro, error)
}

func NewVaqueiro() *d.Vaqueiro {
	vaqueiro := d.Vaqueiro{
		IdVaqueiro: uuid.NewV4().String(),
		DataAt:     time.Now(),
		Ativo:      true,
	}

	return &vaqueiro
}
