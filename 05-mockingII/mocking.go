package main

import (
	"fmt"
	"net/http"
)

type ApiClient struct {
	baseUrl string
}

func NewApiClient(baseUrl string) ApiClient {
	return ApiClient{baseUrl: baseUrl}
}

func (client ApiClient) FetchUser(id int) (int, error) {
	url := fmt.Sprintf("%s/users/%d", client.baseUrl, id)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching user:", err)
		return 0, err
	}

	defer resp.Body.Close()
	return resp.StatusCode, err
}

func main() {
	fmt.Println("Hello, World!")
	client := NewApiClient("https://jsonplaceholder.typicode.com/")
	statusCode, err := client.FetchUser(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Status Code:", statusCode)
		fmt.Println("User fetched successfully!!")
	}
}
