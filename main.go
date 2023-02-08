package main

import (
	"log"
	"movies/movieApi"
	"net/http"
	"os"
)

const APIKEY string = "APIKEY"
const HOST string = "www.omdbapi.com"

type ServiceConfiguration struct {
	server movieApi.Server
}

func (m *ServiceConfiguration) search(w http.ResponseWriter, r *http.Request) {
	result, err := m.server.Search(r.URL)
	if err != nil {

	}
	w.Write([]byte(result))
}

func (m *ServiceConfiguration) detail(w http.ResponseWriter, r *http.Request) {
}

func main() {
	server := movieApi.NewServer(os.Getenv(APIKEY), HOST)
	serviceConfiguration := ServiceConfiguration{server: server}
	//http.HandleFunc("/detail", serviceConfiguration.detail)
	http.HandleFunc("/", serviceConfiguration.search)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
