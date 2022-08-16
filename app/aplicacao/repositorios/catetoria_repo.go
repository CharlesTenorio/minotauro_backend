package repositorios

import (
	uuid "github.com/satori/go.uuid"
	d "minotauro/app/domain"
	"time"
)

type CategoriaRepository interface {
	GetById(id string) (d.Categoria, error)
	Create(d.Categoria) (d.Categoria, error)
	Update(d.Categoria) (d.Categoria, error)
	Delete(id string) (string, error)
	FindAll() ([]d.Categoria, error)
}

func NewCategoria() *d.Categoria {
	categoria := d.Categoria{
		IdCategoria: uuid.NewV4().String(),
		DataAt:      time.Now(),
		Ativado:     true,
	}

	return &categoria
}
