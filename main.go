package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello world")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("PORT is not in the environment")
	}

	portString := os.Getenv("PORT")

	fmt.Println("PORT:", portString)
}
