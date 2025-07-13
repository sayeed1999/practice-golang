package main

import "effective-go/ring"

type T struct {
	Name  string // name of the object
	Value int    // its value
}

func main() {
	x := 1
	y := 2

	// ********* Order of spaces in arithmetic ops without parenthesis **********

	z := (x << 1) + (y << 2)
	println("Sum:", z)

	//  Note: - in golang >> space & no space has a order of priority if there is no parenthesis!
	z = x<<1 + y<<2
	println("Updated Sum:", z)

	// ********* Package naming conventons **********
	// By convention, packages are given lower case, single-word names;
	// there should be no need for underscores or mixedCaps.

	// *Another convention is that the package name is the base name of its source directory;
	// the package in src/encoding/base64 is imported as "encoding/base64" but has name base64,
	// *not encoding_base64 and not encodingBase64.

	newRing := ring.New(5) // Using the ring package's constructor
	println("Created a ring of size:", newRing.Size)
}
