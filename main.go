package main

import (
	"log"
	"net/http"

	"github.com/DaveBlooman/go-app/config"
	"github.com/DaveBlooman/go-app/general"
	"github.com/DaveBlooman/go-app/persistence"
	"github.com/ant0ine/go-json-rest/rest"
)

var Version string

func main() {
	general := general.API{}

	dbConfig := config.DatabaseFromEnv()

	db := persistence.Persistence{}
	err := db.Init(&dbConfig)
	if err != nil {
		log.Fatal(err, "Failed to connect to database")
	}

	general.DB = &db

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/", general.GetArticles),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
