package domain

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Pagamento struct {
	Id       string    `json:"id"`
	IdParque string    `json:"id parque"`
	Nome     string    `json:"nome"`
	Ativo    bool      `json:"ativo"`
	DataAt   time.Time `json:"data cadastro"`
}

type PagamentoRepository interface {
	GetById(id string) (Pagamento, error)
	Create(Pagamento Pagamento) (Pagamento, error)
	Update(Pagamento Pagamento) (Pagamento, error)
	Delete(id string) (string, error)
	FindAll() ([]Pagamento, error)
}

func NewPagamento() *Pagamento {
	Pagamento := Pagamento{
		Id:     uuid.NewV4().String(),
		Ativo:  true,
		DataAt: time.Now(),
	}

	return &Pagamento
}

func (p *Pagamento) validade() error {
	if p.Nome == "" {
		return errors.New("O Nome o brigatório")
	}

	if p.IdParque == "" {
		return errors.New("O codigo do parque é obrigatório")
	}

	return nil

}
