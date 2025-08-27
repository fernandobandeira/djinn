package main

import (
	"fmt"
	"os"
	
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if it exists (ignore error if it doesn't)
	_ = godotenv.Load()
	
	app, cleanup, err := InitializeApp()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize application: %v\n", err)
		os.Exit(1)
	}
	if cleanup != nil {
		defer cleanup()
	}
	
	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Application error: %v\n", err)
		os.Exit(1)
	}
}