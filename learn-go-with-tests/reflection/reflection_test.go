package reflection_test

import (
	"reflect"
	"reflection"
	"testing"
)

func TestWalk_HappyPath(t *testing.T) {

	expected := "Chris"
	var got []string // for spying that the function is called

	x := struct {
		Name string
	}{
		expected,
	}

	reflection.Walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got: %d, want: %d", len(got), 1)
	}
}

func TestWalk_SeriesOfTests(t *testing.T) {
	tests := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two strig fields",
			struct {
				Name string
				City string
			}{
				"Chris", "London",
			},
			[]string{"Chris", "London"},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var got []string // to spy on that the function is called

			reflection.Walk(test.Input, func(x string) {
				got = append(got, x)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
