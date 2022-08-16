package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Categoria struct {
	IdCategoria string    `json:"id_categoria"`
	IdParque    string    `json:"id_parque"`
	Nome        string    `json:"nome"`
	Valor       float64   `json:"valor"`
	Ativado     bool      `json:"activate"`
	DataAt      time.Time `json:"data"`
}

type CategoriaRepository interface {
	GetById(id string) (Categoria, error)
	Create(Categoria) (Categoria, error)
	Update(Categoria) (Categoria, error)
	Delete(id string) (string, error)
	FindAll() ([]Categoria, error)
}

func NewCategoria() *Categoria {
	categoria := Categoria{
		IdCategoria: uuid.NewV4().String(),
		DataAt:      time.Now(),
		Ativado:     true,
	}

	return &categoria
}
