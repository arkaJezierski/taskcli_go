package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/ajezierski/taskcli/cmd"
	"github.com/ajezierski/taskcli/internal/db"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}

	db.InitDB()

	cmd.Execute()
}
