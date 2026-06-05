package main

import "fmt"

func main() {
	// var codes map[string]string
	// codes["en"] = "English" // THIS THROWS ERROR
	// fmt.Println("Codes: ", codes)

	var codes map[string]string = map[string]string{"en": "English", "fr": "French"}

	fmt.Println("Language Codes: ", codes)

	fmt.Println("English Language Code: ", codes["en"])

	language, found := codes["en"]
	fmt.Println("Language and boolean: ", language, found)

	// KEY DOES NOT EXIST, `value` will be the zero value for the map's value type.

	value2, found2 := codes["ge"]
	fmt.Println("Language and boolean: ", value2, found2)

	// ADDING MAP ENTRIES

	codes["ge"] = "German"
	fmt.Println("Updates codes: ", codes)

	// DELETING MAP ENTRIES

	delete(codes, "fr")

	fmt.Println("Codes after deleting french language: ", codes)

	// ITERATING OVER MAP ENTRIES
	for key, value := range codes {
		fmt.Println(key, " => ", value)
	}

	// TRUNCATING MAP ENTRIES
	codes = make(map[string]string)
	fmt.Println("Codes after re-initializing: ", codes)
}