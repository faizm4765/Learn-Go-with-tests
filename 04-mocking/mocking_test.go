package main

import (
	"bytes"
	"fmt"
	"testing"
)

type MockSleeper interface {
	Sleep()
}

type DummySleeper struct {
	Calls int
}

func (s *DummySleeper) Sleep() {
	fmt.Println("Dummy sleep call ", s.Calls)
	s.Calls++
}

func TestCountdown(t *testing.T) {
	bytes := bytes.Buffer{}
	s := &DummySleeper{}
	Countdown(&bytes, s)
	expected := `3
2
1
Go!
`

	got := bytes.String()
	if got != expected {
		t.Errorf("Got %q Expected %q", got, expected)
	}

	actualSleepCalls := s.Calls
	expectedSleepCalls := 3
	if s.Calls != expectedSleepCalls {
		t.Errorf("Got %q Expected %q", actualSleepCalls, expectedSleepCalls)
	}
}