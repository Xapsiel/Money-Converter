package domain

import (
	"testing"
	"time"
)

func TestReadingError(t *testing.T) {
	c := Currency{From_currency: "USD", To_currency: "RUB"}
	_, err := c.ReadData()
	if err != nil {
		t.Fail()
	}
}
func TestUpdateError(t *testing.T) {
	c := Currency{From_currency: "USD", To_currency: "RUB"}
	err := c.UpdateDB("USD")
	if err != nil {
		if err.Error() != "Курс за этот день уже обновлен" {
			t.Fail()
		}
	}
}

func TestGetCoefficientError(t *testing.T) {
	c := Currency{From_currency: "USD", To_currency: "RUB"}
	now := time.Now()
	_, err := c.GetCoefficient("USD", "RUB", now.Day(), int(now.Month()), now.Year())
	if err != nil {
		t.Fail()
	}
}
