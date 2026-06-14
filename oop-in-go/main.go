package main

type Parent struct {
	Name     string
	Property string
}

func (p *Parent) HasProperty() bool {
	return len(p.Property) > 0
}

type Student struct {
	Class string
	Roll  int
}

type Child struct {
	Parent
	Student
}

func (c *Child) HasProperty() bool {
	return false // consider child does not have any property yet
}

func main() {
	parent := Parent{Name: "Father", Property: "Some Property"}
	println(parent.Name)          // Output: Parent Name
	println(parent.HasProperty()) // Output: true

	child := Child{Parent: parent}
	child.Name = "Child"         // This will override the Name field of Parent
	println(child.Name)          // Output: Parent Name
	println(child.HasProperty()) // Output: false

	// Accessing Parent's fields and methods through Child
	child.Name = "Child Name" // This will override the Name field of Parent
	// Accessing Student's fields and methods through Child
	child.Class = "10th Grade"
	child.Roll = 5

	///        Parent       Student
	///		   \             /
	///		    \           /
	///		     \         /
	///		      \       /
	///		       \     /
	///		        \   /
	///		         Child (double inheritance example)
}
