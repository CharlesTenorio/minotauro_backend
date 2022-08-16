package repositorios

import (
	uuid "github.com/satori/go.uuid"
	d "minotauro/app/domain"
	"time"
)

type InscricaoRepository interface {
	GetById(id string) (d.Inscricao, error)
	Create(Inscricao d.Inscricao) (d.Inscricao, error)
	Update(Inscricao d.Inscricao) (d.Inscricao, error)
	Delete(id string) (string, error)
	FindAll() ([]d.Inscricao, error)
}

func NewInscricao() *d.Inscricao {
	inscricao := d.Inscricao{
		IdInscricao: uuid.NewV4().String(),
		DataAt:      time.Now(),
		Ativa:       true,
	}

	return &inscricao
}
