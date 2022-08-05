package domain

import (
	"errors"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Park struct {
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	Address        string    `json:"endereço"`
	Neighborhood   string    `json:"Bairro"`
	City           string    `json:"cidade"`
	Uf             string    `json:"uf"`
	Active         bool      `json:"activate"`
	Img            string    `json:"img"`
	Referencepoint string    `json:"reference point"`
	CreateAt       time.Time `json:"data"`
}

type ParkRepository interface {
	GetById(id string) (Park, error)
	Create(Park Park) (Park, error)
	Update(Park Park) (Park, error)
	Delete(id string) (string, error)
	FindAll() ([]Park, error)
}

func NewPark() *Park {
	park := Park{
		Id:       uuid.NewV4().String(),
		CreateAt: time.Now(),
		Active:   true,
	}

	return &park
}

func (p *Park) validade() error {
	if p.Name == "" {
		return errors.New("O Nome o brigatório")
	}

	if p.Address == "" {
		return errors.New("Endereço e obrigatório")
	}

	if p.Neighborhood == "" {
		return errors.New("Bairro obriatório")
	}

	if p.City == "" {
		return errors.New("Cidade é obriatório")
	}

	if p.Uf == "" {
		return errors.New("UF obriatório")
	}

	if p.Referencepoint == "" {
		return errors.New("Ponto de referencia e obritatório")
	}
	return nil

}

func (park *Park) formatSpace() {
	park.Name = strings.TrimSpace(park.Name)
	park.Address = strings.TrimSpace(park.Address)
	park.Neighborhood = strings.TrimSpace(park.Neighborhood)
	park.City = strings.TrimSpace(park.City)
	park.Referencepoint = strings.TrimSpace(park.Referencepoint)

}
