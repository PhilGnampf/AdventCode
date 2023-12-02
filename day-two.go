package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func daytwo() {

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
		erg += powerOfSmallest(line)
	}

	fmt.Println(erg)

	// Close File
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

}

// Challenge 2
func powerOfSmallest(s string) int {
	var smallestRed int = 0
	var smallestBlue int = 0
	var smallestGreen int = 0

	splitAll := func(c rune) bool {
		return c == ':' || c == ',' || c == ';'
	}
	parts := strings.FieldsFunc(s, splitAll)

	for _, part := range parts {
		values := strings.Split(part, " ")
		if values[0] != "Game" {
			fmt.Println(values[1] + " " + values[2])
			valuesInt, _ := strconv.Atoi(values[1])
			if values[2] == "red" && smallestRed < valuesInt {
				smallestRed, _ = strconv.Atoi(values[1])
			}
			if values[2] == "green" && smallestGreen < valuesInt {
				smallestGreen, _ = strconv.Atoi(values[1])
			}
			if values[2] == "blue" && smallestBlue < valuesInt {
				smallestBlue, _ = strconv.Atoi(values[1])
			}
		}
	}
	fmt.Println("---")
	return smallestRed * smallestBlue * smallestGreen
}

// Challenge 1
func isTrue(s string) bool {
	var isSetValid bool = true

	splitAll := func(c rune) bool {
		return c == ':' || c == ',' || c == ';'
	}
	parts := strings.FieldsFunc(s, splitAll)

	for _, part := range parts {
		values := strings.Split(part, " ")
		if values[0] != "Game" {
			if !isValidThrow(values[1], values[2]) {
				isSetValid = false
			}
		}
	}

	return isSetValid
}

func isValidThrow(value string, color string) bool {
	fmt.Println(value + " " + color)
	valueInt, _ := strconv.Atoi(value)
	if color == "red" && valueInt > 12 {
		return false
	}
	if color == "green" && valueInt > 13 {
		return false
	}
	if color == "blue" && valueInt > 14 {
		return false
	}
	return true
}
