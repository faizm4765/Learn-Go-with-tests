package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(w io.Writer) {
	fmt.Fprintln(w, 3)
	fmt.Fprintln(w, 2)
	fmt.Fprintln(w, 1)
	fmt.Fprintln(w, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
