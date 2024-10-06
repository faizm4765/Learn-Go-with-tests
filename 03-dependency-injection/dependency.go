package main

import "fmt"

func Greet(name string) {
	fmt.Printf("Hello, %s", name)
}

func main() {
	fmt.Printf("Learning dependency injection\n")
	Greet("Faiz")
}
