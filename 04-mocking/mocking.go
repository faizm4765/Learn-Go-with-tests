package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type DefaultSleeper interface {
	Sleep()
}

type Sleeper struct {
}

func (s *Sleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(w io.Writer, s DefaultSleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}

	fmt.Fprintln(w, "Go!")
}

func main() {
	sleeper := &Sleeper{}
	Countdown(os.Stdout, sleeper)
}

// Actually sleep call nhi krte UTs mein, hum ek method banate jo sleep ko mock krega. Like sleep ki nakal krega wo method phir hum usko use kr lenge UT mein so basially us method ko inject kr re and then we can see what we did is we did mocking followed by dependency injection of that method.
