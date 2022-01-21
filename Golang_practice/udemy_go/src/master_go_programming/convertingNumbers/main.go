package main

import "fmt"

func main() {
	var x = 3 //type int
	var y = 3.1 //type float

	//x = x*y //will give error for invalid operation for mismatched typed
	x = x * int(y)
	
	x = int(float64(x) * y)

	y = float64(x) * y
	

	var a = 5
	fmt.Printf("%T\n", a)

	var b int64 = 2
	fmt.Printf("%T\n", b)

	//a = b ---cannot use b(type int64) for type int in assignment
	a = int(b)

	_, _ = y, a //getting rid of compiltime errors
	
}