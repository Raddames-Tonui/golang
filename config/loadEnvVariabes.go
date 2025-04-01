package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
  // Initialize any necessary configurations or settings here
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")

	fmt.Println("Loaded PORT", os.Getenv("PORT"))
  }
}
