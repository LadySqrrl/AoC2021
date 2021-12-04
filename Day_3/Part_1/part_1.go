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

	firstDigit0 := 0
	secondDigit0 := 0
	thirdDigit0 := 0
	fourthDigit0 := 0
	fifthDigit0 := 0
	sixthDigit0 := 0
	seventhDigit0 := 0
	eighthDigit0 := 0
	ninthDigit0 := 0
	tenthDigit0 := 0
	eleventhDigit0 := 0
	twelfthDigit0 := 0
	firstDigit1 := 0
	secondDigit1 := 0
	thirdDigit1 := 0
	fourthDigit1 := 0
	fifthDigit1 := 0
	sixthDigit1 := 0
	seventhDigit1 := 0
	eighthDigit1 := 0
	ninthDigit1 := 0
	tenthDigit1 := 0
	eleventhDigit1 := 0
	twelfthDigit1 := 0

	for i := 0; i < len(strList); i++ {
		r := []rune(strList[i])

		if r[11] == '0' {
			firstDigit0++
		} else {
			firstDigit1++
		}

		if r[10] == '0' {
			secondDigit0++
		} else {
			secondDigit1++
		}

		if r[9] == '0' {
			thirdDigit0++
		} else {
			thirdDigit1++
		}

		if r[8] == '0' {
			fourthDigit0++
		} else {
			fourthDigit1++
		}

		if r[7] == '0' {
			fifthDigit0++
		} else {
			fifthDigit1++
		}

		if r[6] == '0' {
			sixthDigit0++
		} else {
			sixthDigit1++
		}

		if r[5] == '0' {
			seventhDigit0++
		} else {
			seventhDigit1++
		}

		if r[4] == '0' {
			eighthDigit0++
		} else {
			eighthDigit1++
		}

		if r[3] == '0' {
			ninthDigit0++
		} else {
			ninthDigit1++
		}

		if r[2] == '0' {
			tenthDigit0++
		} else {
			tenthDigit1++
		}

		if r[1] == '0' {
			eleventhDigit0++
		} else {
			eleventhDigit1++
		}

		if r[0] == '0' {
			twelfthDigit0++
		} else {
			twelfthDigit1++
		}
	}

	gamma := 0b0
	epsilon := 0b0

	if firstDigit1 > firstDigit0 {
		gamma++
	} else {
		epsilon++
	}
	if secondDigit1 > secondDigit0 {
		gamma += 2
	} else {
		epsilon += 2
	}
	if thirdDigit1 > thirdDigit0 {
		gamma += 4
	} else {
		epsilon += 4
	}
	if fourthDigit1 > fourthDigit0 {
		gamma += 8
	} else {
		epsilon += 8
	}
	if fifthDigit1 > fifthDigit0 {
		gamma += 16
	} else {
		epsilon += 16
	}
	if sixthDigit1 > sixthDigit0 {
		gamma += 32
	} else {
		epsilon += 32
	}
	if seventhDigit1 > seventhDigit0 {
		gamma += 64
	} else {
		epsilon += 64
	}
	if eighthDigit1 > eighthDigit0 {
		gamma += 128
	} else {
		epsilon += 128
	}
	if ninthDigit1 > ninthDigit0 {
		gamma += 256
	} else {
		epsilon += 256
	}
	if tenthDigit1 > tenthDigit0 {
		gamma += 512
	} else {
		epsilon += 512
	}
	if eleventhDigit1 > eleventhDigit0 {
		gamma += 1024
	} else {
		epsilon += 1024
	}
	if twelfthDigit1 > twelfthDigit0 {
		gamma += 2048
	} else {
		epsilon += 2048
	}

	product := gamma * epsilon

	fmt.Println(product)
}
