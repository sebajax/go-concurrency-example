package pkg

// WeatherData represents weather information for a city
type WeatherData struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	Condition   string  `json:"condition"`
}

// FetchWeather fetches weather data for a given city from an external API (simulated here)
func FetchWeather(city string) *WeatherData {
	// Simulate fetching weather data from an external API
	// In a real-world scenario, this function would make an HTTP request to a weather API
	// and parse the response to extract relevant weather information.
	// For simplicity, we'll just simulate some data here.
	temperature := 20.0 + (float64(len(city)) / 10.0) // Simulated temperature
	condition := "Sunny"                              // Simulated weather condition
	return &WeatherData{City: city, Temperature: temperature, Condition: condition}
}
