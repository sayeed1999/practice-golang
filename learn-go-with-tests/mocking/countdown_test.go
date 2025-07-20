package main

import (
	"bytes"
	"testing"

	"countdown.go/sleeper"
)

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := &sleeper.SpySleeper{}

	CountDown(buffer, sleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if sleeper.Calls != 3 {
		t.Errorf("got %q sleep calls, want %q", got, want)
	}
}
