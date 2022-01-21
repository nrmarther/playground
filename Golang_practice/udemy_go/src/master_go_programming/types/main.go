package main

import "fmt"

func main() {
	//MAIN TYPE
	var i1 int8 = -128 //min size of 8 bit integer. if -129 is used it with give error
	fmt.Printf("%T\n", i1)

	var i2 uint16 = 65530 //max value of 16 bit int.
	fmt.Printf("%T\n", i2)

	//FLOAT TYPE
	// leading or trailing zeros before or after a decimal can be omitted
	var f1, f2, f3 float64 = 1.1, -.2, 5
	fmt.Printf("%T %T %T\n", f1, f2, f3)

	//RUNE TYPE
	var r rune = 'f'
	fmt.Printf("%T\n", r)
	fmt.Println(r)
	fmt.Printf("%x\n", r)

	//BOOL TYPE
	var b bool = true
	fmt.Printf("%T\n", b)
	fmt.Println(b)

	//STRING TYPE
	var s string = "Hello Go!"
	fmt.Printf("%T\n", s)
	fmt.Println(s)

	//If i were to want to, i could omit the type on the lefthand side because
	// the compiler would infer the type from the righthand side.

	//ARRAY TYPE
	var numbers = [4]int{4, 5, -9, 100}
	fmt.Printf("%T\n", numbers)

	// SLICE TYPE -- like an array but it can shrink or grow
	var cities = []string{"London", "Tokyo", "New York"}
	fmt.Printf("%T\n", cities)
	
	// MAP TYPE -- contains key-value pairs
	balances := map[string]float64{
		"USD": 2332.2,
		"EUR": 511.11,
	}
	fmt.Printf("%T\n", balances)
	fmt.Println(balances) //if no key is given, all key-value pairs are given

	//STRUCT TYPE
	type Person struct{
		name string
		age int
	}
	var you Person
	fmt.Printf("%T\n", you)

	you.name = "Nathan"
	you.age = 21
	fmt.Println(you)

	//POINTER TYPE
	var x int = 2
	ptrx := &x
	fmt.Printf("ptrx is a %T with a value of %v\n", ptrx, ptrx)

	//FUNCTION TYPE
	fmt.Printf("%T\n", f)
}

func f(){}