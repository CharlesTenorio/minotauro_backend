package domain

import (
	"errors"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Professional struct {
	Id                 string    `json:"id"`
	Name               string    `json:"name"`
	TypeProfessionalID string    `json:"name"`
	Sex                string    `json:"sex"`
	Birth              time.Time `json:"birth"`
	Cellphone          string    `json:"celular"`
	Active             bool      `json:"activate"`
	img                string    `json:"img"`
	CreateAt           time.Time `json:"data"`
}

type ProfessionalRepository interface {
	GetById(id string) (Professional, error)
	Create(Professional Professional) (Professional, error)
	Update(Professional Professional) (Professional, error)
	Delete(id string) (string, error)
	FindAll() ([]Professional, error)
}

func NewProfessional() *Professional {
	professional := Professional{
		Id:       uuid.NewV4().String(),
		CreateAt: time.Now(),
		Active:   true,
	}

	return &professional
}

func (p *Professional) validade() error {
	if p.Name == "" {
		return errors.New("O Nome o brigatório")
	}

	if p.Sex == "" {
		return errors.New("O sexo o brigatório")
	}

	if p.Birth.IsZero() {
		return errors.New(" Nascimento é o brigatório")
	}

	if p.Cellphone == "" {
		return errors.New("Celular não informado")
	}
	return nil

}

func (professional *Professional) formatSpace() {
	professional.Name = strings.TrimSpace(professional.Name)

}
