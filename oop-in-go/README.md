# Learn OOP in Go from Someone who migrated all his OOP knowledge from Dotnet to Go

## Myth 1: Go doesn't support OOP principles

Coming from a Dotnet background, I was successful practicing most of my OOP knowledge of Dotnet into Go that I often use in day-to-day software development.

The rest? Okay, not all textbook contepts we need equally while building real software.

## Myth 2: Go only supports Composition, not Inheritance

Truth: Go supports both inheritance & composition. It's even better, it solves diamond problem.

## Let's Go To Practice

First create a new go project, define a parent schema & add one method.

```go
package main

type Parent struct {
	Name     string
	Property string
}

func (p *Parent) HasProperty() bool {
	return len(p.Property) > 0
}
```

Then write a main function to test its behavior.

```go
func main() {
	parent := Parent{Name: "Parent Name", Property: "Some Property"}
	println(parent.Name)          // Output: Parent Name
	println(parent.HasProperty()) // Output: true
}
```

### Inheritance

Now define a child schema. Add this code below the parent schema definition to inherit it.

```go
type Child struct {
	Parent
}
```

Notice here, we are not saying `Parent Parent`, we are saying `Parent`.

In Golang, when we write,
```go
type Child struct {
    Parent Parent
}
``` 
it becomes a `HAS-A` relationship. Child contains a prop `Parent` of type `Parent`.

But, when we write,
```go
type Child struct {
    Parent
}
```
it becomes a `IS-A` relationship. Child implicitly inherits all properties and methods of `Parent`.

Now create a instance of child in main function and see you can access all properties & methods of the parent from it.

![image](oop%20in%20go.png)

### Polymorphism > Method Overriding

Now, let's override property of `Parent` in `Child`.

Add the below code under `Child` struct.
```go
func (c *Child) HasProperty() bool {
	return false // consider child does not have any property yet
}
```

And add the below code in main function.
```go
	child := Child{Parent: parent}
	child.Name = "Child"         // This will override the Name field of Parent
	println(child.Name)          // Output: Parent Name
	println(child.HasProperty()) // Output: false
```

### Go is not Traditional OOP Language

When people say this, they mean GoLang doesn't 100% mirror to other OOP languages.

For instance, in Go, **method overloading** is not supported.

But that limitation doesn't restrict us appyling OOP paradigms we need in day-to-day engineering.

Rather, Go is designed to perform even better than most OOP languages!

## Go allows multi-inheritance, what Dotnet, Java don't

In Dotnet, Java - you can inherit multiple interfaces, but at most one class only.

But in Go, we can inherit as much classes (structs) as we want via **struct embedding**.

Modify the above code to this -

```go
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
```

Now inside main function, try accessing the properties of both `Parent` & `Student` from an instance of `Child`.

```
