package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying Hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello Chris!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
