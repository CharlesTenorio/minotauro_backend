package domain

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Parque struct {
	Id                string    `json:"id"`
	Nome              string    `json:"name"`
	Endereco          string    `json:"endereço"`
	Bairro            string    `json:"Bairro"`
	Cidade            string    `json:"cidade"`
	Uf                string    `json:"uf"`
	Ativado           bool      `json:"activate"`
	Img               string    `json:"img"`
	PorntoDeRefrencia string    `json:"reference point"`
	DataAt            time.Time `json:"data"`
}

type ParqueRepository interface {
	GetById(id string) (Parque, error)
	Create(Parque Parque) (Parque, error)
	Update(Parque Parque) (Parque, error)
	Delete(id string) (string, error)
	FindAll() ([]Parque, error)
}

func NewParque() *Parque {
	parque := Parque{
		Id:      uuid.NewV4().String(),
		DataAt:  time.Now(),
		Ativado: true,
	}

	return &parque
}

func (p *Parque) validade() error {
	if p.Nome == "" {
		return errors.New("O Nome o brigatório")
	}

	if p.Endereco == "" {
		return errors.New("Endereço e obrigatório")
	}

	if p.Bairro == "" {
		return errors.New("Bairro obriatório")
	}

	if p.Cidade == "" {
		return errors.New("Cidade é obriatório")
	}

	if p.Uf == "" {
		return errors.New("UF obriatório")
	}

	if p.PorntoDeRefrencia == "" {
		return errors.New("Ponto de referencia e obritatório")
	}
	return nil

}
