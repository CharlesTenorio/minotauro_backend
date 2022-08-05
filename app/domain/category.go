package domain

import (
	"errors"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Category struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"valor"`
	Active   bool      `json:"activate"`
	CreateAt time.Time `json:"data"`
}

type CategoryRepository interface {
	GetById(id string) (Category, error)
	Create(Category Category) (Category, error)
	Update(Category Category) (Category, error)
	Delete(id string) (string, error)
	FindAll() ([]Category, error)
}

func NewCategory() *Category {
	category := Category{
		Id:       uuid.NewV4().String(),
		CreateAt: time.Now(),
		Active:   true,
	}

	return &category
}

func (c *Category) validade() error {
	if c.Name == "" {
		return errors.New("O Nome o brigatório")
	}

	return nil

}

func (payment *Payment) formatSpace() {
	payment.Name = strings.TrimSpace(payment.Name)

}
