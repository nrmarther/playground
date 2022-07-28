package main

import (
	"math"
	"fmt"
)
func plusMinus(arr []int32) {
    // Write your code here
    var pos, neg, zero float64
    for i:=0;i<len(arr);i++ {
        if arr[i] > 0 {
            pos++
        } else if arr[i] < 0 {
            neg++
        } else {
            zero++
        } 
    }
    //create the ratio as a float64 and round to 6 decimal places
    posRatio := pos/float64(len(arr))
    posRatio = math.Round(posRatio * 100000)/100000
    
    negRatio := neg/float64(len(arr))
    negRatio = math.Round(negRatio * 100000)/100000
    
    zRatio := zero/float64(len(arr))
    zRatio = math.Round(zRatio * 100000)/100000
    fmt.Printf("%.6g \n%.6g \n%.6g", posRatio, negRatio, zRatio)
}