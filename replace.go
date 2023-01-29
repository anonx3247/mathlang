package main

//*---------------------*//
// REPLACEMENT FUNCTIONS //
//*---------------------*//

import (
	"fmt"
	"regexp"
	"strings"
)

type Direction bool
type String string

const Right Direction = true
const Left Direction = false

// Replace fonts
// Prepend with "\"
// Symbol transforms
// Parenthesis, brakets, and braces
// Advanced notation
func replace(math String) String {
	return math.replaceFont().
		prefixBackslash().
		replaceFrac().
		replaceParnethesis().
		replacePipe().
		replaceShape().
		replaceSymbol().
		replaceText().
		replaceBlock("matrix", "&").
		replaceBlock("cases", "@")
}

/*
***********************

UTILITY METHODS ON `String`

*************************
*/
func (math String) string() string {
	return string(math)
}

func (math String) regexpReplace(re *regexp.Regexp, repl string) (corrected String) {
	corrected = math
	matched := re.MatchString(math.string())
	if matched {
		corrected = String(re.ReplaceAllString(corrected.string(), repl))
	}
	return
}

/*
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
*/
func (math String) regexpDefReplace(key string, repl string) String {
	def := DefaultMathRegexp()
	re := def[key]
	return math.regexpReplace(re, repl)
}
func (math String) regexpCompileAndReplace(restr string, repl string) String {
	re, e := regexp.Compile(restr)
	check(e)
	return math.regexpReplace(re, repl)
}

func (math String) stringsReplace(search string, repl string) String {
	return String(strings.ReplaceAll(math.string(), search, repl))
}

/************************

Replacing Methods for `String`

**************************/

// Replace fonts
// i.e add latex "mathbb" or "mathcal" prefixes to certain matches
func (math String) replaceFont() String {
	return math.regexpDefReplace("MathbbRegexp", "\\mathbb{$1}").regexpDefReplace("MathcalRegexp", "\\mathcal{$1}")
}

/*
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
*/

// Prepend with "\"
// this makes "sin" -> "\sin" for example
func (math String) prefixBackslash() String {
	return math.regexpDefReplace("FunctionRegexp", "\\$1").
		regexpDefReplace("LogicRegexp", "\\$1").
		regexpDefReplace("LetterRegexp", "\\$1").
		stringsReplace("inf", "\\infty")
}

func (math String) replaceFrac() (corrected String) {
	// check for '/' in text
	hasFractions := func(txt String) (bool, int) {
		for i := 1; i < len(txt)-1; i++ {
			if txt[i] == '/' && txt[i-1] != ' ' && txt[i+1] != ' ' {
				return true, i
			}
		}
		return false, len(txt)
	}

	// get root '/'
	getRoot := func(loc int, txt String) int {
		foundParent := true
		for foundParent {
			foundParent = false
			// search for parent to the right
			for j := loc + 1; j < len(txt)-1; j++ {
				if txt[j:j+2] == "}/" {
					m := getMatchingBracket(txt, j, Left)
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
					m := getMatchingBracket(txt, j, Right)
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

	// get root's children
	getChildrenOf := func(root int, txt String) (children [2][2]int) {
		children[0][0] = root + 1
		children[1][1] = root - 1
		getType := func(loc int, s String) string {
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
			children[0][1] = getMatchingBracket(txt, root+1, Right)
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
			children[1][0] = getMatchingBracket(txt, root-1, Left)
		} else {
			panic(fmt.Sprintf("illegal child type:", leftChildType))
		}
		return
	}

	replaceFracChildren := func(r int, c [2][2]int, txt String) String {
		right := txt[c[0][0] : c[0][1]+1]
		left := txt[c[1][0] : c[1][1]+1]

		fixBracket := func(s String) String {
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
	corrected = math
	for check, loc := hasFractions(corrected); check; check, loc = hasFractions(corrected) {
		root := getRoot(loc, corrected)
		children := getChildrenOf(root, corrected)
		corrected = String(replaceFracChildren(root, children, corrected))
	}
	return
}

func getMatchingBracket(str String, loc int, direction Direction) (match int) {
	if loc == len(str)-1 {
		panic("Open bracket at end of string!")
	}

	depth := 1

	i := loc
	for depth != 0 && i < len(str) && i >= 0 {
		if direction == Right {
			i++
		} else if direction == Left {
			i--
		} else {
			panic(fmt.Sprintf("illegal direction:", direction))
		}
		if str[i] == '{' {
			if direction == Right {
				depth++
			} else {
				depth--
			}
		} else if str[i] == '}' {
			if direction == Right {
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

/*
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
*/
/*
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
*/
// Parenthesis, brakets, and braces
// this will change brakets in favour of size-adjusting ones:
// thus ( -> \left( and ] -> \right]
func (math String) replaceParnethesis() (corrected String) {
	/*
	   func replaceBrackets(math string) (corrected string) {
	*/

	corrected = math
	replBr := func(lStr, rStr string) {
		corrected = corrected.regexpCompileAndReplace("\\"+lStr, "\\left"+lStr).
			regexpCompileAndReplace("\\"+rStr, "\\right"+rStr)
	}

	replBr("(", ")")
	replBr("[", "]")
	replBr("\\{", "\\}")
	return
}

func (math String) replacePipe() (corrected String) {
	// first do it for doubles '||'
	return math.regexpCompileAndReplace("(\\s|^)\\|\\|(\\S)", " \\left|\\left|$2").
		regexpCompileAndReplace("(\\S)\\|\\|(\\s|$)", "$1\\right|\\right| ").
		// then for singles '|'
		regexpCompileAndReplace("(\\s|^)\\|(\\S)", " \\left|$2").
		regexpCompileAndReplace("(\\S)\\|(\\s|$)", "$1\\right| ").
		// a final correction is needed (this is a bit of a hack really)
		// because the double will match '...text|| ' to '...text\right|\right| '
		//and then the single will match '\right| ' to '\right\right| '
		//thus we will remove all instances of '\right\right| ' and replace them with '\right| '
		stringsReplace("\\right\\right| ", "\\right| ")
}

// Symbol transforms
// will change "=>" to "\implies" for example
func (math String) replaceSymbol() (s String) {
	s = math
	symbols := []string{
		"<=>",
		"=>",
		"|->",
		"->",
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
		"\\mapsto",
		"\\to",
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
		s = s.stringsReplace(symbols[i], repls[i])
	}
	return
}

/*
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
*/
func (math String) replaceText() String {
	return math.regexpDefReplace("TextRegexp", "\\text{$1}")
}
func (math String) replaceShape() String {
	matched := DefaultMathRegexp()["ShapeRegexp"].MatchString(math.string())
	/*
		// this makes text syntax easier, using "text" instead of \text{text}
		func replaceText(math string) (corrected string) {
			def, err := DefaultMathRegexp()
			check(err)
			corrected = math
			text := def["TextRegexp"]
			matched := text.MatchString(math)
	*/
	if matched {
		return math.regexpDefReplace("ShapeRegexp", "\\($2){$1}").
			stringsReplace("(_)", "overline").
			stringsReplace("(->)", "overrightarrow").
			stringsReplace("(\\to)", "overrightarrow").
			stringsReplace("(^)", "hat").
			stringsReplace("(~)", "tilde").
			stringsReplace("(.)", "dot")
	} else {
		return math
	}
}
func (math String) replaceBlock(title String, prefix String) (s String) {
	s = math
	found := false
	starts := make([]int, 0)
	for i := 0; i < len(math)-2; i++ {
		if math[i:i+2] == prefix+"{" {
			found = true
			starts = append(starts, i)
		}
	}
	if !found {
		return
	}
	ends := make([]int, 0)

	for _, i := range starts {
		ends = append(ends, getMatchingBracket(math, i+1, Right))
	}

	repl := func(start int, end int) String {
		f := String("\\begin{" + title + "} ")
		m := math[start+2:end].stringsReplace(",", " & ").stringsReplace(";", "\\\\ ")
		f += m
		f += " \\end{" + title + "}"
		return f
	}

	s = math[0:starts[0]]
	for i := 0; i < len(starts); i++ {
		s += repl(starts[i], ends[i])
		if i < len(starts)-1 {
			s += math[ends[i]+1 : starts[i+1]]
		}
	}
	if len(math) > ends[len(ends)-1]+1 {
		s += math[ends[len(ends)-1]+1:]
	}
	return
}

/*
=======
	return
}
*/
/*
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
*/
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
