package domain

import (
	"errors"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Cowboy struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Sex       string    `json:"sex"`
	Birth     time.Time `json:"birth"`
	Cellphone string    `json:"celular"`
	Active    bool      `json:"activate"`
	img       string    `json:"img"`
	CreateAt  time.Time `json:"data"`
}

type CowboyRepository interface {
	GetById(id string) (Cowboy, error)
	Create(Cowboy Cowboy) (Cowboy, error)
	Update(Cowboy Cowboy) (Cowboy, error)
	Delete(id string) (string, error)
	FindAll() ([]Cowboy, error)
}

func NewCyclist() *Cowboy {
	cowboy := Cowboy{
		Id:       uuid.NewV4().String(),
		CreateAt: time.Now(),
		Active:   true,
	}

	return &cowboy
}

func (c *Cowboy) validade() error {
	if c.Name == "" {
		return errors.New("O Nome o brigatório")
	}

	if c.Sex == "" {
		return errors.New("O sexo o brigatório")
	}

	if c.Birth.IsZero() {
		return errors.New(" Nascimento é o brigatório")
	}

	if c.Cellphone == "" {
		return errors.New("Celular não informado")
	}
	return nil

}

func (cowboy *Cowboy) formatSpace() {
	cowboy.Name = strings.TrimSpace(cowboy.Name)

}
