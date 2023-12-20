package test

import (
	domain "converter/domain"
	"converter/repo"
	"testing"
)

func TestRead(t *testing.T) {
	db := repo.DataBase{Db: &domain.Currency{}}
	_, err := db.Read()
	if err != nil {
		t.Errorf("Возникла следующая ошибка:\n%v", err.Error())
	}
}
