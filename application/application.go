package application

import (
	"converter/domain"
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
	return result, nil
}
