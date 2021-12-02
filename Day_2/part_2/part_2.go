// Package main is the main fucntion
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	x := 0
	y := 0
	aim := 0
	product := 0

	for i := 0; i < len(strList); i++ {
		dir := strings.Split(strList[i], " ")

		switch dir[0] {
		case "forward":
			fwd, err := strconv.Atoi(dir[1])
			if err != nil {
				fmt.Println("Error converting string to int")
			}
			x += fwd
			y += (fwd * aim)
		case "up":
			up, err := strconv.Atoi(dir[1])
			if err != nil {
				fmt.Println("Error converting string to int")
			}
			aim -= up
		case "down":
			dwn, err := strconv.Atoi(dir[1])
			if err != nil {
				fmt.Println("Error converting string to int")
			}
			aim += dwn
		}
	}

	product = x * y

	fmt.Println(product)
}
