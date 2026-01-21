package main

import (
	"fmt"
	"go-rest-api-course/cmd/internal/comment"
	"go-rest-api-course/cmd/internal/db"
	transportHttp "go-rest-api-course/cmd/internal/transport/http"
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

	// fmt.Println("Successfully connected to DB")

	cmtService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Hello, World")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
