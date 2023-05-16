package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

// Weather struct maps fields to raw json response from GET api.weatherapi.com/v1/forecast.json
type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`

	Current struct {
		TempC     float64 `json:"temp_c"`
		TempF     float64 `json:"temp_f"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`

	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				TempF     float64 `json:"temp_f"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	weather_api_key, ok := os.LookupEnv("WEATHER_API_TOKEN")
	if !ok || weather_api_key == "" {
		panic("Weather API key not set")
	}

	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=" + weather_api_key + "&q=singapore")
	if err != nil {
		panic(err)
	}

	if res.StatusCode != 200 {
		fmt.Printf("Received status code: %d\n", res.StatusCode)
		panic("Weather API not available")
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	var weather Weather
	// parse json data into memory location of weather
	err = json.Unmarshal(body, &weather)

	if err != nil {
		panic(err)
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	fmt.Printf("%s\n", strings.Repeat("─", 75))
	fmt.Printf("%s, %s: %.0fF, %0.fC\n", location.Name, location.Country, current.TempF, current.TempC)
	fmt.Printf("Today's condition: %s\n", current.Condition.Text)
	fmt.Printf("%s\n", strings.Repeat("─", 75))

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(time.Now()) {
			continue
		}

		if hour.ChanceOfRain >= 50 {
			color.Cyan(
				"[%s - %0.fF, %0.fC] - [%0.f%% rain] - [%s]\n",
				date.Format("15:04"),
				hour.TempF,
				hour.TempC,
				hour.ChanceOfRain,
				hour.Condition.Text,
			)
		} else {
			fmt.Printf(
				"[%s - %0.fF, %0.fC] - [%0.f%% rain] - [%s]\n",
				date.Format("15:04"), // https://stackoverflow.com/questions/28087471/what-is-the-significance-of-gos-time-formatlayout-string-reference-time
				hour.TempF,
				hour.TempC,
				hour.ChanceOfRain,
				hour.Condition.Text,
			)
		}
	}

}
