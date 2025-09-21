package server

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonitaga/weather_service/internal/dto"
	"github.com/tonitaga/weather_service/internal/uri"
)

const (
	GeoLocationBaseUrl = "http://api.openweathermap.org/geo/1.0/direct"
	WeatherBaseUrl     = "https://api.openweathermap.org/data/2.5/weather"
)

var uriBuilder = uri.NewBuilder()

func (s *server) getWeahterHandler(c *gin.Context) {
	var requestBody dto.WeatherRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	var uriBuilder = uri.NewBuilder()

	uriBuilder.BaseUrl(GeoLocationBaseUrl)
	uriBuilder.Param("q", requestBody.City, requestBody.CountryCode, requestBody.State)
	uriBuilder.Param("limit", 1)
	uriBuilder.Param("appid", s.config.ApiKey)

	geolocationsResponse, err := http.Get(uriBuilder.Build())
	if err != nil {
		c.String(http.StatusInternalServerError, "On make request for get geolocation. Cause: %v", err)
		return
	}

	defer geolocationsResponse.Body.Close()

	var geolocationsData []dto.GeolocationBody
	if err := json.NewDecoder(geolocationsResponse.Body).Decode(&geolocationsData); err != nil {
		c.String(http.StatusInternalServerError, "On deserialize of geolocation response. Cause: %v", err)
		return
	}

	geolocation := geolocationsData[0]

	uriBuilder.Reset()

	uriBuilder.BaseUrl(WeatherBaseUrl)
	uriBuilder.Param("lon", geolocation.Longitude)
	uriBuilder.Param("lat", geolocation.Latitude)
	uriBuilder.Param("appid", s.config.ApiKey)

	weatherResponse, err := http.Get(uriBuilder.Build())
	if err != nil {
		c.String(http.StatusInternalServerError, "On make request for get weahter. Cause: %v", err)
		return
	}

	defer weatherResponse.Body.Close()

	var weatherData dto.WeatherResponseBody
	if err := json.NewDecoder(weatherResponse.Body).Decode(&weatherData); err != nil {
		c.String(http.StatusInternalServerError, "On deserialize of weahter response. Cause: %v", err)
		return
	}

	c.JSON(http.StatusOK, weatherData)
}
