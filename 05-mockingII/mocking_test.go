package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserNotFound(t *testing.T) {
	// baseUrl := "https://jsonplaceholder.typicode.com/"
	mocksServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Not Found", http.StatusNotFound)
		}),
	)

	client := NewApiClient(mocksServer.URL)
	statusCode, _ := client.FetchUser(1)
	expectedStatusCode := 404
	if statusCode != expectedStatusCode {
		t.Errorf("Got %d Expected %d", statusCode, expectedStatusCode)
	}
}
