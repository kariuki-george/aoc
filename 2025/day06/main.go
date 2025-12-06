package main

import (
	"aoc/utils"
	"log"
	"strconv"
	"strings"
)

func main() {
	log.Println("day06")

	part1()
	part2()
}

func part2() {
	log.Println("part 2")

	data := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

	data = utils.ReadInputEasy("2025", "06")

	//parse opts line

	splited := strings.Split(data, "\n")
	opsstr := splited[len(splited)-1]
	opsstr, _, _ = strings.Cut(opsstr, "\n")
	runes := []rune(opsstr)

	rows := splited[:len(splited)-1]

	//data to read
	getLength := func(offset int) (int, bool) {
		length := 0
		for i := offset + 1; i < len(runes); i++ {
			r := runes[i]
			if r != ' ' {
				length = i - offset - 1
				return length, true
			}
		}
		return 0, false
	}

	offset := 0
	total := 0
	lastProcessed := false
	for {
		if lastProcessed {
			break
		}
		readN, found := getLength(offset)

		if !found {
			readN = len(rows[0]) - offset
			lastProcessed = true
		}
		subtotal := 0

		operator := runes[offset]

		if operator == '*' {
			subtotal = 1
		}

		for i := readN - 1; i >= 0; i-- {
			digitsStr := ""
			for _, row := range rows {
				runes := []rune(row)

				digit := runes[offset+i]

				if digit == ' ' {
					continue
				}
				digitsStr += string(digit)
			}
			if digitsStr == "" {
				digitsStr = "0"
			}
			val, err := strconv.Atoi(string(digitsStr))
			if err != nil {
				log.Fatalln(err)
			}

			if operator == '*' {
				subtotal *= val
			} else {
				subtotal += val
			}

		}
		total += subtotal
		offset += readN + 1
	}

	log.Println(total)
}

func part1() {
	log.Println("part 1")

	data := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   + `

	data = utils.ReadInputEasy("2025", "06")

	rowsSeq := strings.Lines(data)
	rows := [][]string{}
	t := 0
	for row := range rowsSeq {
		parsed := []string{}
		for r := range strings.SplitSeq(row, " ") {
			if r == "\n" || r == " " || r == "" {
				continue
			}
			parsed = append(parsed, strings.TrimSpace(r))
		}
		rows = append(rows, parsed)
	}
	for index, _ := range rows[0] {
		total := 0
		// check sign
		operator := strings.TrimSpace(rows[len(rows)-1][index])
		if operator == "*" {
			total = 1
		}

		for i := range len(rows) {
			if i == len(rows)-1 {
				break
			}

			valStr := rows[i][index]

			val, err := strconv.Atoi(valStr)
			if err != nil {
				log.Fatalln(err)
			}

			if operator == "*" {
				total *= val
			} else {
				total += val
			}

		}
		t += total
	}
	log.Println(t)
}
