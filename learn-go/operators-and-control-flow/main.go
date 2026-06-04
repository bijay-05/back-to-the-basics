package main

import "fmt"

func main() {

	// COMPARISION OPERATORS
	var city string = "kolkata"
	var city2 string = "calcutta"

	fmt.Println("City is same: ", city == city2)

	fmt.Println("City is not same: ", city != city2)

	var a,b int = 5,5
	fmt.Println("a(4) is less than b (5): ", a < b)

	// ARITHMETIC OPERATION

	var x, y string = "foo", "bar"
	fmt.Println("The final value is: ", x + y)

	var i, j float64 = 79.02, 75.66
	fmt.Printf("The final value is: %.2f", i-j)

	var m,n int = 24, 7
	fmt.Println("The remainder is: ", m % n)

	var e int = 1
	e++
	fmt.Println("The increased value is: ", e)

	// LOGICAL OPERATIONS

	var z int = 10
	fmt.Println("This is true: ", ((z < 100) && (z < 200)))
	fmt.Println("This is false: ", ((z < 100) && (z > 200)))

	var f int = 45
	fmt.Println("f is less than 10: ", !(f < 10))

	// ASSIGNMENT OPERATIONS
	z += b
	var o, p = 15, 5
	o -= p // subtracts p from o => o - p
	fmt.Println("This is value after addition assignment: ", z, "subtraction assignment: ", o)

	// BITWISE OPERATIONS

	var q, r int = 12, 25
	fmt.Println("Result after bitwise AND: ", q & r)

	// CONDITIONAL - IF

	if (a == 5) {
		fmt.Println("The value of a is 5. We are inside the conditional")
	}

	if ( b != 5) {
		fmt.Println("The value of b is not 5.")
	} else {
		fmt.Println("The value of b is 5 and we are inside else block")
	}

	if ( a == 10) {
		fmt.Println("The value of a is equal to 10")
	} else if ( a ==5 ) {
		fmt.Println("The value of a is equal to 5")
	} else {
		fmt.Println("The value of a is none")
	}

	// SWITCH STATEMENTS
	switch a {
	case 500:
		fmt.Println("The switch a is 500")
	case 200:
		fmt.Println("The switch a is 200")
	case 5:
		fmt.Println("The switch a is 5")
	default:
		fmt.Println("The switch a is default")
	}

	switch a {
	case 500:
		fmt.Println("The switchf a is 500")
	case 200:
		fmt.Println("The switchf a is 200")
	case 5:
		fmt.Println("The switchf a is 5")
		fallthrough
	case 20:
		fmt.Println("The switchf a is 20")
		fallthrough
	default:
		fmt.Println("The switchf a is default")
	}

	switch {
	case a == 500:
		fmt.Println("The switchtt a is 500")
	case a == 200:
		fmt.Println("The switchtt a is 200")
	case a == 5:
		fmt.Println("The switchtt a is 5")
	default:
		fmt.Println("The switchtt a is default")
	}

	for d := 1; d <= 5; d++ {
		fmt.Println("This is the value of d in loop: ", d * d)

		if d == 3 {
			break
		}
	}

	for d := 1; d <= 5; d++ {
		if d == 3 {
			continue
		}
		fmt.Println("This is the value of d in loop: ", d * d)
	}
}