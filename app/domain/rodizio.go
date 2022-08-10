package domain

import (
	"errors"

	"time"

	uuid "github.com/satori/go.uuid"
)

type Rotation struct {
	Id       string    `json:"id"`
	Number   uint64    `json:"numero"`
	IdEvent  string    `json:"codigo do evento"`
	Active   bool      `json:"activate"`
	CreateAt time.Time `json:"data"`
}

type RotationRepository interface {
	GetById(id string) (Rotation, error)
	Create(Rotation Rotation) (Rotation, error)
	Update(Rotation Rotation) (Rotation, error)
	Delete(id string) (string, error)
	FindAll() ([]Rotation, error)
}

func NewRotation() *Rotation {
	rotation := Rotation{
		Id:       uuid.NewV4().String(),
		CreateAt: time.Now(),
		Active:   true,
	}

	return &rotation
}

func (r *Rotation) validade() error {
	if r.Number == 0 {
		return errors.New("O Numero o brigatório")
	}

	if r.IdEvent == "" {
		return errors.New("o código do evento é o brigatório")
	}

	return nil

}
