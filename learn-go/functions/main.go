package main

import "fmt"

func sumTwoNumber(a int, b int) int {
	return a + b
}

func printGreeting(name string) {
	fmt.Println("Hello, ", name)
}

// multiple return parameters
func operation(c int, d int) (int, int) {
	return c + d, c - d
}

// variadic function
func sumNumbers(numbers ...int) int {
	sum := 0
	for _, value := range(numbers) {
		sum += value
	}
	return sum
}

// variadic function with multiple parameters
func printDetails(student string, subjects ...string) {
	fmt.Println("Hey ", student, ", here are your subjects:")
	for _, sub := range(subjects) {
		fmt.Printf("%s, ", sub)
	}
}

// factorial with recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n - 1)
}

// high order functions with basic circle property functions
func calcArea(r float64) float64 {
	return 3.14 * r * r
}

func calcPerimeter(r float64) float64 {
	return 2 * 3.14 * r
}

func calcDiameter(r float64) float64 {
	return 2 * r
}

func printResult(radius float64, calcFunction func(r float64) float64) {
	result := calcFunction(radius)
	fmt.Println("Result: ", result)
	fmt.Println("Thank you!")
}

func getFunction(query int) func(r float64) float64 {
	queryToFunc := map[int]func(r float64) float64{
		1: calcArea,
		2: calcPerimeter,
		3: calcDiameter,
	}
	return queryToFunc[query]
}

// defer statement
func printName(name string) {
	fmt.Println("This is name: ", name)
}

func printRollNo(rno int) {
	fmt.Println("This is roll number: ", rno)
}

func printAddress(address string) {
	fmt.Println("This is the address: ", address)
}


func main() {
	fmt.Println("This is the sum of 1 and 2: ", sumTwoNumber(1,2))

	printGreeting("Subu")

	sum, difference := operation(5,2)
	fmt.Println("Sum and difference: ", sum, difference)

	// CALLING VARIADIC FUNCTION
	fmt.Println("The sum of numbers: ", sumNumbers(10,20,30,40))

	// CALLING VARIADIC FUNCTION WITH MULTIPLE PARAMETERS
	printDetails("Subu", "Physics", "Chemistry")

	fact := factorial(3)
	fmt.Println("\nThe factorial of 20: ", fact)

	// STORING ANONYMOUS FUNCTION IN A VARIABLE
	x := func(l int, m int) int {
		return l * m
	}
	fmt.Printf("%T\n", x)
	fmt.Println(x(20,30))

	// HIGH ORDER FUNCTIONS

	var query int
	var radius float64

	fmt.Println("Enter the radius of the circle: ")
	fmt.Scanf("%f", &radius)
	fmt.Printf("Enter \n 1 - Area \n 2 - Perimeter \n 3 - Diameter: ")
	fmt.Scanf("%d", &query)

	// if query == 1 {
	// 	fmt.Println("Result: ", calcArea(radius))
	// } else if query == 2 {
	// 	fmt.Println("Result: ", calcPerimeter(radius))
	// } else if query == 3 {
	// 	fmt.Println("Result: ", calcDiameter(radius))
	// } else {
	// 	fmt.Println("Invalid Query")
	// }

	// USING HIGH-ORDER FUNCTIONS
	printResult(radius, getFunction(query))

	// DEFER STATEMENT
	printName("Jason Bourne")
	defer printRollNo(32)
	printAddress("Biratnagar")
}