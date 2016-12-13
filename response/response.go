package response

import (
	"encoding/json"
	"log"
	"net/http"
)

//Response structure
type Response struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
}

func WriteResponse(w http.ResponseWriter, body interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	j, err := json.Marshal(
		Response{
			Code:   status,
			Result: body,
		})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(j)
}
