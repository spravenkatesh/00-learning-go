package main

import "fmt"

func darray() {
	// 1-D array
	var marks [3] int = [3] int {10, 20, 40}
	fmt.Println("array declaration type 1: ", marks)
	names := [4] string {"avenger", "yamaha", "honda", "java"}
	fmt.Println("array declaration type 2: ", names)
	fruits := [...] string {"apple", "grapes"}
	fmt.Println("array declaration type 3: ", fruits)
	fmt.Println("length of array type 1: ", len(marks), "length of array type 2: ", names, "length of array type 3: ", fruits)
	for i := 0; i < len(fruits); i++ {
		fmt.Println("element of array are: ", fruits[i])
	}
	for index, element := range(marks) {
		fmt.Println("index ", index, "=>", element)
	}
	// access elements of array by index
	fmt.Println("element at index 0 of array names is ", names[0])

	// 2-D array
	records := [3][2] int {{10, 20}, {40, 50}, {100, 120}}
	fmt.Println("2-D array elemnets are: ", records)
	fmt.Println("index 0 element of 3rd array is: ", records[2][0])
}

func main() {
	darray()
}