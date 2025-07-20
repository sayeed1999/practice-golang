package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{
		"test":  "this is a test",
		"test2": "this is another test",
	}

	got := Search(dictionary, "test")
	want := "this is a test"

	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}
