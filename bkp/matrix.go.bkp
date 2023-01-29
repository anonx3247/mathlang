package main

//*----------------------------*//
// MATRIX SYNTAX IMPLEMENTATION //
//*----------------------------*//

import (
	"regexp"
	"strings"
)

// returns Latex matrix syntax from mathlang syntax
func getMatrix(text string) (matrix string) {

	// check if there is a matrix
	hasMatrix := false
	matrix = ""
	if strings.Contains(text, ";") {
		hasMatrix = true
	}

	// find matrix within brackets
	bracket, err := regexp.Compile("{")
	check(err)
	matched := bracket.MatchString(text)
	if !hasMatrix {
		return
	} else {
		if !matched {
			return
		} else {
			//get locations of brackets
			brackets := bracket.FindAllStringIndex(text, 100)
			var locations []int
			for i := range brackets {
				locations = append(locations, brackets[i][0])
			}
			// get substrings
			var subStrings = make(map[int]string)
			for _, loc := range locations {
				match := getMatchingBracket(text, '{', '}', loc)
				substr := text[loc : match+1]
				subStrings[loc] = substr
			}
			// check if any substring is included in another
			minimalSubStrings := getMinimalMatrixSubstrings(subStrings)
			// transform substrings into matrices
			var matrices = make(map[int]string)
			for loc, sub := range minimalSubStrings {
				//endLoc := loc + len(sub) - 1
				matrices[loc] = strToMatrix(sub)
			}
			matrix = inject(text, minimalSubStrings, matrices)
		}
		return
	}
}

// transform substrings into matrices
func strToMatrix(s string) (m string) {
	trimmed := s[1 : len(s)-1]           //remove first and last char i.e. the brackets
	lines := strings.Split(trimmed, ";") //keeps track of lines
	height := len(lines)
	width := 1
	var matrixItem = make([][]string, height) // stored as [row][column]

	// generate table of matrix values
	for row, line := range lines {
		if strings.Contains(line, ",") {
			elements := strings.Split(line, ",")
			width = len(elements)
			matrixItem[row] = make([]string, width)
			for column, element := range elements {
				matrixItem[row][column] = element
			}
		} else {
			matrixItem[row] = make([]string, width)
			matrixItem[row][0] = line
		}
	}

	//generate output string from table values
	m = "\\begin{matrix}"
	for row := range matrixItem {
		for column := range matrixItem[row] {
			m += matrixItem[row][column]
			if column != width-1 {
				m += " & "
			}
		}
		if row != height-1 {
			m += "\\\\"
		}
	}
	m += "\\end{matrix}"
	return
}

// check if any substring is included in another
func getMinimalMatrixSubstrings(substrings map[int]string) (minimal map[int]string) {
	minimal = make(map[int]string)
	isMinimal := func(s string) bool {
		for _, sub := range substrings {
			if strings.Contains(s, sub) && strings.Contains(sub, ";") && s != sub {
				return false
			}
		}
		if strings.Contains(s, ";") {
			return true
		} else {
			return false
		}
	}
	// only get the smallest substrings with ";" in them
	for loc, substrings := range substrings {
		if isMinimal(substrings) {
			minimal[loc] = substrings
		}
	}
	return
}
