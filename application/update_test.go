package application

import "testing"

func TestUpdate(t *testing.T) {
	tests := "USD"

	err := Update(tests)
	if err != nil {
		if err.Error() != "Курс за этот день уже обновлен" {
			t.Fail()
		}
	}
}
