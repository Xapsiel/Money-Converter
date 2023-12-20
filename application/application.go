package application

import (
	"converter/domain"
	"fmt"
	"strconv"
)

func Convert(from_currency string, to_currency string, value float64) (float64, error) {
	conv := domain.Currency{
		From_currency: from_currency,
		To_currency:   to_currency,
	}
	coefficient, err := conv.GetCoefficient(from_currency, to_currency) //получение актуального курса

	if err != nil {
		return 0, err
	}

	conv.Coefficient = coefficient
	result := value * conv.Coefficient
	result, err = strconv.ParseFloat(fmt.Sprintf("%.2f", result), 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func MakeGraph(from_currency string, to_currency string, value float64) (map[string]float64, error) {
	conv := domain.Currency{
		From_currency: from_currency,
		To_currency:   to_currency,
	}
	result := make(map[string]float64)
	coefficient, err := conv.GetAllCoefficient(from_currency, to_currency)
	if err != nil {
		return nil, err
	}
	for key, elem := range coefficient {
		result[key] = elem * value
	}
	return result, nil

}

func Update() error {
	var conv domain.Currency
	err := conv.UpdateDB("USD")
	return err
}
