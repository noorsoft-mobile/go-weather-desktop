package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const openWeatherMapApiKey = "979f40a4b2fa0937f07b8481430d3432"

type WeatherInfo struct {
	List []WeatherListItem `json:list`
}

type WeatherListItem struct {
	Dt      int           `json:dt`
	Main    WeatherMain   `json:main`
	Weather []WeatherType `json:weather`
}

type WeatherMain struct {
	Temp      float32 `json:temp`
	FeelsLike float32 `json:feels_like`
	Humidity  int     `json:humidity`
}

type WeatherType struct {
	Icon string `json:icon`
}

func getWeatherForecast(result interface{}) error {
	var url = fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?q=Voronezh&cnt=4&units=metric&appid=%s", openWeatherMapApiKey)
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
	}
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(result)
}
