package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// not sure of a better name😂
func ReadInputEasy(year, day uint) string {
	file := fmt.Sprintf("./%d/day0%d/input.txt", year, day)

	return ReadInput(file)
}

func ReadInput(file string) string {

	data, err := os.ReadFile(file)

	if err != nil {
		// probably a better way to handle this
		log.Fatal(err)
	}

	return strings.TrimSpace(string(data))
}
