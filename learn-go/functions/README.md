# Functions

Functions in Go allow you to structure your programs into reusable blocks of code, making it easier to debug and enhance your applications over time.

Functions are self-contained units of logic that perform specific tasks. By encapsulating operations within a function, you can divide your program into manageable and repeatable segments instead of maintaining one large block of code.

- Self Contained units of code which carry out a certain job
- Help us divide a program into small manageable, repeatable and organisable chunks

At a higher level of abstraction, a function operates by taking one or more inputs, processing them, and then returning an output.

> [!Important]
> Main Benefits of Using functions
>
> - Reusability: Write a function once and call it multiple times throughout your project
> - Abstraction: Hide complex implementation details, allowing you to focus on what the function does rather than how it does it.

```go
func FunctionToAdd(<params>) <return Type> {
  // body of the function
}
```

## Naming Conventions

In Go, function names must adhere to specific naming rules:

- They must begin with a letter or an underscore
- They can include additional letters, digits and underscores
- They are case sensitive
- They cannot contain spaces

## Parameters vs. Arguments

- Parameters are the variable names defined in the function signature
- Arguments are the actual values passed to the function when it is called

```go
// a and b are parameters
func addNumbers(a int, b int) int {
    sum := a + b
    return sum
}
// 2 and 3 are arguments
func main() {
    sumOfNumbers := addNumbers(2, 3)
    fmt.Println(sumOfNumbers)
}
```

## Input Parameters and Return Values

1. **Input Parameters**: These values are passed into the function either by value or by address
2. **Return Parameters**: These define the outputs returned by the function

## Return Types: Multiple, Named and Variadic

There are various function return types in Go, including single and multiple return values, named return parameters and variadic functions.

> Go provides the flexibility to return multiple values from a function. When declaring such a function, you enclose the return types in parentheses.

```go
func operation(a int, b int) (int, int) {
    sum := a + b
    diff := a - b
    return sum, diff
}
```

### Named Return Parameters

By naming the return parameters in the function signature, Go allows you to simplify the function's return statement.

```go
func operation(a int, b int) (sum int, diff int) {
  sum = a + b
  diff = a -b
  return
}
```

> When you name the return parameters, you can use an empty `return` statement to return the current values of those parameters. This approach can make your code cleaner and easier to maintain.

### Variadic Functions

Variadic functions accept a variable number of arguments. In Go, you declare a variadic parameter by prefixing its type with an ellipsis ( `...` ). This feature is useful for functions where the number of inputs may vary.

```go
func <function_name>(param1 string, param2 int, param3 ...type) <return_type>

func sumNumbers(numbers ...int) int
```

Inside the function, the variadic parameter `numbers` behaves like a slice containing all the passed integers.

### Variadic Functions with Multiple Parameters

A variadic function can accept additional parameters, bu the variadic parameter must always be the last in the function signature.

> [!Important]
> The blank identifier `_` effectively acts as a placeholder to ignore any return value you do not need.

## Recursive Functions

A powerful programming technique where a function calls itself, either directly or indirectly, to solve problems by breaking them down into smaller instances of the same problem. This method is particularly effective when dealing with tasks such as calculating mathematical sequences or traversing data structures.

> [!Important]
> Recursion can simplify complex problems by defining a base case and a recursive step. In the factorial function, the base case is defined when n equals 0, because 0! is defined as 1.

## Anonymous Functions

Unnamed functions used for short-lived operations, encapsulating functionality without polluting the global namespace. These functions are declared without an identifier.

- Storing an Anonymous Function in a Variable

```go
x := func(l int, m int) int {
  return l * b
}

multiple := x(20,30) // 60
```

- Directly invoking an Anonymous Function

```go
x := func(l int, m int) int {
  return l * m
}(20,30)

fmt.Println("The value of x: ", x)
```

## High Order Functions

Functions that either accept another function as an argument or return a function as output. High order functions enable modular composition, allowing you to build complex operations from small, focused functions. This approach not only minimizes bugs but also enhances code readability and maintainability.

> [!Important]
> High-order functions allow you to abstract recurring operations, making your code scalable as you incorporate additional properties or shapes.

```go
package main

import "fmt"

func calcArea(r float64) float64 {
	return 3.14 * r * r
}
func calcPerimeter(r float64) float64 {
	return 2 * 3.14 * r
}
func calcDiameter(r float64) float64 {
	return 2 * r
}

func main() {
	var query int
	var radius float64

	fmt.Print("Enter the radius of the circle: ")
	fmt.Scanf("%f", &radius)
	fmt.Printf("Enter \n 1 - area \n 2 - perimeter \n 3 - diameter: ")
	fmt.Scanf("%d", &query)

	if query == 1 {
		fmt.Println("Result: ", calcArea(radius))
	} else if query == 2 {
		fmt.Println("Result: ", calcPerimeter(radius))
	} else if query == 3 {
		fmt.Println("Result: ", calcDiameter(radius))
	} else {
		fmt.Println("Invalid query")
	}
}
```

While this solution works, it can quickly become cumbersome as you add more shape properties. Leveraging high-order functions simplifies the code and enhances scability.

#### Simplification

We introduce two helper functions: one to display the result and another to map the query number to the appropriate calculation function.

```go
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
```

> [!Important]
> High-order functions reduce code redundancy while clarifying the logical separation between input handling and processing logic - making it easier to extend and maintain your application.

## Defer Statement

A powerful feature that postpones a function's execution until the surrounding function returns. While the arguments for a deferred function call are evaluated immediately, the function itself executes later. This behavious ensures that certain operations, such as resource cleanup, occur right before a function exits.

### Working of Defer Statement

Suppose there are three functions: `printName`, `printRollNo` and `printAddress`

- The `printName` function outputs a provided string
- The `printRollNo` function prints an integer
- The `printAddress` function prints a string representing an address

Within the `main` function, the program flow is as follows:

1. The `printName` function is immediately called to print the name
2. The `defer` keyword schedules the `printRollNo` function to execute after `main` finishes its other statements.
3. The `printAddress` function is called to print the address "street-32"
4. Once `main` is ready to return, the deferred `printRollNo(32)` call is executed.
