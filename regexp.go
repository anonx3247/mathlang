package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func DefaultMathRegexp() (def map[string]*regexp.Regexp, err error) {

	//JSONFile := "/home/neosapien/development/go/mathlang/syntax_regexp.json"
	JSONFile := "/usr/local/share/mathlang/syntax_regexp.json"

	//check if file is readable
	test, readError := os.Open(JSONFile)
	if readError != nil {
		err = readError
		return
	}
	check(readError)
	test.Close()

	read := func(key string) (re *regexp.Regexp) {
		re, _ = regexp.Compile(JSONRead(JSONFile, key))
		return
	}

	// EXTRA
	def = map[string]*regexp.Regexp{
		"FunctionRegexp": read("function"),
		"LetterRegexp":   read("letters"),
		"SymbolRegexp":   read("symbol"),
		"LogicRegexp":    read("logic"),
		"ShapeRegexp":    read("shape"),
		"MathbbRegexp":   read("mathbb"),
		"MathcalRegexp":  read("mathcal"),
		"FracRegexp":     read("frac"),
		"TextRegexp":     read("text"),
		"MatrixRegexp":   read("matrix"),
	}

	/*
		def = map[string]*regexp.Regexp{
			"FunctionRegexp": FunctionRegexp,
			"SymbolRegexp":   SymbolRegexp,
			"LetterRegexp":   LetterRegexp,
			"LogicRegexp":    LogicRegexp,
			"ShapeRegexp":    ShapeRegexp,
			"MathbbRegexp":   MathbbRegexp,
			"MathcalRegexp":  MathcalRegexp,
			"FracRegexp":     FracRegexp,
			"TextRegexp":     TextRegexp,
			"MatrixRegexp":   MatrixRegexp,
		}
	*/
	return
}

func JSONRead(file, key string) (value string) {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	var jsonMap map[string]string

	err2 := json.Unmarshal(data, &jsonMap)

	if err2 != nil {
		panic(err2)
	}

	val, ok := jsonMap[key]

	if !ok {
		fmt.Println("Key:", key, "unavailable")
		panic(ok)
	}

	value = val
	return

}
