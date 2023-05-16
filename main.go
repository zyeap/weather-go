package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`

	Current struct {
		TempF     float64 `json:"temp_f"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
}

func main() {
	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=" + os.Getenv("WEATHER_API_TOKEN") + "&q=chicago")
	if err != nil {
		panic(err)
	}

	if res.StatusCode != 200 {
		fmt.Printf("Received status code: %d\n", res.StatusCode)
		panic("Weather API not available")
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	// fmt.Println(string(body))

	var weather Weather
	err = json.Unmarshal(body, &weather)

	if err != nil {
		panic(err)
	}

	location, current := weather.Location, weather.Current

	fmt.Printf("%s, %s: %.0fF\n", location.Name, location.Country, current.TempF)

}
