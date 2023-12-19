package repo

import (
	domain "converter/domain/interface"
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	resp, err := http.Get("https://v6.exchangerate-api.com/v6/8dcd9d44d24821df9839a14e/pair/" + from_currency + "/" + to_currency) //отправляем запрос по апи
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) //считываем ответ
	if err != nil {
		return -1, err
	}
	var document interface{}
	err = json.Unmarshal(body, &document) //приводим массив байтов в адекватный вид
	if err != nil {
		return -1, err
	}
	f := document.(map[string]interface{})["conversion_rate"] //возвразаем курсы относительно code
	return f.(float64), nil
}
