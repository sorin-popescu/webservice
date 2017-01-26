package handlers

import (
	"bytes"
	"context"
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
	ctx := context.Background()
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		err := fmt.Sprintf("Invalid id: %v", ps.ByName("id"))
		log.Println(err)
		response.WriteResponse(w, err, http.StatusBadRequest)
		return
	}

	repository := developers.Developer{}
	dev := repository.GetByID(&id)

	back := make(chan string)

	go getSearch(ctx, dev, back)
	//
	//for i := 0; i < 2; i++ {
	//	select {
	//		case <-ctx.Done():
	//			fmt.Println(ctx.Err())
	//		case u := <-back:
	//			fmt.Println(u)
	//	}
	//}
	result := struct {
		Developer developers.Developer `json:"developer"`
		Link      string               `json:"link"`
	}{
		dev,
		<-back,
	}
	response.WriteResponse(w, result, http.StatusOK)
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

func getSearch(ctx context.Context, developer developers.Developer, back chan string) {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	resp, err := Crawl("https://www.google.co.uk/search?q=" + developer.Name)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	log.Println(buf.String())

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp.Status)
	resp.Body.Close()
	back <- resp.Status
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}
