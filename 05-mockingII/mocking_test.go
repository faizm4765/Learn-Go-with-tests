package main

import (
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// func TestUserNotFound(t *testing.T) {
// 	// baseUrl := "https://jsonplaceholder.typicode.com/"
// 	mocksServer := httptest.NewServer(
// 		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			http.Error(w, "Not Found", http.StatusNotFound)
// 		}),
// 	)

// 	client := NewApiClient(mocksServer.URL)
// 	statusCode, _ := client.FetchUser(1)
// 	expectedStatusCode := 404
// 	if statusCode != expectedStatusCode {
// 		t.Errorf("Got %d Expected %d", statusCode, expectedStatusCode)
// 	}
// }

// func TestUsersListResponse(t *testing.T) {
// 	// baseUrl := "https://jsonplaceholder.typicode.com/"
// 	mocksServer := httptest.NewServer(
// 		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			time.Sleep(10 * time.Second)
// 			w.WriteHeader(http.StatusOK)
// 			jsonOutput := `
// 				[
// 					{"id": 1, "name": "John Doe"},
// 					{"id": 2, "name": "Alex Doe"},
// 					{"id": 3, "name": "Harry Doe"}
// 				]`

// 			w.Write([]byte(jsonOutput))
// 		}),
// 	)

// 	fmt.Println("Mock server URL:", mocksServer.URL)
// 	client := NewApiClient(mocksServer.URL)
// 	users, _ := client.FetchAllUsers()
// 	expectedUsersCount := 3
// 	if len(users) != expectedUsersCount {
// 		t.Errorf("Got %d Expected %d", len(users), expectedUsersCount)
// 	}
// }

func TestFetchUsersWithTimeout(t *testing.T) {
	mocksServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(200 * time.Second) // Simulate a delay
		}),
	)

	client := NewApiClient(mocksServer.URL, 100*time.Millisecond)
	_, err := client.FetchAllUsers()
	if !isTimeOut(err) {
		t.Errorf("Expected timeout error, but got: %v", err)
	}
}

func isTimeOut(err error) bool {
	// write logic to check if the error is a timeout error
	if err == nil {
		return false
	}

	netErr, ok := err.(net.Error)
	if ok && netErr.Timeout() {
		return true
	}

	return false
}
