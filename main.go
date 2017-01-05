package main

import (
	"net/http"

	"flag"

	"github.com/julienschmidt/httprouter"
	"github.com/sorin-popescu/webservice/handlers"
)

func main() {
	port := flag.String("port", "8000", "Port to run application")
	flag.Parse()
	r := httprouter.New()
	r.GET("/developers", handlers.GetDevelopers)
	r.POST("/developers", handlers.AddDeveloper)
	r.GET("/developers/:id", handlers.GetDeveloperByID)
	http.ListenAndServe(":"+*port, r)
}
