package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(filename string) ([]string, error) {
	file, err := os.Open("/Users/a01/golang-example/code-package/" + filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fmt.Println(lines)
	return lines, scanner.Err()
}

func RunMain1() {
	fmt.Println("Running main")
	lines, err := ReadLines("data.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, line := range lines {
			fmt.Println(line)
		}
	}

}
