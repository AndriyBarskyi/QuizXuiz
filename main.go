package main

import (
	"QuizXuiz/config"
	"QuizXuiz/routes"
	"fmt"
	"log"
)

func main() {
	// Load environment variables
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Connect to the database
	config.ConnectDB(cfg)

	// Set up routes
	r := routes.SetupRouter()

	// Run the server
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
