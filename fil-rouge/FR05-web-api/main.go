package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Meteo struct {
	Latitude     float64      `json:"latitude"`
	Longitude    float64      `json:"longitude"`
	Timezone     string       `json:"timezone"`
	Current      Current      `json:"current"`
	CurrentUnits CurrentUnits `json:"current_units"`
}

type Current struct {
	Temperature2m float64 `json:"temperature_2m"`
	Time          string  `json:"time"`
	WindSpeed10m  float64 `json:"wind_speed_10m"`
}

type CurrentUnits struct {
	Temperature2m string `json:"temperature_2m"`
	Time          string `json:"time"`
	WindSpeed10m  string `json:"wind_speed_10m"`
}

func recupererMeteo(baseURL string, lat, lon float64) (Meteo, error) {
	var meteo Meteo

	url := fmt.Sprintf(
		"%s/v1/forecast?latitude=%f&longitude=%f&current=temperature_2m,wind_speed_10m",
		baseURL, lat, lon,
	)

	resp, err := http.Get(url)
	if err != nil {
		return meteo, fmt.Errorf("error sending request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return meteo, fmt.Errorf("the API returned %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&meteo)
	if err != nil {
		return meteo, fmt.Errorf("error decoding JSON: %w", err)
	}

	return meteo, nil
}

func main() {
	baseURL := "http://localhost:8081"
	latitude := 48.89
	longitude := 2.20

	meteo, err := recupererMeteo(baseURL, latitude, longitude)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Weather information:")
	fmt.Println("Latitude:", meteo.Latitude)
	fmt.Println("Longitude:", meteo.Longitude)
	fmt.Println("Timezone:", meteo.Timezone)

	fmt.Println("Temperature:", meteo.Current.Temperature2m, meteo.CurrentUnits.Temperature2m)
	fmt.Println("Time:", meteo.Current.Time)
	fmt.Println("Wind speed:", meteo.Current.WindSpeed10m, meteo.CurrentUnits.WindSpeed10m)
}
