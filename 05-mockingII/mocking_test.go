package main

import (
	"fmt"
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

func TestUsersListResponse(t *testing.T) {
	// baseUrl := "https://jsonplaceholder.typicode.com/"
	mocksServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			jsonOutput := `
				[
					{"id": 1, "name": "John Doe"},
					{"id": 2, "name": "Alex Doe"},
					{"id": 3, "name": "Harry Doe"}
				]`

			w.Write([]byte(jsonOutput))
		}),
	)

	fmt.Println("Mock server URL:", mocksServer.URL)
	client := NewApiClient(mocksServer.URL)
	users, _ := client.FetchAllUsers()
	expectedUsersCount := 3
	if len(users) != expectedUsersCount {
		t.Errorf("Got %d Expected %d", len(users), expectedUsersCount)
	}
}
