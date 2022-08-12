package domain

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Profissionall struct {
	gorm.Model
	IdProfissional string    `json:"id_profissional" gorm:"type:uuid;primary_key"`
	Nome           string    `json:"name"`
	IdTipoProf     string    `json:"Tipo do profissional"`
	Sexo           string    `json:"sex"`
	Nascimento     time.Time `json:"birth"`
	Celular        string    `json:"celular"`
	Activado       bool      `json:"activate"`
	Img            string    `json:"img"`
	DataAt         time.Time `json:"data"`
}

type ProfissionallRepository interface {
	GetById(id string) (Profissionall, error)
	Create(Profissionall Profissionall) (Profissionall, error)
	Update(Profissionall Profissionall) (Profissionall, error)
	Delete(id string) (string, error)
	FindAll() ([]Profissionall, error)
}

func NewProfissionall() *Profissionall {
	profissionall := Profissionall{
		IdProfissional: uuid.NewV4().String(),
		DataAt:         time.Now(),
		Activado:       true,
	}

	return &profissionall
}

func (p *Profissionall) validade() error {
	if p.Nome == "" {
		return errors.New("O Nome o brigatorio")
	}

	if p.Sexo == "" {
		return errors.New("O sexo o brigatório")
	}

	if p.Nascimento.IsZero() {
		return errors.New(" Nascimento é o brigatório")
	}

	if p.Celular == "" {
		return errors.New("Celular não informado")
	}

	if p.IdTipoProf == "" {
		return errors.New("O Tipo do profissiona é obrigatório")
	}

	return nil

}
