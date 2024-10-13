package main

import (
	"fmt"
	"io"
	"os"
)

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s\n", name)
}

func main() {
	Greet(os.Stdout, "Faiz")
}
