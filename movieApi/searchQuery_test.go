package movieApi

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"os"
	"strings"
	"testing"
)

// Test to see that a search returns a Batman movie
func TestSearch(t *testing.T) {
	server := NewServer(os.Getenv("APIKEY"), "www.omdbapi.com")
	url, err := url.Parse("http://localhost:8080/?s=Batman&page=2")
	assert.Nil(t, err)
	result, err := server.Search(url)
	assert.Nil(t, err)
	// Could use a more specific test here
	assert.True(t, strings.Contains(result, "Batman: The Killing Joke"))
}

// Test that a detail request returns details for the right movie
func TestDetail(t *testing.T) {
	server := NewServer(os.Getenv("APIKEY"), "www.omdbapi.com")
	url, err := url.Parse("http://localhost:8080/detail/tt1569923")
	assert.Nil(t, err)
	result, err := server.Detail(url)
	assert.Nil(t, err)
	// Could use a more specific test here
	assert.True(t, strings.Contains(result, "Brandon Vietti"))
}
