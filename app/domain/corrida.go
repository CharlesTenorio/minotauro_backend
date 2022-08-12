package domain

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Corrida struct {
	gorm.Model
	IdCorrida    string    `json:"id_corrida" gorm:"type:uuid;primary_key"`
	IdRodizio    string    `json:"codiog do Rodizio"`
	IdVaqueiro   string    `json:"codigo do vaqueiro"`
	IdCategoria  string    `json:"celular"`
	PrimeiroBoi  string    `json:"1 Boi"`
	SegundoBoi   string    `json:"2 Boi"`
	TerceirooBoi string    `json:"3 Boi"`
	QuartoBoi    string    `json:"4 Boi"`
	QuintoBoi    string    `json:"5 Boi"`
	SextoBoi     string    `json:"6 Boi"`
	SetimoBoi    string    `json:"7 Boi"`
	OitavoBoi    string    `json:"8 Boi"`
	Active       bool      `json:"activate"`
	Img          string    `json:"imgagem"`
	CreateAt     time.Time `json:"data"`
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
		IdCorrida: uuid.NewV4().String(),
		CreateAt:  time.Now(),
		Active:    true,
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
