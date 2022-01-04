package main

import (
	"testing"
)
func TestSearch( t *testing.T) {

	var testArr = []int{-2, 3, 4, 5, 6, 7, 100} //array to test from
	var target1 = 7 //number to search for
	var target2 = 0 //number to search for
	var index = 5 //index of target
	properResult := Search(testArr, target1) //should return testArr[5] as index for 7
	improperResult := Search(testArr, target2)
	
	//test a normal use case
	if properResult != 5 {
		t.Errorf("ERROR: should be %v, but got %v", index, properResult)
		t.Fail()
	} else {
		t.Logf("PASS | expected %v and got %v", index, properResult)
	}
	//test for number not found
	if improperResult != -1 {
		t.Errorf("ERROR: should return -1 but got %v", improperResult)
		t.Fail()
	} else {
		t.Logf("PASS | Number not found in array, function returned -1")
	}

}
