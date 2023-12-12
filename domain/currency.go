package domain

import (
	"converter/repo"
)

type Currency struct { //структура валюты
	From_currency string
	To_currency   string
	Coefficient   float64
}

func (c *Currency) GetCoefficient(from_currency, to_currency string) float64 {
	a := repo.DataBase{Db: c} // обращаемся к Дб, можно ис
	return a.GetCoef(from_currency, to_currency)
}

func (c *Currency) GetAllCoefficient(from_currency, to_currency string) map[string]float64 {
	a := repo.DataBase{Db: c}
	return a.GetAllCoef(from_currency, to_currency)

}

func (c *Currency) ReadData() []byte {
	a := repo.DataBase{Db: c}
	return a.Read()

}
func (c *Currency) UpdateDB(code string) {
	a := repo.DataBase{Db: c}
	a.Write(c.From_currency)

}
