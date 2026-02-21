package geo_test

import (
	"demo/weather/geo"
	"errors"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	// Arrange - подготовка, expected результат
	city := "London"
	expected := geo.GeoData{City: "London"}

	// Act - выполняемая функция
	got, err := geo.GetMyLocation(city)
	if err != nil {
		t.Error("Ошибка получения города")
	}
	// Assert - проверка результата c expected
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получено %v", expected, got)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "NoCity"
	_, err := geo.GetMyLocation(city)
	if !errors.Is(err, geo.ErrNoCity) {
		t.Errorf("Ожидалось %v, получено %v", geo.ErrNoCity, err)
	}
}
