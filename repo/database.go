package repo

import (
	domain "converter/domain/interface"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type DataBase struct {
	Db domain.Convertion
}

type jsonFormat struct {
	Year             int                    `json: "year"`
	Month            int                    `json: "month"`
	Day              int                    `json: "month"`
	Conversion_rates map[string]interface{} `json:"conversion_rates"`
}
type JsonResult struct {
	result, documentation, terms_of_use, time_last_update_utc, time_next_update_utc, base_code string
	time_last_update_unix, time_next_update_unix                                               int
	conversion_rates                                                                           map[string]float64
}

func (db *DataBase) GetCoef(from_currency, to_currency string) float64 {
	data := db.Read() //считываем json файл
	var document []jsonFormat
	err := json.Unmarshal(data, &document) //превращаем json в формат []jsonFormat
	if err != nil {
		panic(err)
	}
	date := time.Now()
	year := date.Year()
	month := int(date.Month())
	day := date.Day()
	for _, elem := range document {
		if elem.Year == year && elem.Month == month && elem.Day == day { //если дата совпадает с актуальной датой,то возвращаем курс относительно доллара
			return elem.Conversion_rates[to_currency].(float64) / elem.Conversion_rates[from_currency].(float64)
		}
	}
	return -1
}
func (db *DataBase) GetAllCoef(from_currency, to_currency string) map[string]float64 {
	result := make(map[string]float64)
	data := db.Read() //считываем json файл
	var document []jsonFormat
	err := json.Unmarshal(data, &document) //превращаем json в формат []jsonFormat
	if err != nil {
		panic(err)
	}
	for _, elem := range document {
		year := elem.Year
		month := elem.Month
		day := elem.Day
		key := fmt.Sprintf("%v-%v-%v", day, month, year)
		result[key] = elem.Conversion_rates[to_currency].(float64) / elem.Conversion_rates[from_currency].(float64)
	}
	return result

}
func (db *DataBase) Read() []byte {
	file, err := os.OpenFile("data.json", os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return data
}
func (db *DataBase) Write(code string) {
	file := db.Read()
	var document []jsonFormat
	err := json.Unmarshal(file, &document)

	if err != nil {

		panic(err)
	}
	date := time.Now()
	new_document := jsonFormat{}
	new_document.Year = date.Year()
	new_document.Month = int(date.Month())
	new_document.Day = date.Day()
	if !db.check(document, new_document.Day, new_document.Month, new_document.Year) {
		return
	}

	jsonRes := db.makeRequest("USD")
	new_document.Conversion_rates = jsonRes
	document = append(document, new_document)
	result, err := json.Marshal(document)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("data.json", result, 0666)
	if err != nil {
		panic(err)
	}

}
func (db *DataBase) check(array []jsonFormat, day, month, year int) bool {
	for _, elem := range array {
		if elem.Day == day && elem.Month == month && elem.Year == year {
			return false
		}
	}
	return true
}
func (db *DataBase) makeRequest(code string) map[string]interface{} {
	resp, err := http.Get("https://v6.exchangerate-api.com/v6/8dcd9d44d24821df9839a14e/latest/" + code)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var document interface{}
	err = json.Unmarshal(body, &document)
	if err != nil {
		panic(err)
	}
	f := document.(map[string]interface{})["conversion_rates"].(map[string]interface{})

	return f
}
