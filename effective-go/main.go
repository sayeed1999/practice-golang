package main

import (
	"effective-go/files"
	"effective-go/ring"
	"fmt"
)

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

	// Another short example is once.Do;
	// once.Do(setup) reads well and would not be improved by writing
	// once.DoOrWaitUntilDone(setup).
	// Long names don't automatically make things more readable.
	// A helpful doc comment can often be more valuable than an extra long name.

	// Mixed Caps!
	// Finally, the convention in Go is to use MixedCaps or mixedCaps rather than underscores to write multiword names.
	// e.g DoOnce instead of do_once.

	// ********* Getter and Setter Naming Conventions **********
	// If you have a field called owner (lower case, unexported),
	// the getter method should be called Owner (upper case, exported), not GetOwner.
	// But SetOwner suits better for the setter method,
	if success := newRing.SetOwner("lady"); success {
		fmt.Println("Ring Owner:", newRing.Owner())
	}

	// Panic! Uncomment the below line to see a panic in action.
	// newRing.SetOwner("") // how to resolve panic in go?

	fileInfo := files.ReadFile()
	fmt.Println()
	fmt.Println(fileInfo)

	// ******** Reverse array 'a' **********
	// Go has no comma operator and ++ and -- are statements not expressions.
	// Thus if you want to run multiple variables in a for you should use parallel assignment
	// (although that precludes ++ and --).
	a := []string{"H", "e", "l", "l", "o", ",", "W", "o", "r", "l", "d", "!"}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println("Reversed Array:", a)

	// ************ Type Switch ************
	t = func() string { return "Hello, World!" }()

	switch t := t.(type) {
	case int:
		fmt.Println("Type is int:", t)
	case string:
		fmt.Println("Type is string:", t)
	case *int:
		fmt.Println("Type is *int:", t)
	case *string:
		fmt.Println("Type is *string:", t)
	default:
		fmt.Println("Unknown type:", t)
	}
}

var t interface{}
