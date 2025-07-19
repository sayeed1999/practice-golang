package structsinterfaces

import "testing"

func TestArea(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{4, 3}
		// got := rectangle.Area()
		want := 12

		checkArea(t, &rectangle, want)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		// got := circle.Area()
		want := 314

		checkArea(t, &circle, want)
	})
}

// Table Driven Tests!!!
// - Run a series of testcases on a series of expected values !!!
func TestAreaTableDriven(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  int
	}{
		{&Rectangle{3, 4}, 12},
		{&Circle{10}, 314},
		{&Triangle{12, 6}, 36},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		want := tt.want

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func checkArea(t testing.TB, shape Shape, want int) {
	t.Helper()
	area := shape.Area()
	if area != want {
		t.Errorf("got %v, want %v", area, want)
	}
}
