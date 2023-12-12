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
	conv.Coefficient = conv.GetCoefficient(from_currency, to_currency)
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
	coefficient := conv.GetAllCoefficient(from_currency, to_currency)
	for key, elem := range coefficient {
		result[key] = elem * value
	}
	return result

}

func Update() {
	var conv domain.Currency
	conv.UpdateDB("USD")
}
