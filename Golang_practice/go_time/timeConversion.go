package main

import (
	"fmt"
	"strings"
)

func timeConversion(s string) string {

    //if the time is midnight
    if strings.HasSuffix(s,"AM") && s[0:2] == "12"{
        //replace the hour time of 12 with 00
        s = strings.Replace(s,"12","00",1)
      //if it is after noon but not noon
    } else if strings.HasSuffix(s,"PM") && s[0:2] !="12"{
        //create a string called change.
        //change takes the 2 charcters in the hour position
        //and adds 10 to the 10s place and 2 to the ones place
        change := fmt.Sprintf("%c%c",s[0]+1,s[1]+2)
        //replaces the first 2 indexs with the previously made 'change' string
        s = strings.Replace(s,s[0:2],change,1)
    }
    //return the first 8 indexes of the newly made string, excluding the
    //suffix "AM" or "PM"
    return s[0:8]
}