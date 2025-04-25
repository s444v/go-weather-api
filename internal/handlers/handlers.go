package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/s444v/go-weather-api/internal/service"
)

const WarmWeather = "https://i.pinimg.com/736x/bc/5d/82/bc5d8255aeb265bbacdeace5c9200d40.jpg"
const ColdWeather = "https://i.pinimg.com/736x/32/4f/98/324f98074dd22d307d6c8189da6e971e.jpg"

type PageData struct {
	Picture     string
	Country     string
	Temperature string
	Humidity    int
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Cant read file", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	country := r.FormValue("fieldCountry")
	geo, err := service.GeocodingCountryName(country)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	weatherInfo, err := service.WeatherInfo(geo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	weatherImgUrl := WarmWeather
	temperature := weatherInfo.Main.Temp - 273.15
	if temperature < 15 {
		weatherImgUrl = ColdWeather
	}
	data := PageData{
		Picture:     weatherImgUrl,
		Country:     geo[0].LocalNames["ru"],
		Temperature: fmt.Sprintf("%.1f", temperature),
		Humidity:    weatherInfo.Main.Humidity,
	}
	tmp, err := template.ParseFiles("weatherReport.html")
	if err != nil {
		http.Error(w, "Cant read file", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmp.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
