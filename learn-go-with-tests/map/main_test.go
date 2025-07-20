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

		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})

	t.Run("unknown word", func(t *testing.T) {
		// dictionary := map[string]string{
		dictionary := Dictionary{
			"test":  "this is a test",
			"test2": "this is another test",
		}

		// got := Search(dictionary, "test")
		_, err := dictionary.Search("test3")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatalf("expected an error")
		}

		if err.Error() != want {
			t.Errorf("got: %s, want: %s", err.Error(), want)
		}
	})
}
