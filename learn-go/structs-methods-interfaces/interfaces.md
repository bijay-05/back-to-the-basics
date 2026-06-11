# Interfaces

A feature that enables flexible and decoupled code design through method signatures without implementation.

An interface defines a set of method signatures without providing their implementations. This design paradigm creates a blueprint for functionality. Types, such as structs, implement interfaces implicitly simply by defining the required methods. This approach eliminates the need for explicit declarations or keywords such as `implements`.

- An interface specifies a method set and is a powerful way to introduce modularity
- Interface is like a blueprint for a method set
- They describe all the methods of a method set by providing the function signature for each method

An interface in Go is defined much like a struct, using the `type` keyword followed by the `interface` name and the interface keyword. Inside the curly braces, list the method signatures that any implementing type must provide.

```go
type FixedDeposit interface {
    getRateOfInterest() float64
    calcReturn() float64
}
```

> No variable declarations or method definitions are included in the interface—only method signatures.

## Implicit Implementation of Interfaces

Go’s interface implementation is implicit. This means that a type automatically implements an interface if it provides all the methods declared in the interface with matching signatures. For instance, if you have a struct that defines methods corresponding to those in the `FixedDeposit` interface, it implements that interface without any additional syntax requirements.

## Defining the Interface

We begin by defining an interface named shape. This interface requires any implementing type to provide two methods: area and perimeter. Both methods must return a value of type float64, ensuring a consistent design across all shapes.

```go
package main

import "fmt"

type shape interface {
  area() float64
  perimeter() float64
}
```

> Any type that implements these two methods effectively conforms to the `shape` interface.

## Implementing the Interface with a Square

Let's see how to implement the `shape` interface using a `square` struct. The struct includes a single field, `side`, and implements the required `area` and `perimeter` methods. Notice how the method signatures exactly match those defined in the interface, ensuring proper implementation.

```go
type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) perimeter() float64 {
	return 4 * s.side
}
```

> [!Important]
> The `square` type implements the `shape` interface by providing both the `area` and `perimeter` methods, which is a fundamental requirement in Golang for interface satisfaction.

## Implementing the Interface with a Rectangle

```go
type rect struct {
	length, breadth float64
}

func (r rect) area() float64 {
	return r.length * r.breadth
}

func (r rect) perimeter() float64 {
	return 2*r.length + 2*r.breadth
}
```

Both methods in the rect `struct` fulfill the exact signatures required by the `shape` interface, reinforcing the importance of consistency when working with interfaces.

## Using the Interface

To illustrate the power of interfaces, we define a function called `printData` that accepts any type conforming to the `shape` interface. This function prints the area and the perimeter for whichever shape is passed to it. In the `main` function, instances of both `rect` and `square` are created and passed to `printData`, demonstrating how a single function can seamlessly handle different data types.

```go
func printData(s shape) {
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {
	r := rect{length: 3, breadth: 4}
	c := square{side: 5}

	printData(r)
	printData(c)
}
```
