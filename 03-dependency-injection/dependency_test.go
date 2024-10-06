package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	bytes := bytes.Buffer{}
	got := ""
	want := "Hello, Faiz"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
