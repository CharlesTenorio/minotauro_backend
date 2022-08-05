package domain

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type TypeProfessional struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"data"`
}

type TypeProfessionalRepository interface {
	GetById(id string) (TypeProfessional, error)
	Create(TypeProfessional TypeProfessional) (TypeProfessional, error)
	Update(TypeProfessional TypeProfessional) (TypeProfessional, error)
	Delete(id string) (string, error)
	FindAll() ([]TypeProfessional, error)
}

func NewTypeProfessional() *TypeProfessional {
	typeProfessional := TypeProfessional{
		Id:       uuid.NewV4().String(),
		CreateAt: time.Now(),
	}

	return &typeProfessional
}

func (t *TypeProfessional) validade() error {
	if t.Name == "" {
		return errors.New("O Nome o brigatório")
	}

	return nil

}
