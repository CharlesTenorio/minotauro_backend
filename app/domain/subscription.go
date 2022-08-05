package domain

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Subscription struct {
	Id              string    `json:"id"`
	IdCowboy        string    `json:"id"`
	IdPark          string    `json:"id_park"`
	IdCategory      string    `json:"id_categoria"`
	IdPayment       string    `json:"id_payment"`
	IdEvent         string    `json:"id_payment"`
	Esteira         string    `json:"esteria"`
	Ticket          []string  `json:"senhas"`
	InscriptionDate time.Time `json:"data da inscrição"`
	Horse           string    `json:"cavalo"`
	Active          bool      `json:"activate"`
	img             string    `json:"img"`
	CreateAt        time.Time `json:"data"`
}

type SubscriptionRepository interface {
	GetById(id string) (Subscription, error)
	Create(Subscription Subscription) (Subscription, error)
	Update(Subscription Subscription) (Subscription, error)
	Delete(id string) (string, error)
	FindAll() ([]Subscription, error)
}

func NewSubscription() *Subscription {
	subscription := Subscription{
		Id:       uuid.NewV4().String(),
		CreateAt: time.Now(),
		Active:   true,
	}

	return &subscription
}

func (s *Subscription) validade() error {
	if s.Horse == "" {
		return errors.New("O Nome o brigatório")
	}

	if s.Esteira == "" {
		return errors.New("Esteira e o brigatório")
	}

	if s.IdCategory == "" {
		return errors.New("Codigo da catergoria é obrigatório")
	}

	if s.IdCowboy == "" {
		return errors.New("Codigo do vaqueiro é obrigatório")
	}

	if s.IdPark == "" {
		return errors.New("Codigo do Parque é obrigatório")
	}

	if s.IdPayment == "" {
		return errors.New("Codigo do Pagamento é obrigatório")
	}

	if s.IdEvent == "" {
		return errors.New("Codigo do Evento é obrigatório")
	}

	return nil

}
