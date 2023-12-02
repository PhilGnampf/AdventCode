package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func dayone() {
	// Open File
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Scan File
	scanner := bufio.NewScanner(file)

	var erg int

	for scanner.Scan() {
		line := scanner.Text()
		line = getNumberString(line)
		erg = erg + getNumber(line)
	}

	fmt.Println(erg)

	// Close File
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

}

// Challenge 1
func getNumber(s string) int {
	var first, last int
	foundFirst := false

	for _, c := range s {
		if unicode.IsDigit(c) {
			num := int(c - '0')
			if !foundFirst {
				first = num
				foundFirst = true
			}
			last = num
		}
	}

	return first*10 + last
}

// Challenge 2
func getNumberString(s string) string {
	numberString := ""

	wordToDigit := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"zero":  "0",
	}
	fmt.Println(s)
	for pos, cString := range s {
		if unicode.IsDigit(cString) {
			numberString = numberString + string(cString)
		} else {

			for key, value := range wordToDigit {
				var valid bool = true
				for pos2, cKey := range key {
					if pos+pos2 >= len(s) {
						valid = false
						break
					}
					if cKey != rune(s[pos+pos2]) {
						valid = false
					}
				}
				if valid {
					numberString = numberString + value
					fmt.Println(value)
					fmt.Println(numberString)
				}
			}

		}
	}

	return numberString
}
