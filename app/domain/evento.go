package domain

import (
	"errors"

	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Evento struct {
	gorm.Model
	IdEvento        string    `json:"id_evento" gorm:"type:uuid;primary_key"`
	IdParque        string    `json:"idPark"`
	Nome            string    `json:"name"`
	Litime_max_bois uint16    `json:"maximo de bois derrubados"`
	QtdRodadaExtra  uint16    `json:"quatidade de rodas extras"`
	DataDoEvento    time.Time `json:"Data do evento"`
	StatusEvento    string    `json:"status do evento"`
	VaorBoiTv       float64   `json:"valor do boi tv"`
	Ativado         bool      `json:"activate"`
	ImgEvento       string    `json:"img"`
	DataAt          time.Time `json:"data"`
}

type EventRepository interface {
	GetById(id string) (Evento, error)
	Create(Evento Evento) (Evento, error)
	Update(Evento Evento) (Evento, error)
	Delete(id string) (string, error)
	FindAll() ([]Evento, error)
}

func NewEvent() *Evento {
	evento := Evento{
		IdEvento: uuid.NewV4().String(),
		DataAt:   time.Now(),
		Ativado:  true,
	}

	return &evento
}

func (e *Evento) validade() error {
	if e.Nome == "" {
		return errors.New("O Nome o brigatório")
	}

	if e.DataDoEvento.IsZero() {
		return errors.New("Data do evento é o brigatória")
	}

	if e.IdParque == "" {
		return errors.New("O codigo do Parque é obrigatório")
	}

	return nil

}
