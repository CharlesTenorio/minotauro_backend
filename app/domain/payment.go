package domain

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Payment struct {
	Id       string    `json:"id"`
	IdPark   string    `json:"idPark"`
	Name     string    `json:"name"`
	Active   bool      `json:"activate"`
	CreateAt time.Time `json:"data"`
}

type PaymentRepository interface {
	GetById(id string) (Payment, error)
	Create(Payment Payment) (Payment, error)
	Update(Payment Payment) (Payment, error)
	Delete(id string) (string, error)
	FindAll() ([]Payment, error)
}

func NewPayment() *Payment {
	payment := Payment{
		Id:       uuid.NewV4().String(),
		IdPark:   "",
		Name:     "",
		Active:   true,
		CreateAt: time.Now(),
	}

	return &payment
}

func (p *Payment) validade() error {
	if p.Name == "" {
		return errors.New("O Nome o brigatório")
	}

	if p.IdPark == "" {
		return errors.New("O codigo do parque é obrigatório")
	}

	return nil

}
