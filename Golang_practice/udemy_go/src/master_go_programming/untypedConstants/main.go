package main

import "fmt"

func main() {
	const a float64 = 5.1	//typed constants
	const b = 6.7	//untyped constant

	const c float64 = a * b
	const str = "Hello" + "Go!"
	_, _ = c, str

	const d = 5 > 10
	fmt.Println(d)

	// const x int = 5
	// const y float64 = 2.2 * x
	//the above causes problems because of different value types

	const x = 5 //default type is int
	const y = 2.2 * x //default type is float64
	fmt.Printf("%T\n", y) //output is float64
	//this will work because of untyped constants

	//untyped consts do not yet have a delcared value
	//this allows for more freedom like in the examples above for constants.
	//not forced by strict rules that dont allow math with different type values.


	var i int = x //changes x to int
	var j float64 = x //var j float64(x)
	var p byte = x

	fmt.Println(i, j, p)


}