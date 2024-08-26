package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mayckol/stress-test/handler"
	"github.com/mayckol/stress-test/http_client"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using default S.O. environment variables.")
	}

	viaCepClient := http_client.NewViaCepClient(os.Getenv("VIA_CEP_API_URL"), true)

	weatherClient := http_client.NewWeatherClientClient(os.Getenv("WEATHER_API_URL"), os.Getenv("WEATHER_API_KEY"), true)
	weatherHandler := handler.NewWeatherHandler(viaCepClient, weatherClient)

	http.HandleFunc("/weather/{zipCode}", weatherHandler.Weather)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("server running on port %s\n", port)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server")
	}
}
