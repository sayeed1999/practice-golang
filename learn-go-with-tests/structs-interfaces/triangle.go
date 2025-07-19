package structsinterfaces

type Triangle struct {
	base   int
	height int
}

func (t *Triangle) Area() int {
	return int(0.5 * float32(t.base) * float32(t.height))
}
