package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// WeatherInfo struct
type WeatherInfo struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

func temperatureFor(location string, temperature chan float64) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=0b4ee5e61edff05f27f95951b9211da5", location)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching data for %s: %s", location, err)
		close(temperature)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading data for %s: %s", location, err)
		close(temperature)
	}

	var data WeatherInfo
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Error decoding data for %s: %s", location, err)
		close(temperature)
	}

	temperature <- (data.Main.Temp - 273.15)
	fmt.Printf("Temperature at %s is %.2f °C\n", location, (data.Main.Temp - 273.15))
}

func main() {
	cities := [...]string{"Mumbai", "London", "Amsterdam", "Singapore", "Paris"}
	temperature := make(chan float64)

	for _, city := range cities {
		go temperatureFor(city, temperature)
	}

	for range cities {
		<-temperature
	}
}
