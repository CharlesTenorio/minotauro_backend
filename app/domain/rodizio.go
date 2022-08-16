package domain

import (
	"errors"

	"time"

	uuid "github.com/satori/go.uuid"
)

type Rodizio struct {
	IdRodizio string     `json:"id_rodizio"`
	Numero    uint64     `json:"numero"`
	IdEvento  string     `json:"codigo do evento"`
	Corrida   []*Corrida `gorm:"foreignKey:Corrida"`
	Active    bool       `json:"activate"`
	CreateAt  time.Time  `json:"data"`
}

type RodizioRepository interface {
	GetById(id string) (Rodizio, error)
	Create(Rodizio Rodizio) (Rodizio, error)
	Update(Rodizio Rodizio) (Rodizio, error)
	Delete(id string) (string, error)
	FindAll() ([]Rodizio, error)
}

func NewRodizio() *Rodizio {
	rodizio := Rodizio{
		IdRodizio: uuid.NewV4().String(),
		CreateAt:  time.Now(),
		Active:    true,
	}

	return &rodizio
}

func (r *Rodizio) validade() error {
	if r.Numero == 0 {
		return errors.New("O Numero o brigatório")
	}

	if r.IdEvento == "" {
		return errors.New("o código do evento é o brigatório")
	}

	return nil

}
