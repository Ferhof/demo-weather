package weather_test

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"errors"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	geoData := geo.GeoData{City: "London"}
	format := 3

	weatherData, _ := weather.GetWeather(geoData, format)
	if !strings.Contains(weatherData, "London") {
		t.Errorf("Ожидали %v, получили %v", "London", weatherData)
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{"Big format", 147},
	{"Zero format", 0},
	{"Minus format", -1},
}

func TestGetWeatherWrongFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			geoData := geo.GeoData{City: "London"}
			_, err := weather.GetWeather(geoData, tc.format)
			if !errors.Is(err, weather.ErrWrongFormat) {
				t.Errorf("Ожидали %v, получили %v", weather.ErrWrongFormat, err)
			}
		})
	}

}
