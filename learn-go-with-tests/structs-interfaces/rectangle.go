package structsinterfaces

type Rectangle struct {
	length int
	width  int
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}
