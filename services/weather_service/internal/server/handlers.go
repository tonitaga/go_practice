package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/tonitaga/weather_service/internal/dto"
	"github.com/tonitaga/weather_service/internal/uri"
)

const (
	GeoLocationBaseUrl = "http://api.openweathermap.org/geo/1.0/direct"
	WeatherBaseUrl     = "https://api.openweathermap.org/data/2.5/weather"
)

func (s *server) executeTemplateIndex(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := template.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *server) handleWeatherRequest(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Invalid content type. Expected application/json", http.StatusBadRequest)
		return
	}

	var requestBody dto.WeatherRequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	builder := uri.NewBuilder()

	builder.BaseUrl(GeoLocationBaseUrl)
	builder.Param("q", requestBody.City, requestBody.CountryCode, requestBody.State)
	builder.Param("limit", 1)
	builder.Param("appid", s.config.ApiKey)

	geolocationResponse, err := http.Get(builder.Build())
	if err != nil {
		http.Error(w, fmt.Sprintf("On make request for get geolocation. Cause: %v", err), http.StatusInternalServerError)
		return
	}

	defer geolocationResponse.Body.Close()

	var geolocationData []dto.GeolocationBody
	if err := json.NewDecoder(geolocationResponse.Body).Decode(&geolocationData); err != nil {
		http.Error(w, fmt.Sprintf("On deserialize of geolocation response. Cause: %v", err), http.StatusInternalServerError)
		return
	}

	if len(geolocationData) == 0 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	geolocation := geolocationData[0]

	builder.Reset()

	builder.BaseUrl(WeatherBaseUrl)
	builder.Param("lon", geolocation.Longitude)
	builder.Param("lat", geolocation.Latitude)
	builder.Param("appid", s.config.ApiKey)

	weatherResponse, err := http.Get(builder.Build())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer weatherResponse.Body.Close()

	var weatherData dto.WeatherResponseBody
	if err := json.NewDecoder(weatherResponse.Body).Decode(&weatherData); err != nil {
		http.Error(w, fmt.Sprintf("On deserialize of weahter response. Cause: %v", err), http.StatusInternalServerError)
		return
	}

	log.Println(weatherData)

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(&weatherData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
