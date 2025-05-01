package main

import (
	"testing"
)

func TestUserNotFound(t *testing.T) {
	baseUrl := "https://jsonplaceholder.typicode.com/"
	client := NewApiClient(baseUrl)
	statusCode, _ := client.FetchUser(1)
	expectedStatusCode := 404
	if statusCode != expectedStatusCode {
		t.Errorf("Got %d Expected %d", statusCode, expectedStatusCode)
	}
}
