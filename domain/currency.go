package domain

import (
	"converter/repo"
)

type Currency struct { //структура валюты
	From_currency string
	To_currency   string
	Coefficient   float64
}

func (c *Currency) GetCoefficient(from_currency, to_currency string) (float64, error) {
	a := repo.DataBase{Db: c} // обращаемся к Дб
	return a.GetCoef(from_currency, to_currency)
}
