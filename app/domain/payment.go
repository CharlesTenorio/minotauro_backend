package domain

import (
	"errors"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Payment struct {
	Id       string    `json:"id"`
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
		CreateAt: time.Now(),
		Active:   true,
	}

	return &payment
}

func (p *Payment) validade() error {
	if p.Name == "" {
		return errors.New("O Nome o brigatório")
	}

	return nil

}

func (payment *Payment) formatSpace() {
	payment.Name = strings.TrimSpace(payment.Name)

}
