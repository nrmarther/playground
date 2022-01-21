package main

import (
	"fmt"
	"math"
)

func main() {
	var x uint8 = 255 //max value of uint8
	x++ //will cause overflow

	fmt.Println(x) //value of x will reset to lowest value, 0
	//error not detected at compiletime and instead at runtime.
	//go operates assuming that no overflow will occur.

	//a := int8(255+1) //go catches this error at compiletime because of declaration

	var b int8 = 127

	fmt.Printf("%d\n", b+1) //no compiletime error
	//at runtime if it overflows, goes to minimum value for its type, -128

	b = -128
	b-- //underflow of variable type. will go to max value

	fmt.Println(b) //prints 127

	f := float32(math.MaxFloat32) //shows max of float32
	fmt.Println(f) 
	f *= 1.2 // f = +Inf or +Infinite
	fmt.Println(f) //float numbers overflow to infinite!!
	
	//constants cannot overflow because it will be caught at compiletime
}