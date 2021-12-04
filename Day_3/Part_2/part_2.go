// Package main is the main fucntion
package main

import (
	"bufio"
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

	firstDigit0, secondDigit0, thirdDigit0, fourthDigit0, fifthDigit0, sixthDigit0, seventhDigit0, eighthDigit0, ninthDigit0, tenthDigit0, eleventhDigit0, twelfthDigit0 := 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
	firstDigit1, secondDigit1, thirdDigit1, fourthDigit1, fifthDigit1, sixthDigit1, seventhDigit1, eighthDigit1, ninthDigit1, tenthDigit1, eleventhDigit1, twelfthDigit1 := 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0

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

	common1, common2, common3, common4, common5, common6, common7, common8, common9, common10, common11, common12 := "", "", "", "", "", "", "", "", "", "", "", ""
	less1, less2, less3, less4, less5, less6, less7, less8, less9, less10, less11, less12 := "", "", "", "", "", "", "", "", "", "", "", ""

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
	if fifthDigit1 >= fifthDigit0 {
		common5 = "1"
		less5 = "0"
	} else {
		common5 = "0"
		less5 = "1"
	}
	if sixthDigit1 >= sixthDigit0 {
		common6 = "1"
		less6 = "0"
	} else {
		common6 = "0"
		less6 = "1"
	}
	if seventhDigit1 >= seventhDigit0 {
		common7 = "1"
		less7 = "0"
	} else {
		common7 = "0"
		less7 = "1"
	}
	if eighthDigit1 >= eighthDigit0 {
		common8 = "1"
		less8 = "0"
	} else {
		common8 = "0"
		less8 = "1"
	}
	if ninthDigit1 >= ninthDigit0 {
		common9 = "1"
		less9 = "0"
	} else {
		common9 = "0"
		less9 = "1"
	}
	if tenthDigit1 >= tenthDigit0 {
		common10 = "1"
		less10 = "0"
	} else {
		common10 = "0"
		less10 = "1"
	}
	if eleventhDigit1 >= eleventhDigit0 {
		common11 = "1"
		less11 = "0"
	} else {
		common11 = "0"
		less11 = "1"
	}
	if twelfthDigit1 >= twelfthDigit0 {
		common12 = "1"
		less12 = "0"
	} else {
		common12 = "0"
		less12 = "1"
	}

	var oxygen []string
	var c02 []string

	for j := 0; j < 12; j++ {
		goodStr := ""
		c := 'a'

		for i := 0; i < len(strList); i++ {
			r := []rune(strList[i])
			switch j {
			case 1:
				c = []rune(common1[1])
				break
			case 2:
				c = []rune(common2)
				break
			case 3:
				c = []rune(common3)
				break
			case 4:
				c = []rune(common4)
				break
			case 5:
				c = []rune(common5)
				break
			case 6:
				c = []rune(common6)
				break
			case 7:
				c = []rune(common7)
				break
			case 8:
				c = []rune(common8)
				break
			case 9:
				c = []rune(common9)
				break
			case 10:
				c = []rune(common10)
				break
			case 11:
				c = []rune(common11)
				break
			case 12:
				c = []rune(common12)
				break
			}
		}
	}
}
