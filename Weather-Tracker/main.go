package main

import (

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"github.com/rs/cors"
)

type WeatherAPI struct {
	API string `json:"OpenWeatherAPI"`
}

type WeatherInfo struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type GetWeatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`

	Weather []WeatherInfo `json:"weather"`
}


func APIConfig(filename string) (WeatherAPI, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return WeatherAPI{}, err
	}
	var x WeatherAPI
	err = json.Unmarshal(bytes, &x)
	if err != nil {
		return WeatherAPI{}, err
	}
	fmt.Printf("u: %+v\n", x)
	return x, nil

}
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from go!\n"))
}
func querry(city string) (GetWeatherData, error) {
	apiConfig, err := APIConfig(".apiConfig")
	if err != nil {
		return GetWeatherData{}, err
	}
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q="+city+"&appid="+apiConfig.API)
	if err != nil {
		return GetWeatherData{}, err
	}
	fmt.Println("response",resp.Body)
	defer resp.Body.Close()
	var d GetWeatherData
	json.NewDecoder(resp.Body).Decode(&d)
	d.Main.Kelvin = d.Main.Kelvin -273
	d.Main.Kelvin = float64(int(d.Main.Kelvin*100)) / 100
	if err != nil {
		return GetWeatherData{}, err
	}
	return d, nil
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/weather/",
		func(w http.ResponseWriter, r *http.Request) {
			city := strings.SplitN(r.URL.Path, "/", 3)[2]
			data, err := querry(city)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(data)
		})

	// Enable CORS for all routes using cors.Default().Handler
	handler := cors.Default().Handler(mux)

	log.Println("Listening....")
	http.ListenAndServe(":8080", handler)
}