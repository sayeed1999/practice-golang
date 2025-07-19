package main

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("z", 5)
	expected := "zzzzz"
	assertCorrectMessage(t, repeated, expected)
}

// Benchmark output between the inefficient and efficient implementations !!!
// BenchmarkRepeat-8                8557082               137.9 ns/op
// BenchmarkRepeatFast-8           36490062                34.83 ns/op

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("z", 5)
	}
}

func BenchmarkRepeatFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatFast("z", 5)
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
