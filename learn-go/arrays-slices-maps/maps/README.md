# Maps

A map is a robust data structure that stores an unordered collection of key-value pairs. Depending on the language, similar structures might be known as associative arrays in PHP, hash tables in Java, or dictionaries in Python. Maps are ideal for quickly retrieving data based on a key, thanks to their efficient implementation using hash tables.

Consider a map that associates language codes with their respective languages.

## Declaring a Map

To declare a map in Go, use the `var` keyword along with the map type specification. For instance, to declare a map with string keys and integer values,

```go
var lang_map map[string]int
```

> [!Important]
> This syntax initializes a nil map. The zero value of a map in Go is nil, meaning it doesn't contain any keys. Trying to add a key-value pair to a nil map will result in a runtime error.

## Adding Values to a Map

```go
package main

import "fmt"

func main() {
    var codes map[string]string
    codes["en"] = "English"
    fmt.Println(codes)
}
```

```bash
panic: assignment to entry in nil map
```

Since the map is nil , it cannot accept new key-value pairs until it is properly initialized.

## Initializing a Map with Values

```go
codes := map[string]string{"en": "English", "fr": "French"}
```

### Using the `make` function

An alternative to direct initialization is to use the built-in make function. This function allows you to define the map type and, optionally, an initial capacity:

```go
codes := make(map[string]int)
fmt.Println(codes) // => OUTPUT: map[]
```

### Determining the Length of a Map

To obtain the number of key-value pairs in a map, utilize the built-in `len` function.

## Accessing Map Values

To access a map value, refer to its key enclosed in square brackets.

> [!Important]
> It's important to note that indexing a map in Go returns two values: the value itself and a boolean indicating whether the key exists.

```go
value, found := map_name[key]
```

## Adding and Updating Map Entries

Adding a new key-value pair to a mapp is straighforward; simply assign a value to a new key:

```go
package main

import "fmt"

func main() {
    codes := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}
    codes["it"] = "Italian"
    fmt.Println(codes)
}
```

If you assign a new value to an existing key, the map updates that key's value.

## Deleting Map Entries

To remove a key-value pair from a map, use the built-in `delete` function. This function takes tht map and the key to delete as its arguments.

## Iterating Over a Map

You can iterate over a map using the `range` expression, which retrieves both the key and its associated value.

```go
package main

import "fmt"

func main() {
    codes := map[string]string{
        "en": "English",
        "fr": "French",
        "hi": "Hindi",
    }

    for key, value := range codes {
        fmt.Println(key, "=>", value)
    }
}
```

## Truncating a Map

Truncating a map involves clearing all its elements.

- **Method 1 : Deleting Each Key Iteratively**

- **Method 2 : Re-initializing the Map**

```go
package main

import "fmt"

func main() {
    codes := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}
    codes = make(map[string]string)
    fmt.Println(codes)
}
```
