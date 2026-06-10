# Pointers

When a program runs, it primarily interacts with RAM (Random Access Memory). Every time you declare a variable, the system allocates memory dynamically based on the data type of that variable. As a result, each execution of the program might assign different memory addresses to variables.

A pointer is a special type of variable that holds the memory address of another variable. This means that instead of storing a direct value, the pointer stores the location where the value is kept, allowing you to access or modify it directly.

> [!Important]
> Memory Addresses can differ between program runs due to dynamic memory allocation. This variability is a normal behaviour in memory management.

```go
x := 1
ptrToX := &x
```

In this code snippet, the variable `x` is assigned a value and stored at a memory addres, for example, `0x0301`. When we declare the pointer variable `ptr`, it holds the memory address of `x` (in this case, `0301`)

To summarize, a pointer is a variable that stores the address of another variable. This powerful concept not only allows you to keep track of where the data is stored but also provides direct access to those memory locations, enabling you to read or modify the value as needed.

## Address and Dereference Operator ( `&` and `*`)

A pointer in Go holds the memory address of a variable. To obtain this memory address, you use the Address Operator, represented by the ampersand (&). Conversely, the Dereference Operator, denoted by an asterisk (\*), accesses the value stored at that memory address.

```go
x := 77
&x = 0x0301 // address operator returns the memory address of x
*0x0301 = 77 // dereference operator retrieves the value stored at memory address
```

## Declaring and Initializing a Pointer

### Declaring a Pointer

A Pointer holds the memory address of a variable.

```go
var <pointer_name> *<data_type>

var ptr_i *int
var ptr_s *string
```

### Initialising a Pointer

Once a pointer is declared, you must initialize it by assigning the memory address of an existing variable.

**Method 1: Using the Address Operator &**

```go
var <pointer_name> *<data_type> = &<variable_name>

i := 10
var ptr_i *int = &i
```

**Method 2: Type Inference**
Go supports type inference, which allows you to omit the explicit data type.

```go
s := "hello"
var ptr_s = &s
```

**Method 3: Shorthand Declaration Operator**

```go
s := "hello"
ptr_s := &s
```

### Dereferencing a Pointer

Dereferencing a pointer means accessing the value stored in the memory address that the pointer holds.

> [!Important]
> By placing the dereference operator (an asterik) before a pointer, you gain access to the value stored in the pointer's referenced memory location. Moreover, you can modify the value by using the dereference operator in an assignment.

```go
*<pointer_name>
*<pointer_name> = <new_value>

// dereferencing
x := 77
var ptr *int := &x // ptr stores an address

// when we dereference ptr and assign a new value, x is updated
*ptr = 100
```

## Passing By Value in Functions

We now explore the concept of passing arguments by value in functions. There are two primary methods for passing arguments: passing by value and passing by reference.

Passing by value means that when you pass a variable to a function, its value is copied to a new memory location. Any modifications made to the parameter within the function affect only the copy, leaving the original value unchanged. Basic data types such as integers, floats, booleans, and strings are all passed by value.

> [!Important]
> When a variable is passed by value, its original location in memory remains unchanged even if the function alters the copied parameter.

```go
func modify(a int) {
    a += 100
}

func main() {
    a := 10
    fmt.Println(a) // Output: 10
    modify(a)
    fmt.Println(a) // Output: 10, since a is passed by value
}
```

## Passing By Reference in Functions

We will explore how to pass variables by reference in functions using pointers. Unlike passing by value, where a copy of the variable is sent to the function, passing by reference allows a function to modify the original variable by using its address.

Golang supports pointers, which let you pass the memory address of a variable to a function. When you pass a pointer, any changes made inside the function affect the original variable.

```go
func modify(s *string) {
	*s = "hello world"
}

func main() {
	a := "hello"
	fmt.Println(a)
	modify(&a)
	fmt.Println(a)
}
```

> [!Caution]
> Another key point in Golang is that slices and maps are passed by reference by default. This means modifications inside the function will affect the original data structure.

### Slices and Maps are Reference Types

A slice in Go is a reference type. When you pass the slice to a function, it points to the same underlying array. Hence, changes inside the function are reflected in the original slice.
