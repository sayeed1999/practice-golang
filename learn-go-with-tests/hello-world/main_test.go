package main

import "testing"

// Note: If you try to run the test without creating a go module, you get an error.
// Error: go: cannot find main module, ...

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World!"

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func TestHelloName(t *testing.T) {
	got := HelloName("Mucktadir Sayem")
	want := "Hello, Mucktadir Sayem!"

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
