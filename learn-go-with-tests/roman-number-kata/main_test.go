package main

import "testing"

func TestConvertToRoman(t *testing.T) {

	// table-based tests
	testcases := []struct {
		Description    string
		Input          int
		ExpectedOutput string
	}{
		{"1 gets converted to I", 1, "I"},
		{"1 gets converted to II", 2, "II"},
		{"1 gets converted to III", 3, "III"},
		{"1 gets converted to IV", 4, "IV"},
		{"1 gets converted to V", 5, "V"},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Description, func(t *testing.T) {
			got := ConvertToRomain(testcase.Input)
			want := testcase.ExpectedOutput
			assertString(t, got, want)
		})
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%v', want '%v'", got, want)
	}
}
