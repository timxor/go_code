package main

import (
	"fmt"
)

//struct template
type info struct {
	result string
}

//declare global constants
const (
	globalDoggy = "global doggy"
	globalInteger = 100
	globalIncrementer = iota
)

func main() {

	//declare a variable
	var message string = "message one"

	//declare a variable, short way
	//doggy := "doggy"


	//update int value
	//globalInteger += 1


	//64 bit accurate float
	var pi float64 = 3.14
	fmt.Printf("Value of pi: %f\n",pi)

	// only print two decimal places
	fmt.Printf("Value of pi: %.2f\n",pi)

	//32 bit, less accurate float
	var pi32bit float32 = 3.14

	//type inference
	pi32bitexample2 := float32(3.14)


	// int type
	nine := 9

	// 8 byte int
	five := uint8(5)


	// boolean value
	isTrue := true
	fmt.Printf("Value of isTrue: %t\n",isTrue)

	//byte value
	byteValue := byte(65)
	//print in hex
	fmt.Printf("Value of byteValue: %x\n",byteValue)













	fmt.Println("Go example: %+v\n ",message)
}


func example() {
	// var decl and inits
	//var ex1 string = "meow"
	fmt.Println("Go")
}





