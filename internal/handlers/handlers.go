package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

const apikey = "034f2881ddd93b4e24a5ddbb78ca416b"
const LondonLat = "51.5073219"
const LondonLon = "-0.1276474"
const GetWeatherURL = "https://api.openweathermap.org/data/2.5/weather"

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

func MainHandler(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Cant read file", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?lat=%s&lon=%s&appid=%s", GetWeatherURL, LondonLat, LondonLon, apikey), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", "Bearer 034f2881ddd93b4e24a5ddbb78ca416b")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Статус:", resp.Status)
	w.Write(buf.Bytes())
}

//var answer []WeatherResponse
// if err = json.Unmarshal(buf.Bytes(), &answer); err != nil {
// 	http.Error(w, err.Error(), http.StatusBadRequest)
// 	return
// }
