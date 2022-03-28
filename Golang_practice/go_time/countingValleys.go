package main

func countingValleys(steps int32, path string) int32 {
    //declaring global variables
    var currPos, valleys int32;
    
    //start tracking steps up and down
    for i:=int32(0);i<steps;i++ {
        if path[i] == 'D' {
            currPos--
        } else {
            currPos++
        }
        if path[i] == 'U' && currPos == 0 {
            valleys++
        }
    }
    return valleys;
}