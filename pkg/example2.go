package pkg

import (
	"sync"
)

// Example2 will fetch weather data for multiple cities concurrently
func Example2() ([]*WeatherData, error) {
	// Split the cities string into individual city names
	cityList := []string{"New York", "London", "Paris"} // Example cities, you can change this list

	// Fetch weather data for each city concurrently
	weatherChan := make(chan *WeatherData, len(cityList))
	var wg sync.WaitGroup
	for _, city := range cityList {
		wg.Add(1)
		go func(city string) {
			defer wg.Done()
			weather := FetchWeather(city)
			weatherChan <- weather
		}(city)
	}

	// Wait for all weather data to be fetched
	go func() {
		wg.Wait()
		close(weatherChan)
	}()

	// Collect weather data for all cities from the channel
	var weatherData []*WeatherData
	for weather := range weatherChan {
		weatherData = append(weatherData, weather)
	}

	return weatherData, nil
}
