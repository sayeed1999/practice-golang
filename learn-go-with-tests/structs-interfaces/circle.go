package structsinterfaces

import "math"

type Circle struct {
	radius float64
}

func (c *Circle) Area() int {
	return int(math.Pi * c.radius * c.radius)
}
