package application

import (
	"converter/domain"
	"fmt"
	"strconv"
)

func Convert(from_currency string, to_currency string, day int, month int, year int, value float64) (float64, error) {
	conv := domain.Currency{
		From_currency: from_currency,
		To_currency:   to_currency,
	}
	coefficient, err := conv.GetCoefficient(from_currency, to_currency, day, month, year) //получение актуального курса
	if err != nil {
		return 0, err
	}
	if value < 0 {
		value = 0
	}

	result := value * coefficient
	result, err = strconv.ParseFloat(fmt.Sprintf("%.2f", result), 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func Update(code string) error {
	var conv domain.Currency
	err := conv.UpdateDB(code)
	return err
}
