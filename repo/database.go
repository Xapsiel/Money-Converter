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

func (db *DataBase) GetCoef(from_currency, to_currency string) (float64, error) {
	data, err := db.Read() //считываем json файл
	if err != nil {
		return -1.0, err
	}
	var document []jsonFormat
	err = json.Unmarshal(data, &document) //превращаем json в формат []jsonFormat
	if err != nil {
		return -1.0, err
	}
	date := time.Now()
	year := date.Year()
	month := int(date.Month())
	day := date.Day()
	for _, elem := range document {
		if elem.Year == year && elem.Month == month && elem.Day == day { //если дата совпадает с актуальной датой,то возвращаем курс относительно доллара
			return elem.Conversion_rates[to_currency].(float64) / elem.Conversion_rates[from_currency].(float64), nil
		}
	}
	return -1.0, err
}
func (db *DataBase) GetAllCoef(from_currency, to_currency string) (map[string]float64, error) {
	result := make(map[string]float64)
	data, err := db.Read() //считываем json файл
	if err != nil {
		return nil, err
	}
	var document []jsonFormat
	err = json.Unmarshal(data, &document) //превращаем json в формат []jsonFormat
	if err != nil {
		return nil, err
	}
	for _, elem := range document {
		year := elem.Year
		month := elem.Month
		day := elem.Day
		key := fmt.Sprintf("%v-%v-%v", day, month, year)
		result[key] = elem.Conversion_rates[to_currency].(float64) / elem.Conversion_rates[from_currency].(float64)
	}
	return result, nil

}
func (db *DataBase) Read() ([]byte, error) {
	file, err := os.OpenFile("repo/data.json", os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}
func (db *DataBase) Write(code string) error {
	file, err := db.Read()
	if err != nil {
		return err
	}
	var document []jsonFormat
	err = json.Unmarshal(file, &document)

	if err != nil {
		return err
	}
	date := time.Now()
	new_document := jsonFormat{}
	new_document.Year = date.Year()
	new_document.Month = int(date.Month())
	new_document.Day = date.Day()
	if !db.check(document, new_document.Day, new_document.Month, new_document.Year) {
		return fmt.Errorf("Курс за этот день уже обновлен")
	}

	jsonRes, err := db.makeRequest("USD")
	if err != nil {
		return err
	}
	new_document.Conversion_rates = jsonRes
	document = append(document, new_document)
	result, err := json.Marshal(document)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("data.json", result, 0666)
	if err != nil {
		return err
	}
	return nil
}
func (db *DataBase) check(array []jsonFormat, day, month, year int) bool {
	for _, elem := range array {
		if elem.Day == day && elem.Month == month && elem.Year == year {
			return false
		}
	}
	return true
}
func (db *DataBase) makeRequest(code string) (map[string]interface{}, error) {
	resp, err := http.Get("https://v6.exchangerate-api.com/v6/8dcd9d44d24821df9839a14e/latest/" + code)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var document interface{}
	err = json.Unmarshal(body, &document)
	if err != nil {
		return nil, err
	}
	f := document.(map[string]interface{})["conversion_rates"].(map[string]interface{})

	return f, nil
}
