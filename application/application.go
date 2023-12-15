package application

import (
	"converter/domain"
	"fmt"
)

func Convert(from_currency string, to_currency string, value float64) float64 {
	conv := domain.Currency{
		From_currency: from_currency,
		To_currency:   to_currency,
	}
	coefficient, err := conv.GetCoefficient(from_currency, to_currency)
	if err != nil {
		panic(err)
	}
	conv.Coefficient = coefficient
	if conv.Coefficient < 0 {
		panic(fmt.Errorf("Нет действующего курса"))
	}
	result := value * conv.Coefficient
	return result
}

func MakeGraph(from_currency string, to_currency string, value float64) map[string]float64 {
	conv := domain.Currency{
		From_currency: from_currency,
		To_currency:   to_currency,
	}
	result := make(map[string]float64)
	coefficient, err := conv.GetAllCoefficient(from_currency, to_currency)
	if err != nil {
		panic(err)
	}
	for key, elem := range coefficient {
		result[key] = elem * value
	}
	return result

}

func Update() error {
	var conv domain.Currency
	err := conv.UpdateDB("USD")
	return err
}
