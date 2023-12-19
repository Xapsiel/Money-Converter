package test

import (
	"converter/application"
	"testing"
)

func TestWrite(t *testing.T) {

	err := application.Update()

	if err != nil && err.Error() != "Курс за этот день уже обновлен" {
		t.Errorf("Возникла следующая ошибка:\n%v", err.Error())
	}

}
