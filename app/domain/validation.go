package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func FormatSpace(campo string) string {
	campo = strings.TrimSpace(campo)
	return campo
}

func VDataObrigatoria(dateV time.Time) error {
	if dateV.IsZero() {
		return errors.New("Data e obrigatória")
	}
	return nil
}

func ValidarCampoObrigatorio(campo string) error {
	if campo == "" {
		return errors.New("e obriagatório o" + campo)
	}
	return nil
}
func VmenorValo(val1, val2 float64, campo1, campo2 string) error {
	if val1 <= val2 {
		sval1 := fmt.Sprintf("%f", val1)
		sval2 := fmt.Sprintf("%f", val2)
		return errors.New("o " + campo1 + sval1 + "não pode ser menor ou igual a" + sval2 + " " + campo2)
	}
	return nil
}

func ChecarDatas(dataIni, dataFim time.Time) error {
	if dataIni.Equal(dataFim) {
		return errors.New("as datas não pode ser iguais")
	}
	if dataFim.After(dataIni) {
		return errors.New("a data final não pode ser menor q data inicial")
	}
	return nil
}
