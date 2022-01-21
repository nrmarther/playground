package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := string(99)
	fmt.Println(s) //s will equal "c". this is because 99 is ASCII code for c.

	// s1 := string(44.2) ---cannot convert untyped float constant to string
	var myStr = fmt.Sprintf("%f", 44.2)
	//first argument is what type is converting from, next is the value
	fmt.Println(myStr)

	var myStr1 = fmt.Sprintf("%d", 34234)
	//%d for integers, %f for floats
	fmt.Println(myStr1)

	//if we use string() function it will print the unicode value at that code.
	//in this case it will print "薺"
	fmt.Println(string(34234))


	//STRINGS TO NUMBERS
	s1 := "3.123" //type is string
	fmt.Printf("%T\n", s1)

	var f1, err = strconv.ParseFloat(s1, 64) //first argument is string or var of string
											 //second value is size or precision. in this case float64
	_ = err

	fmt.Println("f1 is", f1)
	fmt.Printf("f1 type is %T\n", f1)
	fmt.Println("f1 * 3.4 is", f1 * 3.4)

	i, err := strconv.Atoi("-50") //Atoi is String or ASCII to int
	_ = err
	s2 := strconv.Itoa(20)

	fmt.Printf("i type is %T, i value is %v\n", i, i)
	fmt.Printf("s2 is type %T, s2 value is %q\n", s2, s2)
}