# Structs, Methods and Interfaces

## Structs

A user defined data type for grouping related data elements. By bundling different fields - each potentially of a different type - into a single composite data structure, structs allow you to manage and access logically connected information more efficiently.

Consider a scenario where you need to manage data about a student. Instead of handling separate variables for name, grades, roll number, phone number, and address, you can combine these related fields into a single entity called _Student_ struct.

## Declaring and Initialising a Struct

### Declaring a Struct

```go
type <struct_name> struct {
  // list of fields
}
```

Here the keyword `type` creates a new user-defined type. The `struct` keyword indicates that you are defining a struct type, followed by a list of fields enclosed in curly braces. Each field consists of a name and a type, and they are laid out sequentially in memory.

```go
type Circle struct {
  x float64
  y float64
  r float64
}

type Student struct {
  name string
  rollNo int
  marks []int
  grades map[string]int
}
```

> [!Important]
> No values are assigned during declaration, so each field is set to its default value (for example, zero for integers and an empty string for strings) until explicitly initialized.

### Initializing a Struct

After declaring a Struct, you can initialize it using several methods. Each method caters to different cases, ranging from simple declaration to partial initialization with explicit field names.

**1. Initializing with the `var` keyword**

Using the `var` keyword declares a variable of the struct type with all its fields automatically set to their zero values:

```go
var s student
```

**2. Initializing with the `new` keyword**

You can also create an instance of a struct using the `new` keyword, which allocates memory, initializes the fields to their default zero values, and returns a pointer to the struct.

```go
st := new(Student)
```

**3. Initializing with a Struct Literal using Field Names**

For more precise control, you can initialise a struct using a literal with field names to assign initial values explicitly:

```go
st := Student{
  name: "Joe",
  rollNo: 12,
}
```

**4. Initializing with a Struct Literal Without Field Names**

It is possible to initialize a struct without specifying field names, simply by listing the values in the order they were declared. Although valid, this method is less maintainable and generally not recommended for clarity:

```go
st := Student{"Joe", 12}
```

> [!Caution]
> Avoid using struct literals without field names in larger codebases, as they can lead to errors if the struct definition changes.

## Accessing Fields

Access and Modify the fields of a struct using the dot operator. The dot operator allows you to reference a specific field of a struct variable by following the variable name with a dot and then the field name.

> [!Tip]
> Remember that structs cannot be partially assigned. Each field must be explicitly set or initialized.

## Passing Structs to Functions

Pass structs to functions in Go by two different methods: passing by value and passing by reference. These concepts are essential for controlling whether changes within a function affect the original data.

> [!Important]
> Passing by value means that modifications made inside the function do not alter the original struct. This is useful when you want to ensure the integrity of the initial data.

## Comparing Structs

Compare structs, focusing on type compatibility and field equality. Learn how the equality operators ( `==` and `!=` ) behave when applied to structs of both the same and different types.

> [!Caution]
> Attempting to compare struct instances of different types results in the error. Ensure that both structs are of the same type before comparing them.

## Methods

Let's see how methods extend the functionality of structs.

A method resembles a function but includes an extra element: the receiver parameter. This parameter, specified immediately after the func keyword, indicates the type the method is associated with. Essentially, the receiver is the instance of the type on which the method operates.

```go
func (<receiver>) <method_name>(<parameters>) <return_params> {
  // CODE
}
```

You can think of the receiver as the instance bound to a method. When there is a close association between a function and a struct, implementing the function as a method is natural. In essence, a method is a function with a designated receiver, which can be a named type (e.g., Circle) or a pointer to that type.

```go
func (c Circle) area() float64 {
    // code
}

func (c *Circle) area() float64 {
    // code
}
```

> [!Important]
> Using pointer receivers allows methods to modify the original struct, while value receivers work on a copy of the struct. Choosing between pointer and value receivers is crucial: Use pointer receivers when your method needs to modify the receiver; use value receivers when the method only needs to read data.

## Method Sets

Collections of methods associated with a given data type. Method sets allow you to encapsulate functionality and define specific behaviours for your custom types.

```go
type Student struct {
	name   string
	grades []int
}

func (s *Student) displayName() {
	fmt.Println(s.name)
}

func (s *Student) calculatePercentage() float64 {
	sum := 0
	for _, grade := range s.grades {
		sum += grade
	}
	// Calculates the average grade. The original multiplication by 100 cancels out.
	return float64(sum) / float64(len(s.grades))
}
```
