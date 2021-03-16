package main

import "fmt"

//types are inferred
//inside functions, := can be used to assign value to variables **NO VAR NEEDED**

/*var name, lastname string = "vinay", "venkatesh"
var age int = 14
var tf bool = true*/

var name, lastname, age, tf = "vinay", "venkatesh", 14, true

func main() {
	fmt.Printf("My name is %s %s and I am %d years old.\n", name, lastname, age)
	k := 3
	j := 6
	fmt.Print(k + j)
}