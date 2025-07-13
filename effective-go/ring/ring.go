package ring

// Constructor naming convention:
// as the package name is "ring", the only exported type is "Ring",
// so instead of callin it "NewRing", we simply call it "New".

type Ring struct {
	Size  int    // size of the ring
	owner string // owner of the ring, for example "lady"
}

func New(size int) *Ring {
	if size <= 0 {
		panic("size must be greater than 0")
	}
	return &Ring{Size: size}
}

// If you have a field called owner (lower case, unexported),
// the getter method should be called Owner (upper case, exported), not GetOwner.
func (r *Ring) Owner() string {
	return r.owner
}

// SetOwner suits better for the setter method, not GetOwner is redundant, and Owner is simple for getter in Go!
func (r *Ring) SetOwner(owner string) bool {
	if owner == "" {
		panic("owner cannot be empty")
	}
	r.owner = owner
	return true
}
