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
	router.HandleFunc("/convert/{from}/{to}/{cash}", converterHandler)
	router.HandleFunc("/makeGraph/{from}/{to}/{cash}", GraphHandler)
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
	forRes, err := application.Convert(from, to, cash)
	if err != nil {
		panic(err)
	}
	result := converterFormat{From: from, To: to, Value: cash, Result: forRes}
	res, _ := json.Marshal(result)
	rw.Write([]byte(res))
}

type GraphFormat struct {
	From   string             `json:"from"`
	To     string             `json:"to"`
	Value  float64            `json:"value"`
	Result map[string]float64 `json:"result"`
}

func GraphHandler(rw http.ResponseWriter, r *http.Request) {
	value := mux.Vars(r)
	from := value["from"]
	to := value["to"]
	cash, _ := strconv.ParseFloat(value["cash"], 64)
	forRes, err := application.MakeGraph(from, to, cash)
	if err != nil {
		panic(err)
	}
	result := GraphFormat{From: from, To: to, Value: cash, Result: forRes}
	res, _ := json.Marshal(result)
	rw.Write([]byte(res))
}

func updateHandler(rw http.ResponseWriter, r *http.Request) {
	err := application.Update()
	if err != nil {
		panic(err)
	}
}
