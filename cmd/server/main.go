package main

import (
	"context"
	"fmt"
	"go-rest-api-course/cmd/internal/comment"
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

	// fmt.Println("Successfully connected to DB")

	cmtService := comment.NewService(db)

	cmtService.PostComment(
		context.Background(),
		comment.Comment{
			ID:     "05744018-7bf5-4d80-a681-3c29328a9ed9",
			Slug:   "testing1232",
			Author: "Elliot",
			Body:   "Hello. Is there anybody",
		},
	)

	fmt.Println(cmtService.GetComment(
		context.Background(),
		"05744018-7bf5-4d80-a681-3c29328a9ed9",
	))

	return nil
}

func main() {
	fmt.Println("Hello, World")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
