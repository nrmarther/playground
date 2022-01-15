package main

import "fmt"

func main() {
	const days int = 7

	var i int
	fmt.Println(i)

	const pi float64 = 3.14
	const secondsInHour = 3600

	duration := 325 //in hours
	fmt.Printf("Duration in seconds: %v\n", duration*secondsInHour)

	x, y := 5, 1

	fmt.Println(x / y)

	const a = 5
	const b = 0

	// fmt.Println(a / b)

	const n, m int = 4, 5
	const n1, m1 = 6, 7

	const (
		min1 = -500
		min2
		min3
	)
	//grouped constants will inherit the value from the constant above it
	//that has been declared. min2 and min3 will both equal -500

	fmt.Println(min1, min2, min3)

	//Constants rules:
	//1. you cannot change a constant
	const temp int = 100
	// temp = 50
	_ = temp

	//2. Youy cannot initiate a constant at runtime
	// const power = math.Pow(2, 3) <- this would require initiation at runtime

	//3. You cannot use a variable to initialize a constant
	// t := 5
	// const tc = t

	//4.
	const l1 = len("hello")
	fmt.Println(l1)
}