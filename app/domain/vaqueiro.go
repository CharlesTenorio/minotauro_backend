package domain

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Vaqueiro struct {
	Id         string    `json:"id"`
	Nome       string    `json:"name"`
	Sexo       string    `json:"sex"`
	Nascimento time.Time `json:"birth"`
	Celular    string    `json:"celular"`
	Ativo      bool      `json:"activate"`
	Img        string    `json:"imgagem"`
	DataAt     time.Time `json:"data"`
}

type VaqueiroRepository interface {
	GetById(id string) (Vaqueiro, error)
	Create(Vaqueiro Vaqueiro) (Vaqueiro, error)
	Update(Vaqueiro Vaqueiro) (Vaqueiro, error)
	Delete(id string) (string, error)
	FindAll() ([]Vaqueiro, error)
}

func NewVaqueiro() *Vaqueiro {
	vaqueiro := Vaqueiro{
		Id:     uuid.NewV4().String(),
		DataAt: time.Now(),
		Ativo:  true,
	}

	return &vaqueiro
}

func (c *Vaqueiro) validade() error {
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
