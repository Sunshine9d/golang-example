package lab

import (
	"fmt"
)

import (
	"bufio"
	"os"
)

func ReadLines(filename string) ([]string, error) {
	file, err := os.Open("/Users/a01/golang-example/lab2-core/" + filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil

}

func Lab2() {
	lines, err := ReadLines("data.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, line := range lines {
			fmt.Println(line)
		}
	}

}
