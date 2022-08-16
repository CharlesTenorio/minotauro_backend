package domain

import (
	"errors"
	"time"
)

type Pagamento struct {
	IdPagamento string    `json:"id_formapg"`
	IdParque    string    `json:"id parque"`
	Nome        string    `json:"nome"`
	Ativo       bool      `json:"ativo"`
	DataAt      time.Time `json:"data cadastro"`
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
