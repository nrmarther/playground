package main

func breakingRecords(scores []int32) [2]int32 {
    // declare variables for high and low scores and
	//tally times records are broken with min and max
    var low int32 = scores[0]
    var high int32 = scores[0]
    var min, max int32

	//loop through entire array
    for i:=0;i<len(scores);i++ {
		//if the index breaks the high score, replace
		//the high score and add to tally
        if scores[i] > high {
            high = scores[i]
            max++
		//if above condition not met
		//check if index breaks low score, replace
		//low score and add to tally
        } else if scores[i] < low {
            low = scores[i]
            min++
        }
    }
	//put the high and low scores into array called records
    var records = [2]int32{max,min}
	//record records.... this is self explanatory
    return records
}
