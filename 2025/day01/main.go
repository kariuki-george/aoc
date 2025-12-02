package main

import (
	"aoc/utils"
	"log"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part2() {
	log.Println("part 2")
	data := utils.ReadInputEasy("2025", "01")
	// data = "L1000"
	// 	data = `L68
	// L30
	// R48
	// L5
	// R60
	// L55
	// L1
	// L99
	// R14
	// L82`

	commands := strings.Split(data, "\n")

	pointer := 50
	pass := 0

	for _, command := range commands {
		runes := []rune(strings.TrimSpace(command))

		direction, distanceStr := runes[0], string(runes[1:])
		distance, err := strconv.Atoi(distanceStr)

		if err != nil {
			log.Fatal(err)
		}

		// limit distance to 0-99
		factor := distance / 100
		distance = distance % 100
		pointerStart := pointer
		pass += factor
		if direction == rune('L') {
			pointer -= distance
		} else {
			pointer += distance
		}

		if pointer == 0 {
			pass += 1
		}

		if pointer == 100 {
			pass += 1
			pointer = 0
		}
		if pointer < 0 {
			pointer += 100

			if pointerStart != 0 {
				pass += 1
			}

		}
		if pointer > 100 {
			pass += 1
			pointer -= 100
		}

	}
	log.Println(pass)

}

func part1() {
	log.Println("part 1")
	data := utils.ReadInputEasy("2025", "01")
	commands := strings.Split(data, "\n")

	pointer := 50
	pass := 0
	for _, command := range commands {

		runes := []rune(strings.TrimSpace(command))

		direction, distanceStr := runes[0], string(runes[1:])
		distance, err := strconv.Atoi(distanceStr)

		if err != nil {
			log.Fatal(err)
		}

		if direction == rune('L') {
			pointer -= distance
		} else {
			pointer += distance
		}
		pointer = pointer % 100
		if pointer == 0 {
			pass += 1
		}

	}
	log.Println(pass)

}
