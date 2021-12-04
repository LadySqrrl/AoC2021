// Package main is the main fucntion
package main

import (
	"bufio"
	"fmt"
	"os"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	strList, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}

	firstDigit0, secondDigit0, thirdDigit0, fourthDigit0 := 0, 0, 0, 0
	firstDigit1, secondDigit1, thirdDigit1, fourthDigit1 := 0, 0, 0, 0

	for i := 0; i < len(strList); i++ {
		r := []rune(strList[i])

		if r[3] == '0' {
			firstDigit0++
		} else {
			firstDigit1++
		}

		if r[2] == '0' {
			secondDigit0++
		} else {
			secondDigit1++
		}

		if r[1] == '0' {
			thirdDigit0++
		} else {
			thirdDigit1++
		}

		if r[0] == '0' {
			fourthDigit0++
		} else {
			fourthDigit1++
		}
	}
	fmt.Println(firstDigit0, ", ", firstDigit1, "; ", secondDigit0, ", ", secondDigit1, "; ", thirdDigit0, ", ", thirdDigit1, "; ", fourthDigit0, ", ", fourthDigit1)

	common1, common2, common3, common4 := "", "", "", ""
	less1, less2, less3, less4 := "", "", "", ""

	if firstDigit1 >= firstDigit0 {
		common1 = "1"
		less1 = "0"
	} else {
		common1 = "0"
		less1 = "1"
	}
	if secondDigit1 >= secondDigit0 {
		common2 = "1"
		less2 = "0"
	} else {
		common2 = "0"
		less2 = "1"
	}
	if thirdDigit1 >= thirdDigit0 {
		common3 = "1"
		less3 = "0"
	} else {
		common3 = "0"
		less3 = "1"
	}
	if fourthDigit1 >= fourthDigit0 {
		common4 = "1"
		less4 = "0"
	} else {
		common4 = "0"
		less4 = "1"
	}

	fmt.Println(common1, common2, common3, common4)
	fmt.Println(less1, less2, less3, less4)
}
