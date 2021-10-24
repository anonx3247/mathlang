package main

//*----------------------*//
// MAIN PROGRAM EXECUTION //
//*----------------------*//

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// take stdin as input
	if os.Stdin != nil {
		scanner := bufio.NewScanner(os.Stdin)
		math := ""
		for scanner.Scan() {
			math += scanner.Text()
		}
		output(math)
		// read file given as shell arg
	} else if os.Args[1] == "-f" {
		file, err := os.Open(os.Args[2])
		check(err)
		var text []byte
		file.Read(text)
		math := string(text)
		output(math)
		// otherwise simply parse the shell args as the expression
	} else {
		math := strings.Join(os.Args[1:], " ")
		output(math)
	}
}

// prints the results of the replace function
func output(math string) {
	fmt.Println(replace(math))
}
