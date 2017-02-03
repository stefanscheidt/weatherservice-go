package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const apiKey = "eb0f044bf44071e91b4232cbde8cd921"

type Report struct {
	Name    string `json:"name"`
	Weather []struct {
		Description string `json:"description"`
	}
	Main struct {
		Temperature float64 `json:"temp"`
	}
	Sys struct {
		Sunrise int `json:"sunrise"`
		Sunset  int `json:"sunset"`
	}
}

func getData(city string) ([]byte, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q={%s}&appid=%s", city, apiKey)
	res, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	if res.StatusCode != 200 {
		log.Print(string(data))
		return nil, fmt.Errorf("Bad response status %d", res.StatusCode)
	}

	return data, nil
}

func GetForecast(city string) (Report, error) {
	data, err := getData(city)
	if err != nil {
		log.Print(err)
		return Report{}, err
	}

	var report Report
	if err := json.Unmarshal(data, &report); err != nil {
		log.Print(err)
		return report, err
	}

	return report, nil
}
