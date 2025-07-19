package arraysnslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := Sum0(numbers)
		want := 15
		assertCorrectMessage(t, got, want)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		got := Sum(numbers)
		want := 10
		assertCorrectMessage(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 6}, []int{3, 4, 5})
	want := []int{3, 6, 12}

	// Note:- Go does not let you use equality operators with slices.
	// You could write a function to iterate over each got and want slice and check their values
	// but for convenience sake, we can use reflect.DeepEqual()
	if !reflect.DeepEqual(got, want) { // got != want {
		t.Errorf("got '%v', want '%v'", got, want)
	}
}

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got '%d', want '%d'", got, want)
	}
}
