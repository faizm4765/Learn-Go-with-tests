package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ApiClient struct {
	baseUrl       string
	timeoutPeriod time.Duration
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewApiClient(baseUrl string, timeoutPeriod time.Duration) ApiClient {
	return ApiClient{
		baseUrl:       baseUrl,
		timeoutPeriod: timeoutPeriod,
	}
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

func (client ApiClient) FetchAllUsers() ([]User, error) {
	url := fmt.Sprintf("%s/users", client.baseUrl)

	http_client := &http.Client{
		Timeout: client.timeoutPeriod,
	}

	resp, err := http_client.Get(url)
	if err != nil {
		fmt.Println("Error fetching users:", err)
		return []User{}, err
	}

	defer resp.Body.Close()
	resp_body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return []User{}, err
	}

	var users []User
	err = json.Unmarshal(resp_body, &users)
	if err != nil {
		fmt.Println("Error unmarshalling response body:", err)
		return []User{}, err
	}

	return users, nil
}

func main() {
	fmt.Println("Hello, World!")
	client := NewApiClient("https://jsonplaceholder.typicode.com/", 10*time.Second)
	statusCode, err := client.FetchUser(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Status Code:", statusCode)
		fmt.Println("User fetched successfully!!")
	}
}
