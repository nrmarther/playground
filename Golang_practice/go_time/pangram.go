package main

import (
	"strings"
)
func pangrams(s string) string {
    // Write your code here
    alphabet := "abcdefghijklmnopqrstuvwxyz"
    s = strings.ToLower(s)
    cMap := make(map[rune]bool)
    for _, c := range s {
        if strings.ContainsRune(alphabet, c) && !cMap[c] {
            cMap[c] = true
        }
    }
    
    if len(cMap) == len(alphabet) {
        return "pangram"
    }
    return "not pangram"
}