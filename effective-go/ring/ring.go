package ring

// Constructor naming convention:
// as the package name is "ring", the only exported type is "Ring",
// so instead of callin it "NewRing", we simply call it "New".

type Ring struct {
	Size int // size of the ring
}

func New(size int) *Ring {
	if size <= 0 {
		panic("size must be greater than 0")
	}
	return &Ring{Size: size}
}
