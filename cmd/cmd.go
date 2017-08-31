package main

import (
	".."

	"fmt"
)

/*
	Make a command line script:	read comma separated numbers from stdin and check are this numbers numerical progression.
*/

func main() {
	resultMessage := progression.IsNotProgression
	if progression.IsArgsProgression() {
		resultMessage = progression.IsProgression
	}

	fmt.Println(resultMessage)
}
