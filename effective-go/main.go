package main

import (
	"bytes"
	"effective-go/files"
	namedresultparams "effective-go/named-result-params"
	"effective-go/ring"
	"fmt"
	"sync"
)

type T struct {
	Name  string // name of the object
	Value int    // its value
}

type SyncBuffer struct {
	lock   sync.Mutex   // mutex to protect the buffer
	buffer bytes.Buffer // buffer to hold data
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
	var t interface{}
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

	// ************ Named Result Parameters ************
	code, message := namedresultparams.GetStatus()
	fmt.Printf("Status Code: %v & Status Message: %v\n", code, message)
	result, err := namedresultparams.Divide(10, 2)
	fmt.Printf("Division Result: %v, Error: %v\n", result, err)

	// defer executes in LIFO order, e.g.,
	defer fmt.Println("Finished deferring...")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

	// ******** Data allocation ***********
	// // in go, the new keyword allocates memory and zeroes it.
	// p := new(SyncBuffer) // type *SyncBuffer
	// var v SyncBuffer     // type SyncBuffer

	// But some time the zero value isn't good enough and and initializing constructor is necessary.

	// // The expression new(File) and &File{} are equivalent,
	// f1 := new(File)              // type *File
	// f2 := &File{}                // type *File
	// f1 = NewFile(1, "file1.txt") // using the constructor

	// // array, silces, strings
	// arr := []string{"no error", "Eio", "invalid argument"}
	// slices := [...]string{"no error", "Eio", "invalid argument"}
	// map1 := map[string]string{"Enone": "no error", "Eio": "Eio", "Einval": "invalid argument"}

	// Allocate a slice of 100 ints in go
	p1 := make([]int, 100) // type []int
	fmt.Println(p1)
	// Note: make is used to allocate only slices, maps, and channels. And does not return a pointer!

	// Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array.
	// A Read function can therefore accept a slice argument rather than a pointer and a count;

	// Two Dimensional Slices in Go!
	type Transform [3][3]float64 // a 3x3 matrix, an array of arrays
	type LinesOfText [][]byte    // a slice of slices

	// Because slices are variable-length, it is possible to have each inner slice be a different length.
	// That can be a common situation, as in our LinesOfText example: each line has an independent length.
	text := LinesOfText{
		[]byte("Now is a good time!"),
		[]byte("for all gophers"),
		[]byte("to come to the aid of their party."),
	}
	fmt.Println(text)

	trn := Transform{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	fmt.Println(trn)

}

type File struct {
	fd   int    // file descriptor
	name string // name of the file

}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}

	f := new(File) // type *File
	f.fd = fd
	f.name = name
	return f
}

// There are lots of boilerplate in the above ctor, we can simplify it using a composite literal
func NewFile2(fd int, name string) *File {
	if fd < 0 {
		return nil
	}

	f := &File{
		fd:   fd,
		name: name,
	}
	return f // return the pointer to the composite literal
}
