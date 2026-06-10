package main

import "fmt"

func modify(stu *string) {
	*stu = "Hello World"
}

func main() {
	i := "helo"

	// print the type and value of the address of i
	fmt.Printf("%T %v\n", &i, &i) // => *int 0x2343
	// the type of &i as *int (a pointer to an integer)

	// dereference the address to obtain and print the value stored at that address
	fmt.Printf("%T, %v\n", *(&i), *(&i))

	// Declaring pointer
	var r *int
	var s *string
	fmt.Println(r) // => uninitialized pointers in Go have the nil value
	fmt.Println(s)

	// Pointer Dereferencing
	stri := "hello"
	fmt.Printf("%T %v\n", stri, stri)
	
	ps := &stri
	*ps = "Hello World"
	fmt.Printf("%T %v\n", stri, stri)

	// Passing By Reference
	cu := "Hello"
	modify(&cu)
	fmt.Println("The value of cu after modify: ", cu)
}