package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		// dictionary := map[string]string{
		dictionary := Dictionary{
			"test":  "this is a test",
			"test2": "this is another test",
		}

		// got := Search(dictionary, "test")
		got, _ := dictionary.Search("test")
		want := "this is a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		// dictionary := map[string]string{
		dictionary := Dictionary{
			"test":  "this is a test",
			"test2": "this is another test",
		}

		// got := Search(dictionary, "test")
		_, err := dictionary.Search("test3")
		want := ErrNotFound.Error()

		if err == nil {
			t.Fatalf("expected an error")
		}

		assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("unknown word", func(t *testing.T) {
		dictonary := make(Dictionary)
		word := "test"
		definition := "this is a test"
		dictonary.Add(word, definition)

		got, _ := dictonary.Search(word)

		assertStrings(t, got, definition)
	})

	t.Run("known word", func(t *testing.T) {
		dictonary := make(Dictionary)
		word := "test"
		definition := "this is a test"
		err := dictonary.Add(word, definition)

		if err != nil {
			t.Fatalf("cannot throw exception")
		}

		err = dictonary.Add(word, definition) // re-adding what's already present!

		if err == nil {
			t.Fatalf("expected an error")
		}

		assertStrings(t, err.Error(), ErrAlreadyPresent.Error())
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}
