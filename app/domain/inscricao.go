package domain

import (
	"errors"
	"strings"
	"time"
)

type Inscricao struct {
	IdInscricao   string    `json:"id_inscricao"`
	IdVaqueiro    string    `json:"id_vaqueiro"`
	IdParque      string    `json:"id_park"`
	IdCategoria   string    `json:"id_categoria"`
	IdPagamento   string    `json:"id_payment"`
	IdEvento      string    `json:"id_evento"`
	Esteira       string    `json:"esteria"`
	Senhas        []string  `json:"senhas"`
	DataInscricao time.Time `json:"data da inscrição"`
	Cavalo        string    `json:"cavalo"`
	Representacao string    `json:"Representação"`
	MaisSenha     bool      `json:"Mais de uma senha"`
	Ativa         bool      `json:"activate"`
	Img           string    `json:"img"`
	DataAt        time.Time `json:"data"`
}

func (inscricao *Inscricao) Prepere() error {
	if erro := inscricao.validade(); erro != nil {
		return erro
	}
	inscricao.formatSpace()
	return nil
}

func (s *Inscricao) validade() error {
	if s.Cavalo == "" {
		return errors.New("O Nome o brigatório")
	}

	if s.Esteira == "" {
		return errors.New("Esteira e o brigatório")
	}

	if s.IdCategoria == "" {
		return errors.New("Codigo da catergoria é obrigatório")
	}

	if s.IdVaqueiro == "" {
		return errors.New("Codigo do vaqueiro é obrigatório")
	}

	if s.IdParque == "" {
		return errors.New("Codigo do Parque é obrigatório")
	}

	if s.IdPagamento == "" {
		return errors.New("Codigo do Pagamento é obrigatório")
	}

	if s.IdEvento == "" {
		return errors.New("Codigo do Evento é obrigatório")
	}

	return nil

}

func (inscricao *Inscricao) formatSpace() {
	inscricao.Esteira = strings.TrimSpace(inscricao.Esteira)
}
