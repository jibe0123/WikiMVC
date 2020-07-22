package main

import (
	"github.com/jibe0123/WikiMVC/database"
	"github.com/jibe0123/WikiMVC/router"
	"log"
	"net/http"
)

func main() {
	port := "8080"

	newRouter := router.NewRouter()
	database.Connect()
	log.Print("\nConnected to database")
	log.Print("\n🚀 Server started on port :  " + port)

	errServ := http.ListenAndServe(":"+port, newRouter)
	if errServ != nil {
		log.Fatalf("could not serve on port %s", port)
	}
}
