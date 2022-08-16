package domain

import (
	"errors"
	"time"
)

type Parque struct {
	IdParque          string    `json:"id_parque"`
	Nome              string    `json:"nome"`
	IdCamponato       string    `json:"id_campeonato"`
	Endereco          string    `json:"endereço"`
	Bairro            string    `json:"Bairro"`
	Cidade            string    `json:"cidade"`
	Uf                string    `json:"uf"`
	Ativado           bool      `json:"activate"`
	Img               string    `json:"img"`
	PorntoDeRefrencia string    `json:"reference point"`
	DataAt            time.Time `json:"data"`
}

func (p *Parque) validade() error {
	if p.Nome == "" {
		return errors.New("O Nome o brigatório")
	}

	if p.Endereco == "" {
		return errors.New("Endereço e obrigatório")
	}

	if p.Bairro == "" {
		return errors.New("Bairro obriatório")
	}

	if p.Cidade == "" {
		return errors.New("Cidade é obriatório")
	}

	if p.Uf == "" {
		return errors.New("UF obriatório")
	}

	if p.PorntoDeRefrencia == "" {
		return errors.New("Ponto de referencia e obritatório")
	}
	return nil

}
