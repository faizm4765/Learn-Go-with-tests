package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello world!"
	if got != want {
		t.Errorf("got %s want %q", got, want)
	}
}
