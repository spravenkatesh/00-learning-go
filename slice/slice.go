package main

import "fmt"

func main() {

	// initialize a slice
	slice0 := [] int {1,2,3}
	fmt.Println("initialized slice: ", slice0)

	// creating slice from an array
	fruits := [3] string {"apple", "grapes", "oranges"}
	slice := fruits[1:3]
	fmt.Println("slice from fruit is: ", slice)

	//selecting all elemnts of slice
	slice = fruits[ : ]
	fmt.Println("default slice from fruit is: ", slice)

	//selecting all elemnts of slice
	slice = fruits[ : 2]
	fmt.Println("slice from fruit except last element is: ", slice)

	// creating empty slice
	e_slice := make([] int, 3, 4)
	fmt.Println("creating an empty slice: ", e_slice)

	// modifiying values in slice
	e_slice[0] = 10
	fmt.Println("assigned value 10 to an empty slice at index 0", e_slice)

	// append 2 slice
	a_slice := append(slice, "berries", "melons")
	fmt.Println("appneded slice is: ", a_slice)
	fmt.Println("length of appened slice is: ", len(a_slice), ", capacity of appened slice is: ", cap(a_slice))

	// deleting element from slice
	i := 1
	slc_0 := a_slice[:i]
	slc_1 := a_slice[i+1:]
	fmt.Println("new slice by deleting element at index 1 is: ", append(slc_0, slc_1...))

	//copying 2 slices
	e_slc := make([] int, 3, 4)
	slc2 := [] int {2, 3, 4, 5}
	ret := copy(e_slc, slc2)
	fmt.Println("copying 2 slices: ", e_slc, "with length: ", ret)

	// looping over slices
	for index, value := range e_slc {
		fmt.Println(index, "=>", value)
	}

	// looping over slices ex:2
	for _, value := range e_slc {
		fmt.Println("vaues without index is: ", value)
	}

}