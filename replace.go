package main

//*---------------------*//
// REPLACEMENT FUNCTIONS //
//*---------------------*//

import (
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
	corrected = getMatrix(corrected)
	corrected = replaceFrac(corrected)
	corrected = replaceBrackets(corrected)
	corrected = replaceShape(corrected)
	corrected = replaceSymbol(corrected)
	corrected = replaceText(corrected)
	s += corrected
	return
}

// Replace fonts
// i.e add latex "mathbb" or "mathcal" prefixes to certain matches
func replaceFont(math string) (corrected string) {
	def, err := DefaultMathRegexp()
	check(err)
	corrected = math
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
	return
}

// Prepend with "\"
// this makes "sin" -> "\sin" for example
func replaceKeywords(math string) (corrected string) {
	def, err := DefaultMathRegexp()
	check(err)
	corrected = math

	//helper function to prefix with backslash where needed
	addBackslash := func(key string) {
		ref := def[key]
		matched := ref.MatchString(corrected)
		if matched {
			corrected = ref.ReplaceAllString(corrected, "\\$1")
		}
		return
	}

	addBackslash("FunctionRegexp")
	addBackslash("LetterRegexp")
	addBackslash("LogicRegexp")

	// replace 'inf' with '\infty'
	inf, err2 := regexp.Compile("inf")
	check(err2)
	matchinf := inf.MatchString(math)
	if matchinf {
		corrected = inf.ReplaceAllString(corrected, "\\infty")
	}
	return
}

// replaces "a/b" or "{a + b}/{c + d} with \frac{a}{b} or \frac{a + b}/{c + d}
// currently a WIP
// nested fractions are unavailable yet
func replaceFrac(math string) (corrected string) {
	def, err := DefaultMathRegexp()
	check(err)

	corrected = math
	frac := def["FracRegexp"]
	matched := frac.MatchString(math)
	if matched {
		corrected = frac.ReplaceAllString(math, "\\frac{$1}{$2}")
	}
	return
}

// Parenthesis, brakets, and braces
// this will change brakets in favour of size-adjusting ones:
// thus ( -> \left( and ] -> \right]
func replaceBrackets(math string) (corrected string) {

	corrected = math
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
	return
}

// Symbol transforms
// will change "=>" to "\implies" for example
func replaceSymbol(math string) (corrected string) {
	corrected = math

	// TODO: this should probably be stored in the syntac json file
	// thus changing syntax could be more modular and extensible
	repls := map[string]string{
		"<=>": "\\iff",
		"=>":  "\\implies",
		"->":  "\\to",
		"|->": "\\mapsto",
		">=":  "\\ge",
		"<=":  "\\le",
		"!=":  "\\neq",
		"~=":  "\\approx",
		"-=":  "\\equiv",
		"xx":  "\\times",
		"+-":  "\\pm",
		"...": "\\cdots",
		".":   "\\cdot",
	}
	for k, v := range repls {
		corrected = strings.ReplaceAll(corrected, k, v)
	}
	return
}

// this makes text syntax easier, using "text" instead of \text{text}
func replaceText(math string) (corrected string) {
	def, err := DefaultMathRegexp()
	check(err)
	corrected = math
	text := def["TextRegexp"]
	matched := text.MatchString(math)
	if matched {
		corrected = text.ReplaceAllString(math, "\\text{$1}")
	}
	return
}

// replaces "u^{_} with "\overline{u}" and other replacements
func replaceShape(math string) (corrected string) { // TODO: find a better name for this function
	def, err := DefaultMathRegexp()
	check(err)
	corrected = math
	shape := def["ShapeRegexp"]
	matched := shape.MatchString(math)

	repl := func(a, b string) {
		corrected = strings.ReplaceAll(corrected, a, b)
	}

	if matched {
		corrected = shape.ReplaceAllString(math, "\\($2){$1}")
		repl("(_)", "overline")
		repl("(->)", "overrightarrow")
		// the only reason this is here is in case I change the order of
		// the replace functions
		repl("(\\to)", "overrightarrow")
		repl("(^)", "hat")
		repl("(~)", "tilde")
		repl("(.)", "dot")
	}
	return
}

// DEPRECATED
/*
func replaceMatrix(math string) (corrected string) {
	def, err := DefaultMathRegexp()
	check(err)
	corrected = math
	matrix := def["MatrixRegexp"]
	matched := matrix.MatchString(math)
	if matched {
		corrected = matrix.ReplaceAllString(corrected, "\\begin{matrix}\n$1\n\\end{matrix}")
		corrected = strings.ReplaceAll(corrected, ";", "\\\\\n")
	}
	return
}
*/
