package main

import (
	"aoc/utils"
	"log"
	"strconv"
	"strings"
)

func main() {
	log.Println("day05")

	part1()
	part2()
}

func part2() {
	log.Println("part2")

	data := `3-5
10-14
16-20
12-18`

	data = utils.ReadInputEasy("2025", "05")
	rangesStr, _, _ := strings.Cut(data, "\n\n")
	ranges := strings.Split(rangesStr, "\n")

	// Parse all ranges first
	allRanges := [][]int{}
	for _, r := range ranges {
		lowerStr, upperStr, _ := strings.Cut(r, "-")
		lower, err := strconv.Atoi(lowerStr)
		upper, err := strconv.Atoi(upperStr)
		if err != nil {
			log.Fatalln(err)
		}
		allRanges = append(allRanges, []int{lower, upper})
	}

	// Sort ranges by start position
	for i := 0; i < len(allRanges); i++ {
		for j := i + 1; j < len(allRanges); j++ {
			if allRanges[i][0] > allRanges[j][0] {
				allRanges[i], allRanges[j] = allRanges[j], allRanges[i]
			}
		}
	}

	// Merge overlapping ranges
	merged := [][]int{allRanges[0]}
	for i := 1; i < len(allRanges); i++ {
		current := allRanges[i]
		lastMerged := merged[len(merged)-1]

		// Check if current overlaps with last merged range
		// Ranges overlap if: current start <= last end + 1 (allowing adjacent ranges to merge)
		if current[0] <= lastMerged[1]+1 {
			// Merge by extending the last merged range
			lastMerged[1] = max(lastMerged[1], current[1])
		} else {
			// No overlap, add as new range
			merged = append(merged, current)
		}
	}

	// Count all values in merged ranges
	count := 0
	for _, elem := range merged {
		count += elem[1] - elem[0] + 1
	}

	log.Println("Total count:", count)
}

func part1() {
	log.Println("part1")

	data := utils.ReadInputEasy("2025", "05")
	//
	// 	data = `3-5
	// 10-14
	// 16-20
	// 12-18
	//
	// 1
	// 5
	// 8
	// 11
	// 17
	// 32`

	rangesStr, ingIds, _ := strings.Cut(data, "\n\n")
	ranges := strings.Split(rangesStr, "\n")
	ing := [][]int{}
	for _, r := range ranges {
		lowerStr, upperStr, _ := strings.Cut(r, "-")
		lower, err := strconv.Atoi(lowerStr)
		upper, err := strconv.Atoi(upperStr)
		if err != nil {
			log.Fatalln(err)
		}

		// could sort in insert to have binary search below
		ing = append(ing, []int{lower, upper})

		// assuming not so big ingredient list
		// load a map of all possible fresh ingredients
		// then check for membership.
		// for i := lower; i <= upper; i++ {
		// 	ing[i] = 1
		// }
	}
	count := 0
outer:
	for ingId := range strings.SplitSeq(ingIds, "\n") {
		id, err := strconv.Atoi(ingId)
		if err != nil {
			log.Fatalln(err)
		}

		for _, r := range ing {
			lower, upper := r[0], r[1]

			if id < lower || id > upper {
				continue
			}
			count += 1
			continue outer
		}

		// _, ok := ing[id]
		//
		// if ok {
		// 	count += 1
		// }

	}

	log.Println(count)
}
