package domain

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Pagamento struct {
	gorm.Model
	IdPagamento string    `json:"id_formapg" gorm:"type:uuid;primary_key"`
	IdParque    string    `json:"id parque"`
	Nome        string    `json:"nome"`
	Ativo       bool      `json:"ativo"`
	DataAt      time.Time `json:"data cadastro"`
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
		IdPagamento: uuid.NewV4().String(),
		Ativo:       true,
		DataAt:      time.Now(),
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
