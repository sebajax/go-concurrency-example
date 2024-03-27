package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sebajax/go-concurrency-example/pkg"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Define a route to fetch weather data for a given city
	app.Get("/weather/:city", func(c *fiber.Ctx) error {
		// Get the city parameter from the URL
		city := c.Params("city")

		weatherData, err := pkg.Example1(city)
		if err != nil {
			log.Println(err)
		}

		return c.JSON(weatherData)
	})

	// Define a route to fetch weather data for a given city
	app.Get("/weather", func(c *fiber.Ctx) error {
		weatherData, err := pkg.Example2()
		if err != nil {
			log.Println(err)
		}

		return c.JSON(weatherData)
	})

	// Start the HTTP server
	if err := app.Listen(":3000"); err != nil {
		fmt.Println("Error:", err)
	}
}
