package domain

import (
	"time"
)

type Categoria struct {
	IdCategoria string    `json:"id_categoria"`
	IdParque    string    `json:"id_parque"`
	Nome        string    `json:"nome"`
	Valor       float64   `json:"valor"`
	Ativado     bool      `json:"activate"`
	DataAt      time.Time `json:"data"`
}
