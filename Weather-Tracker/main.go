package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type WeatherAPI struct {
	API string `json:"OpenWeatherAPI"`
}

func APIConfig(filename string) (WeatherAPI, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return WeatherAPI{}, os.ErrClosed
	}
	var x WeatherAPI
	err = json.Unmarshal(bytes, &x)
	if err != nil {
		return WeatherAPI{},nil
	}
	fmt.Printf("u: %+v\n", x)
	return x,err

}
func main() {
	filename := ".apiConfig"
	a, err := APIConfig(filename)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(a.API)

}
