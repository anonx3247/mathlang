package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if os.Stdin != nil {
		scanner := bufio.NewScanner(os.Stdin)
		math := ""
		for scanner.Scan() {
			math += scanner.Text()
		}
		output(math)
	} else if os.Args[1] == "-f" {
		file, err := os.Open(os.Args[2])
		check(err)
		var text []byte
		file.Read(text)
		math := string(text)
		output(math)
	} else {
		math := strings.Join(os.Args[1:], " ")
		output(math)
	}
}

func check(err error) {
	if err != nil {

		panic(err)
	}
}

func output(math string) {
	fmt.Printf(replace(math))
}
