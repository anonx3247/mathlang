package main

//*-----------------------*//
// JSON FILE REGEXP PARSER //
//*-----------------------*//

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

// fetches the regexp strings from "syntax_regexp.json" file
func DefaultMathRegexp() (def map[string]*regexp.Regexp) {

	bkpJSONFile := "syntax_regexp.json"
	JSONFile := "/usr/local/share/mathlang/syntax_regexp.json"
	file := JSONFile
	// This will serve as a tool to get access to all the regexps stored in
	// the "syntax_regexp.json" file
	/*
		func DefaultMathRegexp() (def map[string]*regexp.Regexp, err error) {
	*/
	//JSONFile := os.ExpandEnv("$HOME/.config/mathlang/syntax_regexp.json")

	//check if file is readable
	test, readError := os.Open(JSONFile)
	if readError != nil {
		bkp, bkpError := os.Open(bkpJSONFile)
		if bkpError != nil {
			panic("'/usr/local/share/mathlang/syntax_regexp.json' file missing!")
		} else {
			file = bkpJSONFile
		}
		bkp.Close()
	}
	test.Close()

	// helper function to parse json file and then compile its regexp
	read := func(key string) (re *regexp.Regexp) {
		re, _ = regexp.Compile(JSONRead(file, key))
		return
	}

	// definition of the returned dictionary
	def = map[string]*regexp.Regexp{
		"FunctionRegexp": read("function"),
		"LetterRegexp":   read("letters"),
		"LogicRegexp":    read("logic"),
		"ShapeRegexp":    read("shape"),
		"MathbbRegexp":   read("mathbb"),
		"MathcalRegexp":  read("mathcal"),
		"FracRegexp":     read("frac"),
		"TextRegexp":     read("text"),
		"MatrixRegexp":   read("matrix"),
	}
	return
}

// parses json file with a given key, and returns
// the value stored for that key
func JSONRead(file, key string) (value string) {
	data, err := ioutil.ReadFile(file)

	// check for errors
	if err != nil {
		panic(err)
	}

	// create a map of keys to values from json file
	var jsonMap map[string]string
	err2 := json.Unmarshal(data, &jsonMap)

	if err2 != nil {
		panic(err2)
	}

	// check if the requested key is in the file
	val, ok := jsonMap[key]

	if !ok {
		fmt.Println("Key:", key, "unavailable")
		panic(ok)
	}

	value = val
	return

}
