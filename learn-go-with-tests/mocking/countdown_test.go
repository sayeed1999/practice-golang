package main

import (
	"bytes"
	"reflect"
	"testing"

	"countdown.go/sleeper"
)

func TestCountDown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
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
	})

	t.Run("sleeps before each print", func(t *testing.T) {
		sleeper := &sleeper.SpySleeper{}

		CountDown(sleeper, sleeper)

		want := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !reflect.DeepEqual(sleeper.Calls, want) {
			t.Errorf("got %q sleep calls, want %q", sleeper.Calls, want)
		}
	})
}
