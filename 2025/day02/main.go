package main

import (
	"aoc/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	log.Println("day02")

	part1()
	part2()
}

func part2() {

	data := utils.ReadInputEasy("2025", "02")

	// data = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

	ranges := strings.Split(data, ",")

	vals := make(map[int]int)

	for _, prange := range ranges {

		//parse
		bounds := strings.Split(prange, "-")
		smallestStr, largestStr := bounds[0], bounds[1]

		smallest, err := strconv.Atoi(smallestStr)
		largest, err := strconv.Atoi(largestStr)
		if err != nil {
			log.Fatal(err)
		}

		for i := smallest; i <= largest; i++ {
			//check length
			val := fmt.Sprintf("%d", i)
			l := len(val)
			// checkout tries
			for winSize := 1; winSize <= l/2; winSize++ {

				// build a sliding window

				offset := 0
				prevWindow := ""
				for {
					win1Start := offset
					win1End := winSize + offset
					win2Start := win1End
					win2End := win1End + winSize
					offset += winSize
					if win2End > l {
						if win2Start == l {
							// only consider those values whose all windos have been consumed entirely
							vals[i] = i
						}
						break
					}
					win1 := val[win1Start:win1End]
					win2 := val[win2Start:win2End]

					// log.Println("val ", val, " ", win1, " - ", win2, "offset", offset, "winSize", winSize)
					if win1 != win2 {
						break
					}

					if prevWindow != "" && prevWindow != win1 {
						break
					}
					prevWindow = win1

				}

			}

		}

	}
	total := 0
	for _, v := range vals {
		total += v
	}

	log.Println(total)
}
func part1() {

	data := utils.ReadInputEasy("2025", "02")

	// data = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

	ranges := strings.Split(data, ",")

	total := 0

	for _, prange := range ranges {
		//parse
		bounds := strings.Split(prange, "-")
		smallestStr, largestStr := bounds[0], bounds[1]

		smallest, err := strconv.Atoi(smallestStr)
		largest, err := strconv.Atoi(largestStr)
		if err != nil {
			log.Fatal(err)
		}

		for i := smallest; i <= largest; i++ {
			//check length
			val := fmt.Sprintf("%d", i)
			l := len(val)
			if l%2 != 0 {
				continue
			}
			// split two
			if val[:l/2] == val[l/2:] {
				total += i
			}

		}

	}

	log.Println(total)
}
