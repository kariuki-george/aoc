package main

import (
	"aoc/utils"
	"log"
	"strings"
)

func main() {
	log.Println("day 04")
	part1()
	part2()
}

func part2() {
	log.Println("part02")

	data := utils.ReadInputEasy("2025", "04")
	// 	data = `..@@.@@@@.
	// @@@.@.@.@@
	// @@@@@.@.@@
	// @.@@@@..@.
	// @@.@@@@.@@
	// .@@@@@@@.@
	// .@.@.@.@@@
	// @.@@@.@@@@
	// .@@@@@@@@.
	// @.@.@@@.@.`

	dim, _, _ := strings.Cut(data, "\n")
	cols := len(dim)
	data = strings.ReplaceAll(data, "\n", "")

	total := 0

	// outer:
	for {
		picked := []int{}
		for index := range len(data) {
			if string(data[index]) != "@" {
				// skip if no roll
				continue
			}
			count := 0
			for _, dir := range []string{"top", "bottom", "right", "left", "bottomright", "bottomleft", "topright", "topleft"} {
				if roll := rollExists(data, cols, index, dir); roll {
					count += 1
				}
			}
			if count < 4 {
				picked = append(picked, index)
			}
		}

		// round complete
		// replace removables
		if len(picked) == 0 {
			break
		}
		total += len(picked)

		for _, index := range picked {
			data = replaceAtIndex(data, index, 'x')
		}

	}

	log.Println(total)

}

func part1() {
	log.Println("part01")

	data := utils.ReadInputEasy("2025", "04")
	//

	// 	data = `..@@.@@@@.
	// @@@.@.@.@@
	// @@@@@.@.@@
	// @.@@@@..@.
	// @@.@@@@.@@
	// .@@@@@@@.@
	// .@.@.@.@@@
	// @.@@@.@@@@
	// .@@@@@@@@.
	// @.@.@@@.@.`

	dim, _, _ := strings.Cut(data, "\n")
	cols := len(dim)
	data = strings.ReplaceAll(data, "\n", "")

	total := 0

	for index := range len(data) {
		if string(data[index]) != "@" {
			// skip if no roll
			continue
		}
		count := 0
		for _, dir := range []string{"top", "bottom", "right", "left", "bottomright", "bottomleft", "topright", "topleft"} {
			if roll := rollExists(data, cols, index, dir); roll {
				count += 1
			}
		}
		if count < 4 {
			total += 1
		}
	}

	log.Println(total)

}

func replaceAtIndex(s string, index int, newRune rune) string {
	runes := []rune(s)
	if index >= 0 && index < len(runes) {
		runes[index] = newRune
	}
	return string(runes)
}

func rollExists(data string, cols int, startPos int, direction string) bool {

	// check if start position exists
	if startPos < 0 || startPos >= len(data) {
		return false
	}

	row := startPos / cols
	col := startPos % cols
	pos := startPos

	switch direction {
	case "top":
		pos = ((row - 1) * cols) + col
	case "bottom":
		pos = ((row + 1) * cols) + col
	case "left":
		pos = ((row) * cols) + col - 1
		if pos < (cols * row) {
			return false
		}
	case "right":
		pos = ((row) * cols) + col + 1
		if pos >= (cols*row)+cols {
			return false
		}

	case "topleft":
		pos = ((row - 1) * cols) + col - 1
		if pos < (row-1)*cols {
			return false
		}

	case "topright":
		pos = ((row - 1) * cols) + col + 1
		if pos >= (row-1)*cols+cols {
			return false
		}

	case "bottomleft":
		pos = ((row + 1) * cols) + col - 1
		if pos < (row+1)*cols {
			return false
		}
	case "bottomright":
		pos = ((row + 1) * cols) + col + 1
		if pos >= (row+1)*cols+cols {
			return false
		}

	}

	// check if pos exists

	if pos < 0 || pos >= len(data) {
		return false
	}
	return string(data[pos]) == "@"
}
