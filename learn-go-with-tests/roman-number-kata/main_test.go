package main

import "testing"

func TestConvertToRoman(t *testing.T) {
	t.Run("1 should return I", func(t *testing.T) {
		got := ConvertToRomain(1)
		want := "I"
		assertString(t, got, want)
	})

	t.Run("2 should return II", func(t *testing.T) {
		got := ConvertToRomain(2)
		want := "II"
		assertString(t, got, want)
	})

	t.Run("3 should return III", func(t *testing.T) {
		got := ConvertToRomain(3)
		want := "III"
		assertString(t, got, want)
	})

}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%v', want '%v'", got, want)
	}
}
