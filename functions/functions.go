package main

import "fmt"

/*
signature:
func <functionName>(<parameters>) <returnType> {
	<functionBody>
}
*/

// add numbers and return result function

func addNums(a int, b int) int {
	return a + b
}

/*
add, sub two numbers and return multiple results

	function
*/
func addSub1(a int, b int) (int, int) {
	return a + b, a - b
}

func addSub2(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

/*
Variadic function
accepts variable parameters in a function

func <funcName>(<parameter> type, ...type) {
	<func body>
}
*/

func nums(num ...int) int {
	for _, val := range num {
		fmt.Println("numbers are: ", val)
	}
	return 0
}

func studentData(name string, subjects ...string) {
	fmt.Println("Hey ", name, "your subjects are: ")
	for index, sub := range subjects {
		fmt.Println("sub ", index, ": ", sub)
	}
}

// recursive function

func fact(n int) int {
	if n == 1 {
		return 1
	}
	return n * fact(n-1)
}

/*
Higher Order Function
Defer statement
*/

func main() {
	result := addNums(2, 3)
	fmt.Println("sum of 2 & 3 is: ", result)

	r1, r2 := addSub1(3, 5)
	fmt.Println("result of addition of 3 & 5 is : ", r1, ", result of sub of 3 & 5 is: ", r2)

	sum, diff := addSub2(4, 8)
	fmt.Println("result of sum of 4 & 8 is : ", sum, ", result of diff of 4 & 8 is: ", diff)

	nums(10, 20, 30)

	studentData("prasanna", "maths", "physics", "chemistry")

	fmt.Println("result of factorial of 5 is: ", fact(5))

	// Anonymus function 1
	x := func(n int) int {
		return n * n
	}
	fmt.Println("calling anonymus function to square a number: ", x(2))
	fmt.Printf("Type of x is : %T\n", x)
	// Anonymus function 2
	y := func(n int) int {
		return n * n
	}(3)
	fmt.Printf("Type of y is : %T\n", y)
	fmt.Printf("Val of y is: %d\n", y)
}
