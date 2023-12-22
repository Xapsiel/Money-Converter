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
	//возвращаемые от апи данные
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
		//проверка на дату и существование данной валюты
		if elem.Year == year && elem.Month == month && elem.Day == day && elem.Conversion_rates[from_currency] != nil && elem.Conversion_rates[to_currency] != nil { //если дата совпадает с актуальной датой,то возвращаем курс относительно доллара
			return elem.Conversion_rates[to_currency].(float64) / elem.Conversion_rates[from_currency].(float64), nil
		}
	}
	return -1.0, fmt.Errorf("Курс (%v-%v) за эту дату не был зафиксирован сервером", from_currency, to_currency)
}
func (db *DataBase) GetCoefByDate(from_currency, to_currency string, day, month, year int) (float64, error) {
	data, err := db.Read() //считываем json файл
	if err != nil {
		return -1.0, err
	}
	var document []jsonFormat
	err = json.Unmarshal(data, &document) //превращаем json в формат []jsonFormat
	if err != nil {
		return -1.0, err
	}
	for _, elem := range document {
		//проверка на дату и существование данной валюты
		if elem.Year == year && elem.Month == month && elem.Day == day && elem.Conversion_rates[from_currency] != nil && elem.Conversion_rates[to_currency] != nil { //если дата совпадает с актуальной датой,то возвращаем курс относительно доллара
			return elem.Conversion_rates[to_currency].(float64) / elem.Conversion_rates[from_currency].(float64), nil
		}
	}
	return -1.0, fmt.Errorf("Курс (%v-%v) за эту дату не был зафиксирован сервером", from_currency, to_currency)

}
func (db *DataBase) Read() ([]byte, error) {
	file, err := os.OpenFile("repo/data.json", os.O_RDWR, 0666) //открыаем бд
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(file) //считываем что внутри
	if err != nil {
		return nil, err
	}

	return data, nil //возвращаем в формате массива байтов
}
func (db *DataBase) Write(code string) error {
	file, err := db.Read() //получаем массив данных
	if err != nil {
		return err
	}
	var document []jsonFormat
	err = json.Unmarshal(file, &document) //все что было в бд до этого момента - записываем в document

	if err != nil {
		return err
	}
	date := time.Now()
	new_document := jsonFormat{}
	new_document.Year = date.Year()
	new_document.Month = int(date.Month())
	new_document.Day = date.Day()
	if !db.check(document, new_document.Day, new_document.Month, new_document.Year) { //если курс за этот день существует, то...
		return fmt.Errorf("Курс за этот день уже обновлен")
	}

	jsonRes, err := db.makeRequest("USD") //получаем мапу курсов
	if err != nil {
		return err
	}
	new_document.Conversion_rates = jsonRes   //записываем курс
	document = append(document, new_document) //добавляем новый документ в старый
	result, err := json.Marshal(document)     //превращем в массив байтов
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("repo/data.json", result, 0666) //записываем в файл
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
	resp, err := http.Get("https://v6.exchangerate-api.com/v6/8dcd9d44d24821df9839a14e/latest/" + code) //отправляем запрос по апи
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) //считываем ответ
	if err != nil {
		return nil, err
	}
	var document interface{}
	err = json.Unmarshal(body, &document) //приводим массив байтов в адекватный вид
	if err != nil {
		return nil, err
	}
	f := document.(map[string]interface{})["conversion_rates"].(map[string]interface{}) //возвразаем курсы относительно code

	return f, nil
}
