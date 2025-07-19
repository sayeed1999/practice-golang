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
// For helper functions, it's a good idea to accept a testing.TB which is an interface
// that *testing.T and *testing.B both satisfy.

// t.Helper() is needed to tell the test suite that this method is a helper. By doing this,
// when it fails, the line number reported will be in our function call rather than inside
// our test helper. This will help other developers track down problems more easily.
func assertCorrectMessage(t testing.TB, got, want string) {
	// ** comment this t.Helper() out to see why it's important for failing tests! **
	t.Helper() // This tells Go that this function is a helper function for tests.
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
