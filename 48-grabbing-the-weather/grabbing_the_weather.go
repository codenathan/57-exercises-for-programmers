package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type WeatherData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

func fetchWeather(city string) (*WeatherData, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return nil, err
	}

	apiKey := os.Getenv("WEATHER_KEY")

	baseURL := "http://api.openweathermap.org/data/2.5/weather"
	query := url.Values{}
	query.Add("q", city)
	query.Add("appid", apiKey)
	query.Add("units", "imperial")

	resp, err := http.Get(baseURL + "?" + query.Encode())
	if err != nil {
		return nil, fmt.Errorf("error fetching weather data: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	var data WeatherData
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("error parsing weather data: %v", err)
	}

	return &data, nil
}

func displayWeather(data *WeatherData) {
	if data == nil {
		fmt.Println("No weather data available")
		return
	}
	fmt.Printf("%s weather:\n", data.Name)
	fmt.Printf("%.0f degrees Fahrenheit\n", data.Main.Temp)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Where are you? ")
	city, _ := reader.ReadString('\n')
	city = strings.TrimSpace(city)

	if city == "" {
		fmt.Println("You must enter a city")
		return
	}

	weather, err := fetchWeather(city)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	displayWeather(weather)

}
