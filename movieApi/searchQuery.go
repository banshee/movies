package movieApi

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const apiKeyParameter string = "apikey"
const idParameter string = "i"

// Server contains the data needed to talk to the backend server
type Server struct {
	apikey string
	host   string
}

// NewServer creates a new server configuration
func NewServer(apikey, host string) Server {
	return Server{
		apikey: apikey,
		host:   host,
	}
}

// RefactorSearch takes a URL and refactors it to pass through a search (not a detail)
// query to the backend
func (m *Server) RefactorSearch(q *url.URL) url.URL {
	return refactorSearchUrl(q, m.host, m.apikey)
}

// RefactorDetail takes a URL and refactors it to pass through a detail query
// to the backend
func (m *Server) RefactorDetail(q *url.URL) url.URL {
	return refactorDetailUrl(q, m.host, m.apikey)
}

// refactorSearchUrl adds the api key parameter, keeping the other search parameters
func refactorSearchUrl(incomingQuery *url.URL, destinationHost, apikey string) url.URL {
	result := *incomingQuery
	result.Host = destinationHost
	result.Scheme = "http"
	query := incomingQuery.Query()
	query.Add(apiKeyParameter, apikey)
	result.RawQuery = query.Encode()
	return result
}

// refactorSearchUrl adds the api key parameter and extracts the detail id from the path
func refactorDetailUrl(incomingQuery *url.URL, destinationHost, apikey string) url.URL {
	result := *incomingQuery
	result.Host = destinationHost
	result.Scheme = "http"
	query := url.Values{}
	pathElements := strings.Split(incomingQuery.Path, "/")
	query.Add(idParameter, pathElements[len(pathElements)-1])
	query.Add(apiKeyParameter, apikey)
	result.RawQuery = query.Encode()
	result.Path = ""
	return result
}

func (m *Server) Search(query *url.URL) (string, error) {
	url := m.RefactorSearch(query)
	return m.retrieveResults(url)
}

func (m *Server) Detail(query *url.URL) (string, error) {
	url := m.RefactorDetail(query)
	return m.retrieveResults(url)
}

// retrieveResults actually retrieves the results from the remote server
func (m *Server) retrieveResults(url url.URL) (string, error) {
	s := url.String()
	res, err := http.Get(s)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return "", errors.New(fmt.Sprintf("Status code %d", res.StatusCode))
	}
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return string(body), nil
}
