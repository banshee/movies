package main

import (
	"log"
	"movies/movieApi"
	"net/http"
	"os"
	"strings"
)

const APIKEY string = "APIKEY"
const HOST string = "www.omdbapi.com"

type ServiceConfiguration struct {
	server movieApi.Server
}

func (m *ServiceConfiguration) search(w http.ResponseWriter, r *http.Request) {
	result, err := m.server.Search(r.URL)
	if err == nil {
		w.Write([]byte(result))
	}
}

func (m *ServiceConfiguration) detail(w http.ResponseWriter, r *http.Request) {
	result, err := m.server.Detail(r.URL)
	if err == nil {
		w.Write([]byte(result))
	}
}

func (m *ServiceConfiguration) handler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, "/detail") {
		m.detail(w, r)
	} else {
		m.search(w, r)
	}
}

func main() {
	server := movieApi.NewServer(os.Getenv(APIKEY), HOST)
	serviceConfiguration := ServiceConfiguration{server: server}
	http.HandleFunc("/", serviceConfiguration.handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
