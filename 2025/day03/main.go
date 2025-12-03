package main

import (
	"aoc/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	log.Println("day03")
	part1()
	part2()
}
func part2() {
	log.Println("part2")
	data := utils.ReadInputEasy("2025", "03")

	// data = `987654321111111
	// 811111111111119
	// 234234234234278
	// 818181911112111`

	banks := strings.SplitSeq(data, "\n")
	sum := 0
	for bank := range banks {
		runes := []rune(bank)
		val := ""
		// just check for the bigggest number

		// val, runes = biggerChecker(runes, 2)
		for j := 12; j > 0; j-- {
			maxval, r := biggerChecker(runes, j)
			val = fmt.Sprintf("%s%d", val, maxval)
			runes = r
		}

		j, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		sum += j
	}
	log.Println(sum)
}

// don't fault the name😂
func biggerChecker(runes []rune, remainingJs int) (int, []rune) {
	for num := 9; num >= 0; num-- {

		for index, rune := range runes[:len(runes)-(remainingJs-1)] {
			if string(rune) == fmt.Sprintf("%d", num) {
				//largest val
				return num, runes[index+1:]
			}
		}

	}
	return 0, nil
}

func part1() {
	log.Println("part1")

	data := utils.ReadInputEasy("2025", "03")

	// data = `987654321111111
	// 811111111111119
	// 234234234234278
	// 818181911112111`

	banks := strings.SplitSeq(data, "\n")
	sum := 0
	for bank := range banks {
		runes := []rune(bank)
		val := ""
		// just check for the bigggest number
	outer:
		for num := 9; num >= 0; num-- {

			for index, rune := range runes {
				// skip if the rune is the last rune
				if index == len(runes)-1 {
					break
				}
				if string(rune) == fmt.Sprintf("%d", num) {
					// check for the second largest number
					// could separate this into a reusable func

					largest := 0
					for _, rune := range runes[index+1:] {

						num, err := strconv.Atoi(string(rune))
						if err != nil {
							log.Fatal(err)
						}

						if num > largest {
							largest = num
						}

					}

					val = fmt.Sprintf("%s%d", string(rune), largest)
					break outer
				}
			}

		}
		j, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		sum += j
	}
	log.Println(sum)
}
