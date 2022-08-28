package services

import (
	"fmt"
	r "minotauro/app/aplicacao/repositorios"
	"minotauro/app/domain"
	"time"
)

type CampeonatoService struct {
	RepoCamoponanto r.CampeonatoRepository
}

var validacao domain.MessagenErros

func NewCamponatoServece(repo *r.CamponatoRepoDb) *CampeonatoService {
	return &CampeonatoService{RepoCamoponanto: repo}
}

func (c *CampeonatoService) GetById(id string) (domain.Campeonato, error) {
	cp, err := c.RepoCamoponanto.GetById(id)
	if err != nil {
		return domain.Campeonato{}, err
	}
	return cp, nil
}

func (c *CampeonatoService) Create(nome string, data_ini, data_final time.Time, vpLugar,
	vsLugar, vtLugar, vqLugar float64, cartaz string) (domain.Campeonato, error) {
	camp := domain.NovoCampeonato()
	validacao.CampoObriatorio = nome
	validacao.DataInicial = data_ini
	validacao.DataFinal = data_final
	validacao.DataObriagoria = data_ini
	validacao.ChecarMenorValor()

	camp.Nome = nome
	camp.DataInicial = data_ini
	camp.DataFinal = data_final
	camp.ValPrimeiroLugar = vpLugar
	camp.ValSegundoLugar = vsLugar
	camp.ValTerceiroLugar = vtLugar
	camp.ValQuartoLugar = vqLugar
	camp.Cartaz = cartaz

	cp, err := c.RepoCamoponanto.Create(*camp)
	if err != nil {
		return domain.Campeonato{}, err
	}
	return cp, nil
}

func (c *CampeonatoService) Update(id, nome string, data_ini, data_final time.Time, vpLugar,
	vsLugar, vtLugar, vqLugar float64, cartaz string, ativo bool) (domain.Campeonato, error) {
	camp := domain.NovoCampeonato()
	camp.IdCamponato = id
	camp.Nome = nome
	camp.DataInicial = data_ini
	camp.DataFinal = data_final
	camp.ValPrimeiroLugar = vpLugar
	camp.ValSegundoLugar = vsLugar
	camp.ValTerceiroLugar = vtLugar
	camp.ValQuartoLugar = vqLugar
	camp.Ativo = ativo

	cp, err := c.RepoCamoponanto.Update(*camp)
	if err != nil {
		return domain.Campeonato{}, err
	}
	return cp, nil

}

func (c *CampeonatoService) Delete(id string) (string, error) {
	_, err := c.RepoCamoponanto.Delete(id)
	if err != nil {
		fmt.Print(err)
		return "DEU RIIM", err
	}

	return "Dado excluido com sucesso", nil
}

func (c *CampeonatoService) FindAll() ([]domain.Campeonato, error) {
	camp, err := c.RepoCamoponanto.FindAll()
	if err != nil {
		return []domain.Campeonato{}, err
	}
	return camp, nil
}
