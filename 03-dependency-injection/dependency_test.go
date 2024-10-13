package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	bytes := bytes.Buffer{}
	Greet(&bytes, "Faiz")
	got := bytes.String()
	want := "Hello, Faiz\n"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
