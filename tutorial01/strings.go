package main

import (
	"fmt"
)

func main() {
	substring := "kitty cat meow all day long."

	fmt.Printf("\n\nPrinting out a range of a string.\n\n")
	fmt.Printf("start to point: %s\n", substring[0:9])
	fmt.Printf("start to point: %s\n", substring[:9])
	fmt.Printf("point to point: %s\n", substring[10:20])
	fmt.Printf("point to end: %s\n", substring[10:])


	//length of the string
	fmt.Printf("length: %d\n", len(substring))


	// iterate with the range function
	// for loop over the string
	for i, r := range substring {
		//index then character
		fmt.Printf("%d %c \n", i, r)
	}

	// for loop over the string, without the index value
	for _, r := range substring {
		//index then character
		//fmt.Printf("%c \n", r)
		// of all on one line
		fmt.Printf("%c ", r)
	}

}