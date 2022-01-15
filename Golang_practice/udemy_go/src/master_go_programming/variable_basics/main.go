package main

import "fmt"

func main() {
	var age int = 30
	fmt.Println("Age: ", age)

	//inferred declaration
	var name = "Dan"
	// fmt.Println("Your name is: ", name)
	_ = name //ignores error that would be had if name had not been used

	s := "Learning glang!"
	fmt.Println(s)

	name = "Andrei" //changes name of existing variable
	name1 := "John"
	_ = name1

	car,cost := "Audi", 50000
	fmt.Println(car,cost)


//can use the short hand/multi-line declaration
//as long as at least one variable is new
	car, year := "CMW", 2021
	_ = year

	var opened = false
	opened, file := true, "a.txt"

	_, _ = opened, file

	//multiple declarations with good readability
	var (
		salary float64
		firstName string
		gender bool
	)

	fmt.Println(salary, firstName, gender)

	//multiple declarations
	var a, b, c int
	//multiple assignment
	a, b, c = 5, 6, 7
    _, _, _ = a, b, c

	a, b, c = c, b, a //swapping variables

	fmt.Println(a, b, c)

	sum := 5 + 2.3
	fmt.Println(sum)


}