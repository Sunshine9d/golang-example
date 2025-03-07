package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sumFromFile() {
	file, err := os.Open("/root/code/xerox/output.txt")
	if err != nil {
		return
	}
	defer file.Close()

	var sum = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rs, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		sum += rs
		//lines = append(lines, scanner.Text())
	}
	fmt.Println(sum)
}
