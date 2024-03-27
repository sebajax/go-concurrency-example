package pkg

import (
	"sync"
)

// Example1 will fetch weather data for a city concurrently
func Example1(city string) (*WeatherData, error) {
	// Fetch weather data for the city concurrently
	weatherChan := make(chan *WeatherData)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		weather := FetchWeather(city)
		weatherChan <- weather
	}()

	// Wait for the weather data to be fetched in another go rutine
	go func() {
		wg.Wait()
		close(weatherChan)
	}()

	// Retrieve weather data from the channel
	weather := <-weatherChan
	return weather, nil
}
