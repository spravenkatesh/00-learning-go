package main

import "fmt"

//arithmetic operators

func Add2n(a, b int) int{
	return a + b
}

func Sub2n(a, b int) int {
	return a - b
}

func Mul2n(a, b int) int {
	return a * b
}

func Div2n(a, b float32) float32 {
	return a / b
}

func Mod2n(a, b int) int {
	return a % b
}

func inc(a int) int {
	a++
	return a
}

func dec(a int) int {
	a--
	return a
}

//logical operators

func land(x int) bool {
	return x < 100 && x > 20
}

func lor(x int) bool {
	return x < 100 || x > 20
}

func lnot(x bool) bool {
	return !x
}

// assignment operators

func eq(x, y int) int {
	x += y
	return x
}

func deq(x, y int) int {
	x -= y
	return x
}

func meq(x, y int) int {
	x *= y
	return x
}

func dqeq(x, y int) int {
	x /= y
	return x
}

func moeq(x, y int) int {
	x %= y
	return x
}

// bitwise operation

func band(x, y int) int {
	return x & y
}

func bor(x, y int) int {
	return x | y
}

func bxor(x, y int) int {
	return x ^ y
}

func blsh(x, y int) int {
	return x << y
}

func brsh(x, y int) int {
	return x >> y
}
// driver code

func main() {
	// variable initialization
	var a, b int = 3, 5

	// arithmetic operators
	fmt.Printf("addition of number %v & %v is: %v \n", a, b, Add2n(a, b))
	fmt.Printf("substration of number %v, %v is : %v \n", a , b, Sub2n(a, b))
	fmt.Printf("multiplication of number %v, %v is : %v \n", a , b, Mul2n(a, b))
	fmt.Printf("division of number %v, %v is : %v \n", a , b, Div2n(float32(a), float32(b)))
	fmt.Printf("Modulus of number %v, %v is : %v \n", a , b, Mod2n(a, b))

	// comparison operators
	fmt.Printf("increment operator %v is %v \n", a, inc(a))
	fmt.Printf("decrement operator %v is %v \n", a, dec(a))

	//logical operators

	fmt.Printf("%v is greater than 20 and less than 100 : %v \n", 30, land(30))
	fmt.Printf("%v is greater than 20 or less than 100 : %v \n", 10, lor(10))
	fmt.Printf("complement of %v is : %v \n", true, lnot(true))

	//assignment operators

	fmt.Printf("add and assign value for %v & %v is : %v \n", a, b, eq(a, b))
	fmt.Printf("subtract and assign value for %v & %v is : %v \n", a, b, deq(a, b))
	fmt.Printf("multiply and assign value for %v & %v is : %v \n", a, b, meq(a, b))
	fmt.Printf("divide and assign quotient value for %v & %v is : %v \n", a, b, dqeq(a, b))
	fmt.Printf("modulud and assign remainder value for %v & %v is : %v \n", a, b, moeq(a, b))

	// bitwise operators
	fmt.Printf("bitwise or for %v & %v is : %v \n", a, b, bor(a,b))
	fmt.Printf("bitwise and for %v & %v is : %v \n", a, b, band(a,b))
	fmt.Printf("bitwise xor for %v & %v is : %v \n", a, b, bxor(a,b))
	fmt.Printf("bitwise left-shift for %v & %v is : %v \n", a, b, blsh(a,b))
	fmt.Printf("bitwise right-shift for %v & %v is : %v \n", a, b, brsh(a,b))
}