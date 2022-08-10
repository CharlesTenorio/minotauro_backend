package domain

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type bois struct {
	numero    uint
	maracacao string
}

type Corrida struct {
	Id          string `json:"id"`
	IdRodizio   string `json:"codiog do Rodizio"`
	IdCorrida   string `json:"Codigo do vaqueiro"`
	IdVaqueiro  string `json:"codigo do vaqueiro"`
	IdCategoria string `json:"celular"`
	Bois        map[string]*bois
	Active      bool      `json:"activate"`
	Img         string    `json:"imgagem"`
	CreateAt    time.Time `json:"data"`
}

type CorridaRepository interface {
	GetById(id string) (Corrida, error)
	Create(Corrida Corrida) (Corrida, error)
	Update(Corrida Corrida) (Corrida, error)
	Delete(id string) (string, error)
	FindAll() ([]Corrida, error)
}

func NewCyclist() *Corrida {
	corrida := Corrida{
		Id:       uuid.NewV4().String(),
		CreateAt: time.Now(),
		Active:   true,
	}

	return &corrida
}

func (c *Corrida) validade() error {
	if c.IdCategoria == "" {
		return errors.New("O Nome o brigatório")
	}

	if c.IdRodizio == "" {
		return errors.New("O sexo o brigatório")
	}

	if c.IdVaqueiro == "" {
		return errors.New("Celular não informado")
	}
	return nil

}