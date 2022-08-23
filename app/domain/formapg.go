package domain

import (
	"errors"
	"strings"
	"time"
)

type Pagamento struct {
	IdPagamento string    `json:"id_formapg"`
	IdParque    string    `json:"id parque"`
	Nome        string    `json:"nome"`
	Ativo       bool      `json:"ativo"`
	DataAt      time.Time `json:"data cadastro"`
}

func (pagamento *Pagamento) Prepere() error {
	if erro := pagamento.validade(); erro != nil {
		return erro
	}
	pagamento.formatSpace()
	return nil
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

func (pagamento *Pagamento) formatSpace() {
	pagamento.Nome = strings.TrimSpace(pagamento.Nome)
}
