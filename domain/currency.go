package domain

import (
	"converter/repo"
)

type Currency struct { //структура валюты
	From_currency string
	To_currency   string
	Coefficient   float64
}

func (c *Currency) ReadData() ([]byte, error) {
	a := repo.DataBase{Db: c}
	return a.Read()

}
func (c *Currency) UpdateDB(code string) error {
	a := repo.DataBase{Db: c}
	return a.Write(c.From_currency)

}
func (c *Currency) GetCoefficient(from_currency, to_currency string, day, month, year int) (float64, error) {
	a := repo.DataBase{Db: c}
	return a.GetCoef(from_currency, to_currency, day, month, year)

}
