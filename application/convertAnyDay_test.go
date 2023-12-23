package application

import (
	"fmt"
	"testing"
	"time"
)

func TestConvert(t *testing.T) {
	type testStruct struct {
		from  string
		to    string
		value float64
		day   int
		month int
		year  int
	}
	now := time.Now()
	yesterday := now.Add(time.Hour * (-24))
	tommorow := now.Add(time.Hour * 24)
	testsList := []testStruct{
		testStruct{from: "USD", to: "RUB", value: 100, day: now.Day(), month: int(now.Month()), year: now.Year()},
		testStruct{from: "USD", to: "RUB", value: 100, day: yesterday.Day(), month: int(yesterday.Month()), year: yesterday.Year()},
		testStruct{from: "USD", to: "RUB", value: 100, day: tommorow.Day(), month: int(tommorow.Month()), year: tommorow.Year()},

		testStruct{from: "USD", to: "RUB", value: -100, day: now.Day(), month: int(now.Month()), year: now.Year()},
		testStruct{from: "USD", to: "RUB", value: -100, day: yesterday.Day(), month: int(yesterday.Month()), year: yesterday.Year()},
		testStruct{from: "USD", to: "RUB", value: -100, day: tommorow.Day(), month: int(tommorow.Month()), year: tommorow.Year()},

		testStruct{from: "USD", to: "RUB", value: 0, day: now.Day(), month: int(now.Month()), year: now.Year()},
		testStruct{from: "USD", to: "RUB", value: 0, day: yesterday.Day(), month: int(yesterday.Month()), year: yesterday.Year()},
		testStruct{from: "USD", to: "RUB", value: 0, day: tommorow.Day(), month: int(tommorow.Month()), year: tommorow.Year()},
	}
	for _, elem := range testsList {
		expected, err := Convert(elem.from, elem.to, elem.day, elem.month, elem.year, elem.value)
		if err != nil {
			if err.Error() == "Курс за эту дату не был зафиксирован сервером" && elem.day != tommorow.Day() {
				t.Fail()
			}
		}
		fmt.Println(expected)
	}

}
