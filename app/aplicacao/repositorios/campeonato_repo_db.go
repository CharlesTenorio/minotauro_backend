package repositorios

import (
	"database/sql"
	"fmt"
	d "minotauro/app/domain"
)

type CamponatoRepoDb struct {
	db *sql.DB
}

var validation d.Campeonato

func NewCampeonatoDb(db *sql.DB) *CamponatoRepoDb {
	return &CamponatoRepoDb{db: db}
}

func (c *CamponatoRepoDb) FindAll() ([]d.Campeonato, error) {
	var campeonato d.Campeonato
	var campeonatos []d.Campeonato
	sqlSelect := "SELECT id_campeonato, nome, data_inicial, data_fimal, valor_primeiro_lugar, valor_segundo_lugar, valor_terceiro_lugar, valor_quarto_lugar, cartaz, ativo, data_at, soft_delete"
	stmt, err := c.db.Query(sqlSelect + " FROM campeonato")
	if err != nil {
		fmt.Print(err)
		return []d.Campeonato{}, err
	}
	for stmt.Next() {
		err = stmt.Scan(&campeonato.IdCamponato, &campeonato.Nome, &campeonato.DataInicial, &campeonato.DataFinal,
			&campeonato.ValPrimeiroLugar, &campeonato.ValSegundoLugar, &campeonato.ValTerceiroLugar, &campeonato.ValQuartoLugar, &campeonato.Cartaz,
			&campeonato.Ativo, &campeonato.DataAt, &campeonato.SoftDelete)
		campeonatos = append(campeonatos, campeonato)
	}

	return campeonatos, nil

}

func (c *CamponatoRepoDb) GetById(id string) (d.Campeonato, error) {

	var campeonato d.Campeonato
	sqlSelect := "SELECT id_campeonato, nome, data_inicial, data_fimal, valor_primeiro_lugar, valor_segundo_lugar, valor_terceiro_lugar, valor_quarto_lugar, cartaz, ativo, data_at, soft_delete"
	stmt, err := c.db.Prepare(sqlSelect + " FROM campeonato where id_campeonato=$1")

	if err != nil {
		fmt.Print(err)
		return d.Campeonato{}, err
	}
	err = stmt.QueryRow(id).Scan(&campeonato.IdCamponato, &campeonato.Nome, &campeonato.DataInicial, &campeonato.DataFinal,
		&campeonato.ValPrimeiroLugar, &campeonato.ValSegundoLugar, &campeonato.ValTerceiroLugar, &campeonato.ValQuartoLugar,
		&campeonato.Cartaz, &campeonato.Ativo, &campeonato.DataAt, &campeonato.SoftDelete)
	if err != nil {

		return d.Campeonato{}, err
	}

	return campeonato, nil

}

func (c *CamponatoRepoDb) Create(campeonato d.Campeonato) (d.Campeonato, error) {

	stmt, err := c.db.Prepare("INSERT INTO campeonato (id_campeonato, nome, data_inicial, data_fimal," +
		" valor_primeiro_lugar, valor_segundo_lugar, valor_terceiro_lugar, valor_quarto_lugar, " +
		"cartaz, ativo, data_at, soft_delete)" +
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);")
	if err != nil {
		return d.Campeonato{}, err
	}
	_, err = stmt.Exec(campeonato.IdCamponato, campeonato.Nome, campeonato.DataInicial, campeonato.DataFinal,
		campeonato.ValPrimeiroLugar, campeonato.ValSegundoLugar,
		campeonato.ValTerceiroLugar, campeonato.ValQuartoLugar,
		campeonato.Cartaz, campeonato.Ativo, campeonato.DataAt, campeonato.SoftDelete)
	if err != nil {
		return d.Campeonato{}, err
	}
	return campeonato, nil
}

func (c *CamponatoRepoDb) Update(campeonato d.Campeonato) (d.Campeonato, error) {
	/*err := validation.Prepere()
	if err != nil {
		return d.Campeonato{}, err
	}*/
	stmt, err := c.db.Prepare("UPDATE campeonato " +
		"SET nome=$2, data_inicial=$3, data_fimal=$4, valor_primeiro_lugar=$5," +
		"valor_segundo_lugar=$6, valor_terceiro_lugar=$7, valor_quarto_lugar=$8," +
		"cartaz=$9, ativo=$10, data_at=$11" +
		"WHERE id_campeonato=$1;")
	if err != nil {
		return d.Campeonato{}, err
	}
	_, err = stmt.Exec(campeonato.IdCamponato, campeonato.Nome, campeonato.DataInicial, campeonato.DataFinal,
		campeonato.ValPrimeiroLugar, campeonato.ValSegundoLugar,
		campeonato.ValTerceiroLugar, campeonato.ValQuartoLugar)
	if err != nil {
		return d.Campeonato{}, err
	}
	return campeonato, nil
}

// soft delete nao execlui do banco apenas marca como excluido
func (c *CamponatoRepoDb) Delete(id string) (string, error) {

	stmt, err := c.db.Prepare("UPDATE campeonato SET  soft_delete=true WHERE id_campeonato=$1;")
	if err != nil {
		fmt.Print(err)

		return "deu erro", err

	}
	_, err = stmt.Exec(id)
	if err != nil {
		return "deu ruim", err
	}
	err = stmt.Close()
	if err != nil {
		return "deu ruim", err
	}
	return "Dado excluido com sucesso", nil

}
