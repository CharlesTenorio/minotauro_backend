package domain

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Categoria struct {
	gorm.Model
	IdCategoria string    `json:"id_categoria" gorm:"type:uuid;primary_key"`
	IdParque    string    `json:"id_parque" gorm:"column:id_parque;type:uuid;notnull"`
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

func (c *Categoria) validade() error {
	if c.Nome == "" {
		return errors.New("O Nome o brigatório")
	}

	return nil

}
