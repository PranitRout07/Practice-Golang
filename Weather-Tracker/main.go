package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type WeatherAPI struct {
	API string `json:"OpenWeatherAPI"`
}

type GetWeatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func APIConfig(filename string) (WeatherAPI, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return WeatherAPI{}, os.ErrClosed
	}
	var x WeatherAPI
	err = json.Unmarshal(bytes, &x)
	if err != nil {
		return WeatherAPI{}, nil
	}
	fmt.Printf("u: %+v\n", x)
	return x, err

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
	defer resp.Body.Close()
	var d GetWeatherData
	json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		return GetWeatherData{}, err
	}
	return d, nil
}

func main() {
	// filename := ".apiConfig"
	// a, err := APIConfig(filename)
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(a.API)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/weather/",
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
	http.ListenAndServe(":8080",nil)
}
