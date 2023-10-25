package main

import "fmt"

// testing if condition

func tif(name string) {
	if name == "apple" {
		fmt.Println("condition under if matched: ", name)
	}
}

// testig if-else condition

func tifele(name string) {
	if name == "apple" {
		fmt.Println("condition under if matched: ", name)
	} else {
		fmt.Println("condition under else matched: ", name)
	}
}

// testing if-else-if condition

func tifelif(name string) {
	if name == "apple" {
		fmt.Println("condition under if matched: ", name)
	} else if name == "orange" {
		fmt.Println("condition under else-if matched: ", name)
	} else {
		fmt.Println("condition under else matched: ", name)
	}
}

// switch-case

func tsc1(name string){
	switch name {
		case "solar":
			fmt.Println("received expected name", name)
		case "martian", "lunar":
			fmt.Println("received expected name", name)
		default:
			fmt.Println("received unexpected name", name)
	}
}


func tsc2(name string){
	switch {
		case name == "solar":
			fmt.Println("received expected name", name)
		case name == "martian" || name == "lunar":
			fmt.Println("received expected name", name)
		default:
			fmt.Println("received unexpected name", name)
	}
}

func tsc3(name string){
	switch {
		case name == "solar":
			fmt.Println("received expected name", name)
			fallthrough
		case name == "martian" || name == "lunar":
			fmt.Println("received expected name: ", name)
		default:
			fmt.Println("received unexpected name", name)
	}
}

// for loop

func tfor1() {
	for i:= 0 ; i < 10; i++ {
		fmt.Println("square of" ,i ,"is :" , i*i)
	}
}

func tfor2(i int) {
	for ; i < 10; i++ {
		fmt.Println("square of" ,i ,"is :" , i*i)
	}
}

func tfor3(i int) {
	for ; i < 10; i++ {
		if i == 0 {
			continue
		}
		fmt.Println("square of" ,i ,"is :" , i*i)
	}
}

func tfor4(i int) {
	for ; i < 10; i++ {
		if i == 6 {
			break
		}
		fmt.Println("square of" ,i ,"is :" , i*i)
	}
}

func main() {
	tif("apple")
	tifele("grapes")
	tifelif("orange")
	tsc1("solar")
	tsc1("martian")
	tsc2("lunar")
	tsc3("solar")
	tfor1()
	tfor2(1)
	tfor3(0)
	tfor4(2)
}
