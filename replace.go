package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Replace fonts
// Prepend with "\"
// Symbol transforms
// Parenthesis, brakets, and braces
// Advanced notation
func replace(math string) (s string) {
	corrected := math
	corrected = replaceFont(corrected)
	corrected = replaceKeywords(corrected)
	corrected = replaceFrac(corrected)
	corrected = replaceParnethesis(corrected)
	corrected = replaceShape(corrected)
	corrected = replaceSymbol(corrected)
	corrected = replaceText(corrected)
	corrected = replaceMatrix(corrected)
	s += corrected
	return
}

// Replace fonts
// i.e add latex "mathbb" or "mathcal" prefixes to certain matches
func replaceFont(math string) (s string) {
	def, err := DefaultMathRegexp()
	check(err)
	corrected := math
	mathbb := def["MathbbRegexp"]
	mathcal := def["MathcalRegexp"]
	matched := mathbb.MatchString(math)
	matched2 := mathcal.MatchString(math)
	if matched {
		corrected = mathbb.ReplaceAllString(math, "\\mathbb{$1}")
	}
	if matched2 {
		corrected = mathcal.ReplaceAllString(corrected, "\\mathcal{$1}")
	}
	s += fmt.Sprintf(corrected)
	return
}

// Prepend with "\"
// this makes "sin" -> "\sin" for example
func replaceKeywords(math string) (s string) {
	def, err := DefaultMathRegexp()
	check(err)

	corrected := math
	addBackslash := func(key string) {
		ref := def[key]
		matched := ref.MatchString(corrected)
		if matched {
			corrected = ref.ReplaceAllString(corrected, "\\$1")
		}
	}

	addBackslash("FunctionRegexp")
	addBackslash("LetterRegexp")
	addBackslash("LogicRegexp")

	inf, err2 := regexp.Compile("inf")
	check(err2)
	matchinf := inf.MatchString(math)
	if matchinf {
		corrected = inf.ReplaceAllString(corrected, "\\infty")
	}

	s += fmt.Sprintf(corrected)
	return
}

func replaceFrac(math string) (s string) {
	// check for '/' in text
	hasFractions := func(txt string) (bool, int) {
		for i := 0; i < len(txt); i++ {
			if txt[i] == '/' {
				return true, i
			}
		}
		return false, len(txt)
	}

	// get root '/'
	getRoot := func(loc int, txt string) int {
		foundParent := true
		for foundParent {
			foundParent = false
			// search for parent to the right
			for j := loc + 1; j < len(txt)-1; j++ {
				if txt[j:j+2] == "}/" {
					m := getMatchingBracket(txt, j, "left")
					// if the found '/' is a child of 'j'
					if m < loc && loc < j {
						foundParent = true
						loc = j + 1
					}
				}
			}

			// search for parent to the left
			for j := loc - 1; j > 0; j-- {
				if txt[j-1:j] == "/{" {
					m := getMatchingBracket(txt, j, "right")
					// if the found '/' is a child of 'j'
					if j < loc && loc < m {
						foundParent = true
						loc = j - 1
					}
				}
			}
		}

		return loc
	}

	// get root children
	getChildrenOf := func(root int, txt string) (children [2][2]int) {
		children[0][0] = root + 1
		children[1][1] = root - 1
		getType := func(loc int, s string) string {
			t := ""
			if s[loc] == '{' || s[loc] == '}' {
				t = "bracket"
			} else {
				t = "nobracket"
			}
			return t
		}

		rightChildType := getType(root+1, txt)
		leftChildType := getType(root-1, txt)

		if rightChildType == "nobracket" {
			i := root
			for i < len(txt) && txt[i] != ' ' && txt[i] != '}' {
				i++
			}
			children[0][1] = i - 1
		} else if rightChildType == "bracket" {
			children[0][1] = getMatchingBracket(txt, root+1, "right")
		} else {
			panic(fmt.Sprintf("illegal child type: ", rightChildType))
		}

		if children[0][1] == len(txt) {
			children[0][1] = len(txt) - 1
		}

		if leftChildType == "nobracket" {
			i := root - 1
			for i >= 0 && txt[i] != ' ' && txt[i] != '{' {
				i--
			}
			children[1][0] = i + 1
		} else if leftChildType == "bracket" {
			children[1][0] = getMatchingBracket(txt, root-1, "left")
		} else {
			panic(fmt.Sprintf("illegal child type:", leftChildType))
		}
		return
	}

	replaceFracChildren := func(r int, c [2][2]int, txt string) string {
		right := txt[c[0][0] : c[0][1]+1]
		left := txt[c[1][0] : c[1][1]+1]

		fixBracket := func(s string) string {
			if s[0] == '{' && s[len(s)-1] == '}' {
				return s
			} else {
				return "{" + s + "}"
			}
		}
		s := txt[0:c[1][0]]
		s += "\\frac"
		s += fixBracket(left)
		s += fixBracket(right)
		s += txt[c[0][1]+1:]
		return s
	}
	corrected := math
	for check, loc := hasFractions(corrected); check == true; check, loc = hasFractions(corrected) {
		root := getRoot(loc, corrected)
		children := getChildrenOf(root, corrected)
		corrected = replaceFracChildren(root, children, corrected)
	}
	s = corrected
	return
}

/*func replaceFrac(math string) (s string) {
	def, err := DefaultMathRegexp()
	check(err)

	corrected := math
	frac := def["FracRegexp"]
	matched := frac.MatchString(math)
	if matched {
		corrected = frac.ReplaceAllString(math, "\\frac{$1}{$2}")
	}

	s += fmt.Sprintf(corrected)
	return
}*/

func getMatchingBracket(str string, loc int, direction string) (match int) {
	if loc == len(str)-1 {
		panic("Open bracket at end of string!")
	}

	depth := 1

	i := loc
	for depth != 0 && i < len(str) && i >= 0 {
		if direction == "right" {
			i++
		} else if direction == "left" {
			i--
		} else {
			panic(fmt.Sprintf("illegal direction:", direction))
		}
		if str[i] == '{' {
			if direction == "right" {
				depth++
			} else {
				depth--
			}
		} else if str[i] == '}' {
			if direction == "right" {
				depth--
			} else {
				depth++
			}
		}
	}
	if i == len(str) {
		panic("out of bounds search for bracket!")
	}
	match = i
	return
}

// Parenthesis, brakets, and braces
// this will change brakets in favour of size-adjusting ones:
// thus ( -> \left( and ] -> \right]
func replaceParnethesis(math string) (s string) {

	corrected := math
	replBr := func(lStr, rStr string) {
		left, err := regexp.Compile("\\" + lStr)
		right, err2 := regexp.Compile("\\" + rStr)
		check(err)
		check(err2)
		matchedl := left.MatchString(math)
		matchedr := right.MatchString(math)
		if matchedl {
			corrected = left.ReplaceAllString(corrected, "\\left"+lStr)
		}
		if matchedr {
			corrected = right.ReplaceAllString(corrected, "\\right"+rStr)
		}
	}

	replBr("(", ")")
	replBr("[", "]")
	s += fmt.Sprintf(corrected)
	return
}

// Symbol transforms
// will change "=>" to "\implies" for example
func replaceSymbol(math string) (s string) {
	s = math
	symbols := []string{
		"<=>",
		"=>",
		"->",
		"|->",
		">=",
		"<=",
		"!=",
		"~=",
		"-=",
		"xx",
		"+-",
		"...",
		".",
	}

	repls := []string{
		"\\iff",
		"\\implies",
		"\\to",
		"\\mapsto",
		"\\ge",
		"\\le",
		"\\neq",
		"\\approx",
		"\\equiv",
		"\\times",
		"\\pm",
		"\\cdots",
		"\\cdot",
	}

	for i := range repls {
		s = strings.ReplaceAll(s, symbols[i], repls[i])
	}

	return
}

func replaceText(math string) (s string) {
	def, err := DefaultMathRegexp()
	check(err)
	corrected := math
	text := def["TextRegexp"]
	matched := text.MatchString(math)
	if matched {
		corrected = text.ReplaceAllString(math, "\\text{$1}")
	}
	s += fmt.Sprintf(corrected)
	return
}

func replaceShape(math string) (s string) {
	def, err := DefaultMathRegexp()
	check(err)
	corrected := math
	shape := def["ShapeRegexp"]
	matched := shape.MatchString(math)

	repl := func(a, b string) {
		corrected = strings.ReplaceAll(corrected, a, b)
	}

	if matched {
		corrected = shape.ReplaceAllString(math, "\\($2){$1}")
		repl("(_)", "overline")
		repl("(->)", "overrightarrow")
		repl("(\\to)", "overrightarrow")
		repl("(^)", "hat")
		repl("(~)", "tilde")
		repl("(.)", "dot")
	}
	s += corrected
	return
}

func replaceMatrix(math string) (s string) {
	s = math
	found := false
	starts := make([]int, 0)
	for i := 0; i < len(math)-2; i++ {
		if math[i:i+2] == "&{" {
			found = true
			starts = append(starts, i)
		}
	}
	if !found {
		return
	}
	ends := make([]int, 0)

	for _, i := range starts {
		ends = append(ends, getMatchingBracket(math, i+1, "right"))
	}

	repl := func(start int, end int) string {
		f := "\\begin{pmatrix} "
		m := math[start+2 : end]

		m = strings.ReplaceAll(m, ",", " & ")
		m = strings.ReplaceAll(m, ";", " \\\\ ")

		f += m

		f += " \\end{pmatrix}"
		return f
	}

	s = math[0:starts[0]]
	for i := 0; i < len(starts); i++ {
		s += repl(starts[i], ends[i])
	}
	if len(math) > ends[len(ends)-1]+1 {
		s += math[ends[len(ends)-1]+1:]
	}
	return
}

/*func replaceMatrix(math string) (s string) {
	def, err := DefaultMathRegexp()
	check(err)
	corrected := math
	matrix := def["MatrixRegexp"]
	matched := matrix.MatchString(math)
	if matched {
		corrected = matrix.ReplaceAllString(corrected, "\\begin{matrix}\n$1\n\\end{matrix}")
		corrected = strings.ReplaceAll(corrected, ";", "\\\\\n")
	}
	s = corrected
	return
}*/
