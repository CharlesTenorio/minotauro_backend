package repositorios

import (
	uuid "github.com/satori/go.uuid"
	d "minotauro/app/domain"
	"time"
)

type PagamentoRepository interface {
	GetById(id string) (d.Pagamento, error)
	Create(Pagamento d.Pagamento) (d.Pagamento, error)
	Update(Pagamento d.Pagamento) (d.Pagamento, error)
	Delete(id string) (string, error)
	FindAll() ([]d.Pagamento, error)
}

func NewPagamento() *d.Pagamento {
	Pagamento := d.Pagamento{
		IdPagamento: uuid.NewV4().String(),
		Ativo:       true,
		DataAt:      time.Now(),
	}

	return &Pagamento
}
