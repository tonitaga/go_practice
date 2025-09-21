package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tonitaga/weather_service/internal/config"
	"github.com/tonitaga/weather_service/internal/dto"
	"github.com/tonitaga/weather_service/internal/server"
	"github.com/tonitaga/weather_service/internal/uri"
)

type GeoData struct {
	Name       string            `json:"name"`
	LocalNames map[string]string `json:"local_names,omitempty"`
	Latitude   float64           `json:"lat"`
	Longitude  float64           `json:"lon"`
	Country    string            `json:"country"`
	State      string            `json:"state,omitempty"`
}

// https://openweathermap.org/api/geocoding-api

func main() {
	server := server.NewServer(config.NewConfig())
	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}

	config := config.NewConfig()

	uriBuilder := uri.NewBuilder()
	uriBuilder.BaseUrl("http://api.openweathermap.org/geo/1.0/direct")
	uriBuilder.Param("q", "Kazan", "RU", "")
	uriBuilder.Param("limit", 1)
	uriBuilder.Param("appid", config.ApiKey)

	fmt.Println(uriBuilder.Build())
	response, err := http.Get(uriBuilder.Build())
	if err != nil {
		log.Fatalln(err)
	}

	var geolocationData []dto.GeolocationData
	if err := json.NewDecoder(response.Body).Decode(&geolocationData); err != nil {
		log.Fatalln(err)
	}

	geolocation := geolocationData[0]

	fmt.Println("Latitude:", geolocation.Latitude)
	fmt.Println("Longitude:", geolocation.Longitude)
	fmt.Println("Country:", geolocation.Country)
	fmt.Println("Name:", geolocation.Name)

	uriBuilder.Reset()

	uriBuilder.BaseUrl("https://api.openweathermap.org/data/2.5/weather")
	uriBuilder.Param("lat", geolocation.Latitude)
	uriBuilder.Param("lon", geolocation.Longitude)
	uriBuilder.Param("appid", config.ApiKey)

	response, err = http.Get(uriBuilder.Build())
	if err != nil {
		log.Fatalln(err)
	}

	var weatherData dto.WeatherResponse
	if err := json.NewDecoder(response.Body).Decode(&weatherData); err != nil {
		log.Fatalln(err)
	}

	const AbsoluteZero = 273.15

	fmt.Println("Temperature:", weatherData.Main.Temp-AbsoluteZero)
	fmt.Println("Feels like:", weatherData.Main.FeelsLike-AbsoluteZero)
}
