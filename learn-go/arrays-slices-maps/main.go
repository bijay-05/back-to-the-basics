package main

import "fmt"

func main() {
	// INITIALIZING ARRAY
	var grades [5]int
	var fruits [3]string

	// PRINT EMPTY ARRAY
	fmt.Println("Grades array: ", grades) // => [0 0 0 0 0]
	fmt.Println("Fruits array: ", fruits) // => [ ]

	// EXPLICIT INITIALIZATION OF ARRAY
	// IF THE NUMBER OF PROVIDED VALUES IS LESS THAN THE DECLARED SIZE
	// THE REMAINING ELEMENTS ARE SET TO THE ZERO VALUE
	var gradess [3]int = [3]int{10, 20, 30}

	// SHORT HAND DECLARATION
	marks := [3]int{1, 2, 3}

	// USING ELLIPSES
	// LET THE COMPILER DETERMINE THE ARRAY LENGTH AUTOMATICALLY BY USING ELLIPSES
	// THIS AUTOMATICALLY SETS THE SIZE OF THE names ARRAY TO THE NUMBER OF ELEMENTS PROVIDED
	names := [...]string{"Rachel", "Phoebe", "Monica"} 

	fmt.Println("Gradess: ", gradess)
	fmt.Println("Marks: ", marks)
	fmt.Println("Names: ", names)

	// GETTING THE LENGTH OF AN ARRAY WITH BUILT-IN len() FUNCTION

	fmt.Println("Length of Names: ", len(names))

	// ACCESS NAME AT INDEX 2
	fmt.Println("Name at index 2: ", names[2])

	// UPDATING THE NAME AT INDEX @
	names[2] = "Monika"

	fmt.Println("Updated array names: ", names)

	// LOOPING THROUGH ARRAYS

	for i:=0; i < len(names); i++ {
		fmt.Println("Name at position: ", i, " name: ", names[i])
	}

	// range KEYWORD RETURNS BOTH THE INDEX AND THE ELEMENT
	for index, element := range names {
		fmt.Println(index, "=>", element)
	}

	// MULTI-DIMENSIONAL ARRAYS

	var arr [3][2]int = [3][2]int{
		{1,2},
		{3,4},
		{5,6},
	}
	fmt.Println("two-dimensional array: ", arr)
	fmt.Println("second row and second column element: ", arr[1][1])

	fmt.Println("================SLICES SECTION==================")

	var domes []int = []int{32, 33, 34}
	fmt.Println("Domes slice: ", domes)

	test_arr_1 := [5]int{1,2,3,4,5}
	slice_1 := test_arr_1[1:4]
	fmt.Println("Slice_1: ", slice_1)

	slice_2 := test_arr_1[3:]
	fmt.Println("Slice_2: ", slice_2)

	slice_3 := test_arr_1[:2]
	fmt.Println("Slice_3: ", slice_3)

	// SLICE WITH make FUNCTION
	slice_test := make([]int, 5, 8)
	fmt.Println("Slice Test: ", slice_test)
	fmt.Println("Slice Test length: ", len(slice_test))
	fmt.Println("Slice Test Capacity: ", cap(slice_test))

	// MODIFYING A SLICE
	arr_test_2 := [5]int{10, 20, 30, 40, 50}
    // Create a slice containing indices 0, 1, and 2.
    slice_test_2 := arr_test_2[:3]
    fmt.Println("Before modification:")
    fmt.Println("Array:", arr_test_2)
    fmt.Println("Slice:", slice_test_2)
    
    // Modify the element at index 1 in the slice_test_2.
    slice_test_2[1] = 9000

    fmt.Println("After modification:")
    fmt.Println("Array:", arr_test_2) // ARRAY ALSO CHANGES
    fmt.Println("Slice:", slice_test_2)

	arr_test_3 := [5]int{10, 20, 30, 40, 50}
	slice_test_3 := arr_test_3[1:3]

	fmt.Println("Length and capacity: ", len(slice_test_3), cap(slice_test_3)) // 2 and 4

	fmt.Println("Original Slice: ", slice_test_3)
	fmt.Println("Length: ", len(slice_test_3)) // 2
	fmt.Println("Capacity: ", cap(slice_test_3)) // 4

	slice_test_3 = append(slice_test_3, 200, 300, 500) // EXCEED ORIGINAL CAPACITY OF 4, SO DOUBLES THE CAPACITY

	fmt.Println("Updated Slice: ", slice_test_3)
	fmt.Println("New length: ", len(slice_test_3)) // 5
	fmt.Println("New Capacity: ", cap(slice_test_3)) // 8

	// APPEND ONE SLICE TO ANOTHER

	slice_test_3 = append(slice_test_3, slice_test_2...)
	fmt.Println("Appended Slice: ", slice_test_3)

	// DELETING ELEMENTS FROM A SLICE
	arr_test_4 := [5]int{1,2,3,4,5}
	delete_index := 2
	fmt.Println("Original array: ", arr_test_4)

	slice_test_a := arr_test_4[:2] // ELEMENTS BEFORE INDEX 2
	slice_test_b := arr_test_4[delete_index + 1:] // ELEMENTS AFTER INDEX 2

	final_slice_test := append(slice_test_a, slice_test_b...)
	fmt.Println("New slice after deletion of element at index 2: ", final_slice_test)

	// COPYING ELEMENTS BETWEEN SLICES WITH copy 

	src_slice_1 := []int{10, 20, 30, 40}
	dest_slice_1 := make([]int, 3) // DESTINATION SLICE WITH LENGTH AND CAPACITY 3
	num_elements_copied := copy(dest_slice_1, src_slice_1)
	fmt.Println("Destination Slice: ", dest_slice_1)
	fmt.Println("Number of elements copied: ", num_elements_copied) // => 3 ELEMENTS COPIED AS THE DESTINATION SLICE'S LENGTH IS 3

	// LOOPING THROUGH SLICES

	for index, value := range src_slice_1 {
		fmt.Println(index, "=>", value)
	}
}