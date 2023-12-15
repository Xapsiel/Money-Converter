package test

import (
	"converter/application"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConverterHandler(t *testing.T) {
	resp, err := http.Get("http://localhost/convert/RUB/USD/100")
	if err != nil {
		panic(err)
	}
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
	expected := document.(map[string]interface{})["result"].(float64)
	assert.Equal(t, expected, application.Convert("RUB", "USD", 100))
}
