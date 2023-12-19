package application

import (
	"converter/application"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	from := "RUB"
	to := "USD"
	value := 1.234567890
	URL := fmt.Sprintf("https://v6.exchangerate-api.com/v6/8dcd9d44d24821df9839a14e/pair/%v/%v/%v", from, to, value)
	resp, err := http.Get(URL)
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
	expected := document.(map[string]interface{})["conversion_result"].(float64)

	actual, err := application.Convert(from, to, value)
	if err != nil {
		panic(err)
	}

	result := math.Abs(expected - actual)
	assert.GreaterOrEqual(t, 0.01, result)
}
