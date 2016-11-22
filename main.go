package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"flag"
	"log"
)

//Response structure
type Response struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Code:   200,
		Result: fmt.Sprintf("%s %s", "hello", r.FormValue("name")),
	}
	json, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func main() {
	port := flag.String("port", ":8000", "Port to run application")
	flag.Parse()
	http.HandleFunc("/", validateName(hello))
	http.ListenAndServe(":"+*port, nil)
}

func validateName(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if ok := r.URL.Query()["name"]; ok == nil {
			err := fmt.Errorf("Please provide a name")
			writeResponse(w, err, http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")
		if !isValidName(name) {
			err := fmt.Errorf("Name has %d character(s).Please provide more than 1 character.", len(name))
			writeResponse(w, err, http.StatusBadRequest)
			return
		}
		fn(w, r)
	}
}

func isValidName(name string) bool {
	if len(name) > 1 {
		return true
	}
	return false
}

func writeResponse(w http.ResponseWriter, err error, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	j, err := json.Marshal(
		Response{
			Code:   status,
			Result: err.Error(),
		})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(j)
}
