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
		{"2 gets converted to II", 2, "II"},
		{"3 gets converted to III", 3, "III"},
		{"4 gets converted to IV", 4, "IV"},
		{"5 gets converted to V", 5, "V"},
		{"6 gets converted to VI", 6, "VI"},
		{"7 gets converted to VII", 7, "VII"},
		{"8 gets converted to VIII", 8, "VIII"},
		{"9 gets converted to IX", 9, "IX"},
		{"10 gets converted to X", 10, "X"},
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
