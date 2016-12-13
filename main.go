package main

import (
	"net/http"

	"flag"

	"github.com/julienschmidt/httprouter"
	"github.com/sorin-popescu/webservice/handlers"
)

// func hello(w http.ResponseWriter, r *http.Request) {
// 	result := fmt.Sprintf("%s %s", "hello", r.FormValue("name"))
// 	writeResponse(w, result, http.StatusOK)
// }

func main() {
	port := flag.String("port", "8000", "Port to run application")
	flag.Parse()
	r := httprouter.New()
	r.GET("/developers", handlers.GetDevelopers)
	r.GET("/developers/:id", handlers.GetDeveloperByID)
	http.ListenAndServe(":"+*port, r)
}
