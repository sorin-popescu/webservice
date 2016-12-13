package handlers

import (
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
