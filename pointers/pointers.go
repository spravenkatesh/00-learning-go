package main

import "fmt"

func main() {
	var x int = 10
	//referencing
	fmt.Printf("type of x: %T\n, address of x: %v\n", &x, &x)
	//de-referencing
	fmt.Printf("type of x: %T\n, value of x: %v\n", *(&x), *(&x))

	// initializing pointers
	var p_x *int = &x
	fmt.Printf("type of p_x is: %T\naddress of p_x is: %v\nadress of x is: %v\n", &p_x, &p_x, &x)
	var p1_x = &x
	fmt.Printf("type of p_x is: %T\naddress of p_x is: %v\nadress of x is: %v\n", &p1_x, &p1_x, &x)
	p2_x := &x
	fmt.Printf("type of p_x is: %T\naddress of p_x is: %v\nadress of x is: %v\n", &p2_x, &p2_x, &x)

	// usage of de-referencing
	s := "hello"
	fmt.Printf("%T %v\n", s, s)
	ptr_s := &s
	fmt.Printf("%T %v\n", &ptr_s, &ptr_s)
	*ptr_s = "world"
	fmt.Printf("%T %v\n", &ptr_s, *ptr_s)
	fmt.Printf("%T %v\n", &ptr_s, &ptr_s)

}
