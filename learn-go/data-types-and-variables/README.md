# Data Types and Variables

> [!Important]
> Remember that the `fmt.Print` function does not append a newline by default. When printing multiple items consecutively, they will appear without any line breaks. `fmt.Println` adds new line at the end of each argument. `fmt.Printf` for formatting output.

### Common Format Specifiers

| **Format Specifier** | **Description**                        |
| -------------------- | -------------------------------------- |
| `%v`                 | Default format for the value           |
| `%T`                 | Data type of the value                 |
| `%d`                 | Decimal Integer                        |
| `%c`                 | Character                              |
| `%q`                 | String enclosed in quotes              |
| `%s`                 | plain string                           |
| `%t`                 | Boolean value                          |
| `%f`                 | Floating-point number                  |
| `%.2f`               | floating numbers upto 2 decimal places |

> [!Important]
> Always ensure that the value you assign matches the declared variable type. This prevents compile-time errors and maintains the integrity of your code.

### Reassigning Variables

Variables declared with shorthand notation can be reassigned new values of the same type. Take the following example:

```go
package main

import "fmt"

func main() {
    name := "Lisa"
    name = "Peter"
    fmt.Println(name)
}
```

Since `name` was inferred as a string during its initial declaration, it can only be reassigned a string value. Attempting to assign a different type will result in a compile-time error.

> [!Caution]
> Reassign variables only with values of the same type. Mixing data types will lead to compile-time errors due to Go's strict type safety.

## Variable Scope

n Go, variable scope defines the region in a program where a variable can be accessed or referenced. Scopes in Go are determined by code blocks, which are denoted by matching pairs of curly braces `{ }`. An outer block can include one or more inner blocks. While inner blocks can use variables declared in their outer blocks, outer blocks cannot access variables declared within their inner blocks.

```go
{
  // Outer block
  {
    // Inner block
  }
}
```

> [!Important]
> Remember, outer blocks cannot access variables declared in their inner blocks.

## Zero Values

In Golang, every variable declared without an explicit initialization automatically receives a default value known as its zero value. These zero values differ based on the variable’s data type. Understanding how zero values work is essential for building robust Go applications and can help prevent common programming errors.

When a variable is declared without an initialization value, Go assigns the following defaults:

| **Data Type** | **Value** |
| ------------- | --------- |
| bool          | `false`   |
| int           | `0`       |
| float64       | `0.00`    |
| string        | `""`      |

For more complex types — such as pointers, functions, slices, maps, channels, and interfaces—the zero value is `nil`. These topics will be discussed in more detail in later sections.

> [!Important]
> Zero values are not only default assignments; they also help in determining whether a variable has been explicitly initialized. This behaviour is especially useful when using conditional checks or when passing variables to functions.

## User Input

we’ll explore how to read input from a user in Go using the fmt package. One common technique is to use the Scanf function, which requires a format string and a sequence of variables (passed with the address operator) where the input will be stored.

```go
fmt.Scanf("<format specifier>", object_arguments)

// READ A STRING USING THE %s SPECIFIER
fmt.Scanf("%s", object_arguments)
```

### Handling Input Errors

The fmt.Scanf function returns two values: the number of arguments successfully read and an error if one occurred. This is useful when you need to handle input errors.

## Find the Type of Variable

Determine the data type of a variable in Go using `%T` and the reflect package's `TypeOf` function.
We will cover two primary techniques: using the `%T` format specifier with `fmt.Printf` and using the reflect package's `TypeOf` function, which can also handle literal values.

## Converting between Types

How to perform type conversion (also known as type casting) in Go by converting variables from one data type to another. While converting variables is straightforward, be aware that the actual value may change after conversion due to differences in data representation.

Covering integer to float, float to integer, and string conversions using the strconv package.

> [!Important]
> When converting an integer to a float, the numeric value remains the same, but may be represented differently with added precision.

> [!Caution]
> Keep in mind that converting a float to an integer results in a loss of precision. Ensure this behavior is acceptable for your application’s use case.

### String conversion with `strconv` package

Go’s built-in package, strconv, makes it easy to convert between strings and integers. Two frequently used functions in this package are:

- `Itoa`: Converts an integer to a string
- `Atoi`: Converts a string to an integer and returns both the integer value and an error if the conversion fails.

To convert a string to an integer, use the Atoi function, which returns both the converted integer and an error value. When the string represents a valid integer, the error will be nil.
