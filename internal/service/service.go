package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const APIkey = "034f2881ddd93b4e24a5ddbb78ca416b"
const GetWeatherURL = "https://api.openweathermap.org/data/2.5/weather"
const GetCountryURL = "http://api.openweathermap.org/geo/1.0/direct"

// http://api.openweathermap.org/geo/1.0/direct?q=London&limit=5&appid={API key}
type Geocoding struct {
	Name       string            `json:"name"`
	LocalNames map[string]string `json:"local_names,omitempty"` // если поле отсутствует, не сериализуется
	Lat        float64           `json:"lat"`
	Lon        float64           `json:"lon"`
	Country    string            `json:"country"`
	State      string            `json:"state,omitempty"`
}
type WeatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Rain struct {
		OneH float64 `json:"1h"`
	} `json:"rain"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int64 `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func Detector(s string) string {
	return "clean"
}

func GeocodingCountryName(country string) ([]Geocoding, error) {
	var geo []Geocoding
	var buf bytes.Buffer
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?q=%s&limit=1&appid=%s", GetCountryURL, country, APIkey), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", APIkey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(buf.Bytes(), &geo); err != nil {
		return nil, err
	}
	if len(geo) == 0 {
		return nil, errors.New("couldnt find city")
	}
	return geo, nil
}

func WeatherInfo(geo []Geocoding) (WeatherResponse, error) {
	var weather WeatherResponse
	var buf bytes.Buffer
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?lat=%f&lon=%f&appid=%s", GetWeatherURL, geo[0].Lat, geo[0].Lon, APIkey), nil)
	if err != nil {
		return weather, err
	}
	req.Header.Set("Authorization", APIkey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return weather, err
	}
	defer resp.Body.Close()
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return weather, err
	}
	if err = json.Unmarshal(buf.Bytes(), &weather); err != nil {
		return weather, err
	}
	return weather, nil
}
