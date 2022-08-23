package domain

import (
	"errors"
	"strings"

	"time"
)

type Evento struct {
	IdEvento        string    `json:"id_evento"`
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

func (evento *Evento) Prepere() error {
	if erro := evento.validade(); erro != nil {
		return erro
	}
	evento.formatSpace()
	return nil
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
func (evento *Evento) formatSpace() {
	evento.Nome = strings.TrimSpace(evento.Nome)
}
