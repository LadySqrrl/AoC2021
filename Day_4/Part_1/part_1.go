// Package main is the main fucntion
package main

import (
	"bufio"
	"fmt"
	"os"

	//"strconv"
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

	//draws := strings.Split(strList[0], ",")
	var card [5][5]string
	//var found [5][5]bool

	//for x := 0; x < len(draws); x++ {
	//currDraw := draws[x]
	for i := 1; i < len(strList); i++ {
		line := strings.Split(strList[i], " ")
		for j := 0; j < len(line); j++ {
			if line[0] != "Z" {
				card[i-1][j] = line[j]
			}
		}
	}
	fmt.Println(card)
	//}

}
