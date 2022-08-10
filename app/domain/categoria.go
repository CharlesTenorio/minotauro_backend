package domain

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Categoria struct {
	Id      string    `json:"id"`
	Nome    string    `json:"nome"`
	Valor   float64   `json:"valor"`
	Ativado bool      `json:"activate"`
	DataAt  time.Time `json:"data"`
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
		Id:      uuid.NewV4().String(),
		DataAt:  time.Now(),
		Ativado: true,
	}

	return &categoria
}

func (c *Categoria) validade() error {
	if c.Nome == "" {
		return errors.New("O Nome o brigatório")
	}

	return nil

}
