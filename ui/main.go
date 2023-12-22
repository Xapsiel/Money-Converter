package main

import (
	"converter/application"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../www/"))))
	router.HandleFunc("/convert/{from}/{to}/{cash}", converterHandler)
	router.HandleFunc("/history_convert/{from}/{to}/{cash}/{day}/{month}/{year}", convertAnyDateHandler)

	router.HandleFunc("/update", updateHandler)
	http.ListenAndServe(":8080", router)
}

type converterFormat struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Value  float64 `json: "value"`
	Result float64 `json:"result"`
}

func converterHandler(rw http.ResponseWriter, r *http.Request) {
	value := mux.Vars(r)
	from := value["from"]
	to := value["to"]
	cash, _ := strconv.ParseFloat(value["cash"], 64)
	rw.Header().Set("Content-Type", "application/json")
	forRes, err := application.ConvertToday(from, to, cash)
	if err != nil {
		panic(err)
	}
	result := converterFormat{From: from, To: to, Value: cash, Result: forRes}
	res, _ := json.Marshal(result)
	rw.Write([]byte(res))
}

func convertAnyDateHandler(rw http.ResponseWriter, r *http.Request) {
	value := mux.Vars(r)
	from := value["from"]
	to := value["to"]
	cash, _ := strconv.ParseFloat(value["cash"], 64)
	day, _ := strconv.Atoi(value["day"])
	month, _ := strconv.Atoi(value["month"])
	year, _ := strconv.Atoi(value["year"])
	rw.Header().Set("Content-Type", "application/json")
	forRes, err := application.ConverAnyDay(from, to, day, month, year, cash)
	result := converterFormat{From: from, To: to, Value: cash, Result: forRes}

	if err != nil {
		if err.Error() == "Курс (USD-RUB) за эту дату не был зафиксирован сервером" {
			result.Result = 0
		} else {
			panic(err)
		}
	}
	res, _ := json.Marshal(result)
	rw.Write([]byte(res))

}
func updateHandler(rw http.ResponseWriter, r *http.Request) {
	err := application.Update()
	if err != nil {
		panic(err)
	}
}
