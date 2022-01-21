package main

import "fmt"

type km float64
type mile float64

func main() {
	type age int //source type is int, defined type is age
	type oldAge age //int is its underlying or source type
	type veryOldAge age //int is still underlying type

	type speed uint //uint is its underlying type

	var s1 speed = 10 //type speed
	var s2 speed = 20 //type speed
	fmt.Printf("type of s1 is %T and type of s2 is %T\n", s1, s2)

	fmt.Println(s2 - s1)
	var x uint
	//x = s1 will not work because x and s1 are different types. requires conversion
	x = uint(s1) //correct way
	_ = x

	var s3 speed = speed(x)
	_= s3

	var parisToLondon km = 465
	var distanceInMiles mile
	distanceInMiles = mile(parisToLondon) / 0.621
	fmt.Printf("Type of distanceToMiles is %T\n", distanceInMiles)
	fmt.Printf("The distance from paris to london is %f\n", distanceInMiles)
}