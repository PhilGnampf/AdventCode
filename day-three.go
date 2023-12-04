package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	
	// Open File
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Scan File
	scanner := bufio.NewScanner(file)

	var input [][]rune	

	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		input = append(input, runes)
	}

	printSolution(findNumbers(input))

	// Close File
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

}

func findNumbers(input [][]rune) [][]int{
	var solution [][]int
	var foundNumbersIterator int = 0
	var inDigitSequence bool = false
	var tmpDigitSequence string
	var tmpSequenceStart int
	var tmpErr error
	
	for x := 0; x < len(input); x++ {
		for y := 0; y < len(input[x]); y++ {
			if unicode.IsDigit(input[x][y]) {
				if inDigitSequence {
					tmpDigitSequence += string(input[x][y])
				} else {
					tmpDigitSequence = string(input[x][y])
					tmpSequenceStart = y
				}

				inDigitSequence = true
			}
			
			if !unicode.IsDigit(input[x][y]) && inDigitSequence {
				fmt.Println(x)
				solution[foundNumbersIterator][0] = x
				solution[foundNumbersIterator][1] = tmpSequenceStart
				solution[foundNumbersIterator][2], tmpErr= strconv.Atoi(tmpDigitSequence)
				solution[foundNumbersIterator][3] = y - tmpSequenceStart
				if tmpErr != nil {
					log.Fatal(tmpErr)
				}
				foundNumbersIterator++
				inDigitSequence = false
			}
		}
	}

	return solution
}

func printSolution(solution [][]int) {
	for _, line := range solution {
		for _, number := range line {
			print(number)
		}
		println()
	}
}