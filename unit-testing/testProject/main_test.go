package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	total := Add(1, 3)
	assert.NotNil(t, total, "The total should not be nil")
	assert.Equal(t, 4, total, "Expecting 4")
}

func TestSubstract(t *testing.T) {
	total := Substract(1, 3)
	assert.NotNil(t, total, "The total should not be nil")
	assert.Equal(t, -2, total, "Expecting -2")
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	return router
}

func TestRootEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "OK Response is expected")
	assert.Equal(t, "Hello worl", response.Body.String(), "Incorrect body found")
}
