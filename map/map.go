package main

import (
	"fmt"
)

func main() {
	// initialize the map
	/*	var lang map[string] string
		lang["en"] = "english"
		fmt.Println(lang)
	>>	panic: assignment to entry in nil map
	*/
	lang := map[string] string {"en": "english", "fr": "french", "de": "german"}
	fmt.Println("initialized a map lang: ", lang)
	fmt.Println("length of map lang: ", len(lang))

	// creating empty map
	fruits := make(map[string] string, 4)
	fmt.Println("created empty map fruits: ", fruits)
	fruits["citrus"] = "Orange"
	fmt.Println("created empty map fruits: ", fruits)
	
	//fetch value from map by key
	value, status := fruits["citrus"]
	fmt.Println("status: ", status, "value: ", value)

	// fetch value from map for un-defined key -> retunns status: false , value: nil
	value, status = fruits["locked"]
	fmt.Println("status: ", status, "value: ", value)

	// adding key-value pair
	fruits["berries"] = "cherry"
	fmt.Println("new fruits map is : ", fruits)

	// delete key-value pair
	delete(lang, "en")
	fmt.Println("lang after delete is : ", lang)

	// iterate over maps
	for key, val := range fruits {
		fmt.Println(key, "=>", val)
	}

	// truncate (empty the map)
	for key, _ := range fruits {
		delete(fruits, key)
	fmt.Println("truncated maps is: ", fruits)
	}
	// or
	lang = make(map[string] string, 4)
	fmt.Println("Truncated lang is: ", lang)

}