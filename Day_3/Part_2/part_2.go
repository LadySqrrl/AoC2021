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

	var common1, common2, common3, common4, common5, common6, common7, common8, common9, common10, common11, common12 string

	if firstDigit1 >= firstDigit0 {
		common1 = "1"
	} else {
		common1 = "0"
	}
	if secondDigit1 >= secondDigit0 {
		common2 = "1"
	} else {
		common2 = "0"
	}
	if thirdDigit1 >= thirdDigit0 {
		common3 = "1"
	} else {
		common3 = "0"
	}
	if fourthDigit1 >= fourthDigit0 {
		common4 = "1"
	} else {
		common4 = "0"
	}
	if fifthDigit1 >= fifthDigit0 {
		common5 = "1"
	} else {
		common5 = "0"
	}
	if sixthDigit1 >= sixthDigit0 {
		common6 = "1"
	} else {
		common6 = "0"
	}
	if seventhDigit1 >= seventhDigit0 {
		common7 = "1"
	} else {
		common7 = "0"
	}
	if eighthDigit1 >= eighthDigit0 {
		common8 = "1"
	} else {
		common8 = "0"
	}
	if ninthDigit1 >= ninthDigit0 {
		common9 = "1"
	} else {
		common9 = "0"
	}
	if tenthDigit1 >= tenthDigit0 {
		common10 = "1"
	} else {
		common10 = "0"
	}
	if eleventhDigit1 >= eleventhDigit0 {
		common11 = "1"
	} else {
		common11 = "0"
	}
	if twelfthDigit1 >= twelfthDigit0 {
		common12 = "1"
	} else {
		common12 = "0"
	}

	commonStr := common12 + common11 + common10 + common9 + common8 + common7 + common6 + common5 + common4 + common3 + common2 + common1
	fmt.Println(commonStr)

	oxygen := make([]string, len(strList))
	c02 := make([]string, len(strList))

	copyOxygen := copy(oxygen, strList)
	copyc02 := copy(c02, strList)
	fmt.Println(copyOxygen, copyc02)

	for j := 0; j < 11; j++ {
		rOxy := []rune(oxygen[j])
		rCommon := []rune(commonStr)
		for i := 0; i < len(oxygen)-1; i++ {
			if rOxy[j] != rCommon[j] {
				oxygen = append(oxygen[:i], oxygen[i+1:]...)
			} else {
				c02 = append(c02[:i], c02[i+1:]...)
			}
		}
		if rOxy[j] != rCommon[j] {
			oxygen = append(oxygen[:len(oxygen)-1])
		} else {
			c02 = append(c02[:len(c02)-1])
		}
	}

	for i := 0; i < len(oxygen)-1; i++ {
		rOxy := []rune(oxygen[len(oxygen)-1])
		rCommon := []rune(commonStr)
		if rOxy[11] != rCommon[11] {
			oxygen = append(oxygen[:i])
		} else {
			c02 = append(c02[:i])
		}
	}

	fmt.Println(oxygen, "\n", c02)
}
