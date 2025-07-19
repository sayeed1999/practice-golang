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

	assertCorrectMessage(t, got, want)
}

func TestHelloNameProperly(t *testing.T) {
	t.Run("saying hello when name is empty", func(t *testing.T) {
		got := HelloName("")
		want := "Hello, Sayeed!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people", func(t *testing.T) {
		got := HelloName("Sifat")
		want := "Hello, Hamim!"

		assertCorrectMessage(t, got, want)
	})
}

// Helper function to avoid code duplication in tests.
func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper() // This tells Go that this function is a helper function for tests.
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
