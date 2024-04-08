package main

import "fmt"

func Hello(name string, language string) string {
	if name == "" {
		return "Hello world!"
	}

	if language == "" {
		return "Hello " + name + "!"
	}

	if language == "Spanish" {
		return "Hola " + name + "!"
	}

	return ""
}

func main() {
	fmt.Println(Hello("faiz", ""))
}
