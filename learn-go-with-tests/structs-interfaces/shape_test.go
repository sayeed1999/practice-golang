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

func checkArea(t testing.TB, shape Shape, want int) {
	t.Helper()
	area := shape.Area()
	if area != want {
		t.Errorf("got %v, want %v", area, want)
	}
}
