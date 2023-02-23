package main

import (
	"context"
	"fmt"
	"sdivyansh59/GO-REST-API-V2/internal/comment"
	"sdivyansh59/GO-REST-API-V2/internal/db"
)

// Run - is going to be responsible for the
// instantiation and startup of our go application
func Run() error {
	fmt.Println("Starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)

	cmtService.PostComment(
		context.Background(),
		comment.Comment{
			ID: "605338af-3783-4e80-bfa1-22c23804ae48",
			Slug: "manual-test",
			Author: "Elliot",
			Body: "Hello World",
		},
	)


	fmt.Println(cmtService.GetComment(
		context.Background(),
		"605338af-3783-4e80-bfa1-22c23804ae48",
	))

	return nil
}



func main() {
	fmt.Println("Go REST Api")
	if err := Run(); err != nil {
		fmt.Println(err)
	}

}