package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestAllRegexp(t *testing.T) {
	lines := []string{}
	math := Math{lines}
	file, err := os.Open("test.math")
	check(err)
	math.SetupWithFile(file)
	fmt.Println("TESTING REGEXPS...")
	def, err := DefaultNoteRegexp()
	if err != nil {
		fmt.Println("Unable to access regexp file.")
	}

	for key, re := range def {
		fmt.Println("Matches for: [", key, "] '", re, "'")
		if strings.TrimSpace(re.String()) != "" {
			for i, line := range math.lines {
				matched := re.MatchString(line)
				if matched {
					fmt.Println(i+1, line)
				}
			}
		}

	}
	check(err)
}
