package main

import (
	"fmt"
	"go-rest-api-course/cmd/internal/db"
)

func Run() error {
	fmt.Println("Starting up our app")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	fmt.Println("Successfully connected to DB")

	return nil
}

func main() {
	fmt.Println("Hello, World")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
