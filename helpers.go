package main

import (
	"fmt"
	"sort"
	"strings"
)

//*------------------------*//
// VARIOUS HELPER FUNCTIONS //
//*------------------------*//

// general function to panic on errors
func check(err error) {
	if err != nil {
		panic(err)
	}
}

//injects strings at given locations
// takes three arguments: string to modify, map of position/things to replace, map of position / things to replace with
func inject(s string, before, after map[int]string) (ret string) {
	// first check that the maps have the same locations
	var beforeLocs []int
	var afterLocs []int
	for i := range before {
		beforeLocs = append(beforeLocs, i)
	}
	for i := range after {
		afterLocs = append(afterLocs, i)
	}
	beforeLocs = sort.IntSlice(beforeLocs)
	afterLocs = sort.IntSlice(afterLocs)
	for i := 0; i < len(beforeLocs); i++ {
		if beforeLocs[i] != afterLocs[i] {
			fmt.Println("before: ", beforeLocs)
			fmt.Println("after: ", afterLocs)
			panic("start and end locations aren't the same")
		}
	}

	//mark locations to inject strings
	marker := ">*|*<"
	ret = s
	for _, i := range beforeLocs {
		ret = strings.Replace(ret, before[i], marker, 1)
	}
	//generate slice of replacements ordered
	var replacements []string
	for _, i := range afterLocs {
		replacements = append(replacements, after[i])
	}
	//make the final string
	matricesToReplace := strings.Count(ret, marker)
	for i := 0; i < matricesToReplace; i++ {
		ret = strings.Replace(ret, marker, replacements[i], 1)
	}
	return
}

// returns index of matching braket in string
func getMatchingBracket(text string, left, right rune, startIndex int) (index int) {
	counter := 0
	i := startIndex
	// check that the index isn't the last rune
	if startIndex >= len(text)-2 {
		panic("index out of range")
	}
	// check that the left rune matches the rune at the start index
	if rune(text[startIndex]) != left {
		fmt.Println(text[startIndex])
		panic("braket never started")
	} else {
		counter++
	}
	for counter > 0 {
		c := rune(text[i+1])
		i++
		if c == left {
			counter++
		} else if c == right {
			counter--
		}
	}
	return i
}
