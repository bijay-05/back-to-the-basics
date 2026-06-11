package main

import "fmt"

type Student struct {
	name string
	rollNo int
	marks []int
	grades map[string]int
}

type Circle struct {
	x int
	y int
	radius float64
	area float64
}

func calcArea(c Circle) {
	var area float64
	area = 3.14 * c.radius * c.radius
	fmt.Printf("\nArea of circle: %.2f", area)
}

func calcAreaR(c *Circle) {
	var area float64
	area = 3.14 * c.radius * c.radius
	fmt.Printf("\nArea of circle: %.2f", area)
	(*c).area = area
}

type s1 struct {
	x int
}

type CircleForMethod struct {
	radius float64
	area float64
}
func (c *CircleForMethod) calcAreaMethod() {
	c.area = 3.14 * c.radius * c.radius
	fmt.Printf("\n Area of circle with method: %+v", c)
}

type CircleForValue struct {
	radius float64
	area float64
}
func (c CircleForValue) calcAreaValue() {
	c.area = 3.14 * c.radius * c.radius
	fmt.Printf("\n Area of circle with value: %+v", c)
}

func main() {
	// INITIALIZATION WITH VAR
	var s Student
	fmt.Printf("%+v", s)

	fmt.Printf("\n")
	// INITIALIZATION WITH NEW KEYWORD
	st := new(Student)
	fmt.Printf("%+v", st) // st is a pointer to Student

	// STRUCT LITERAL USING FIELD NAMES
	stt := Student{
		name: "Sibhu",
		rollNo: 12,
	}

	fmt.Printf("\n%+v", stt)

	// MODIFY STRUCT FIELDS

	stt.name = "Shiva"
	fmt.Printf("\n%+v", stt)

	// PASS STRUCT BY VALUE
	var c Circle = Circle{ x:5, y:5, radius:1, area:0}
	fmt.Printf("\nThis is circle: %+v", c)
	calcArea(c)
	fmt.Printf("\nThis is circle after calculating area: %+v", c)

	// PASS BY REFERENCE
	calcAreaR(&c)
	fmt.Printf("\nThis is area after passing reference: %+v", c)

	// COMPARING STRUCTS
	cr := s1{x : 5}
	c1 := s1{x: 6}
	c2 := s1{x: 5}

	if cr != c1 {
		fmt.Println("\nc and c1 have different values")
	}
	if cr == c2 {
		fmt.Println("c is same as c2")
	}

	// METHODS

	// METHOD WITH POINTER RECEIVER
	cir := CircleForMethod{radius: 5}
	cir.calcAreaMethod()
	fmt.Printf("\n With pointer receiver: %+v", cir)

	// METHOD WITH VALUE RECEIVER
	cirv := CircleForValue{radius: 1}
	cirv.calcAreaValue()
	fmt.Printf("\nWith Value Receiver: %+v", cirv)

}