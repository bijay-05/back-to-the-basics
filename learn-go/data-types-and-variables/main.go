package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// This is a single line comment

/*
This is a multi line comment
*/

var globalVariable string = "Global"

func main() {
	var topic, nextTopic string
	var localVariable string = "Local"
	/*var (
		country string = "Nepal"
		nationality string = "Nepalese"
	)*/

	name := "Bijay" // := required for type inference by compiler
	
	var surname string = "Pachhai"
	var age int = 29
	var male bool = true

	topic = "Modern Web Development"
	nextTopic = "Cloud Computing"
	fmt.Print("The topic is: ", topic)
	fmt.Print("The next topic is: ", nextTopic)
	fmt.Print("\nString literal with new line: ", topic, "\n", nextTopic)

    fmt.Println("Hello World", name, surname, "age: ", age, "male: ", male)

		//STRING FORMATTING
	fmt.Printf("Hello World from user: %s %s \n", name, surname)

	// Arrays & slices
	numbers := [4]int{1, 2, 3, 4}
	words := []string{"foo", "bar"}
	floats := []float64{7.0, 9.43, 0.65}

	//Maps
	// In Maps, keys can be of types like strings or integers
	// mapping := map[string]int{
	// 	"x": 30,
	// }

	// numberMap := map[int]int{
	// 	1: 100,
	// }

	keyValueMap := map[string]string{
		"key": "value",
		"bijay": "pachhai",
	}

	fmt.Println("Numbers: ", numbers)
	fmt.Println("Slice of words: ", words)
	fmt.Println("Slice of arrays: ", floats)
	fmt.Println(keyValueMap)

	// VARIABLE SCOPES

	city := "London"

	{
		country := "United Kingdom"
		fmt.Println("Country: ", country)
		fmt.Println("City: ", city)
	}
	// fmt.Println("This is country, but not accessible: ", country)
	fmt.Println("Global and Local Variables: ", globalVariable, "and", localVariable)

	// USER INPUT
	// var username string
	// fmt.Println("Enter your username: ")
	// fmt.Scanf("%s", &username)
	// fmt.Println("Hey there, ", username)

	//++++++++++++++++++++++++++++
	//INPUT ERROR HANDLING
	//++++++++++++++++++++++++++++
	// var a string
    // var b int
    // fmt.Print("Enter a string and a number: ")
    // count, err := fmt.Scanf("%s %d", &a, &b)
    // fmt.Println("count:", count)
    // fmt.Println("error:", err)
    // fmt.Println("a:", a)
    // fmt.Println("b:", b)

	//+++++++++++++++++++++++++++++
  	// TYPE OF DATA VARIABLE
	//+++++++++++++++++++++++++++++
	// var grades int = 42
  //   var message string = "hello world"
  //   var isCheck bool = true
  //   var amount float32 = 5466.54

  //   fmt.Printf("variable grades = %v is of type %T \n", grades, grades)
  //   fmt.Printf("variable message = '%v' is of type %T \n", message, message)
  //   fmt.Printf("variable isCheck = %v is of type %T \n", isCheck, isCheck)
  //   fmt.Printf("variable amount = %v is of type %T \n", amount, amount)

	fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
  //   fmt.Printf("Type: %v \n", reflect.TypeOf("priyanka"))
  //   fmt.Printf("Type: %v \n", reflect.TypeOf(46.0))
  //   fmt.Printf("Type: %v \n", reflect.TypeOf(true))
	var i int = 90
    var f float64 = float64(i)
    fmt.Printf("%.2f\n", f)

	var ft float64 = 45.89
    var it int = int(ft)
    fmt.Printf("%v\n", it)

	var in int = 42
    var s string = strconv.Itoa(in) // convert int to string
    fmt.Printf("%q", s) // "42"
	
	var stt string = "200"
    itt, err := strconv.Atoi(stt)
    fmt.Printf("%v, %T \n", itt, itt)
    fmt.Printf("%v, %T", err, err)

	//++++++++++++++++++++++++++++++
	// CONSTANTS
	//++++++++++++++++++++++++++++++
	const named = "Harry Potter"
    //named = "Hermione Granger" // This will throw compile time error
    fmt.Printf("%v: %T \n", named, named)

}