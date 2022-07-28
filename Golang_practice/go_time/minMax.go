package main

import (
	"fmt"
)

func miniMaxSum(arr []int32) {
    // variable declarations
    //min max are going to be calculated values without the largest or
    //smallest numbers. total is all 5 numbers from array combined
    var min, max, total int64;
    var small int32 = arr[0];
    var large int32 = arr[0];
    
    //loop through full array
    for i:=0;i<len(arr);i++ {
        //finding the smallest number in array
        if arr[i] < small {
            small = arr[i]
        //finding largeest number in array
        } else if arr[i] > large {
            large = arr[i]
        }
        //calculating total
        total += int64(arr[i])
    }
    //calculate the min and max
    min = total - int64(large)
    max = total - int64(small)
    
    //print results
    fmt.Printf("%d %d", min, max)
}