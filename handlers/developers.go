package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/sorin-popescu/webservice/developers"
	"github.com/sorin-popescu/webservice/response"
)

//GetDevelopers will return all developers
func GetDevelopers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	repository := developers.Developer{}
	devs := repository.FindAll()
	res := make(map[string]developers.Developer)
	for index, value := range devs {
		res[fmt.Sprint(index)] = value
	}

	response.WriteResponse(w, res, http.StatusOK)
}

//GetDeveloperByID will return one developer by id
func GetDeveloperByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		err := fmt.Sprintf("Invalid id: %v", ps.ByName("id"))
		log.Println(err)
		response.WriteResponse(w, err, http.StatusBadRequest)
		return
	}

	repository := developers.Developer{}
	dev := repository.GetByID(&id)

	response.WriteResponse(w, dev, http.StatusOK)
}

//AddDeveloper will add a new developer to the map
func AddDeveloper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	developer := developers.Developer{}
	err := decodeResponse(r, &developer)
	if err != nil {
		err = fmt.Errorf("Error code: %d", 400)
	}

	repository := developers.Developer{}
	developers := repository.AddOne(&developer)

	response.WriteResponse(w, developers, http.StatusCreated)
}

func decodeResponse(r *http.Request, object interface{}) error {
	err := json.NewDecoder(r.Body).Decode(object)
	if err != nil {
		err = fmt.Errorf("Error code: %d", 400)
	}
	return err
}
