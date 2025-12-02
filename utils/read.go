package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// not sure of a better name😂
func ReadInputEasy(year, day string) string {
	file := fmt.Sprintf("./%s/day%s/input.txt", year, day)

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
