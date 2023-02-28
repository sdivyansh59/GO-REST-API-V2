package main

import (
	"fmt"
	"sdivyansh59/GO-REST-API-V2/internal/comment"
	"sdivyansh59/GO-REST-API-V2/internal/db"
	transportHttp "sdivyansh59/GO-REST-API-V2/internal/transport/http"
	//log "github.com/sirupsen/logrus"
)


// Run - is going to be responsible for the
// instantiation and startup of our go application
func Run() error {
	fmt.Println("Starting up our application")

	store, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := store.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	cmtService := comment.NewService(store)

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}



func main() {
	fmt.Println("Go REST Api")
	if err := Run(); err != nil {
		fmt.Println(err)
	}

}