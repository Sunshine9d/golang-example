package lab

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Lab3() {
	file, err := os.Open("/Users/a01/golang-example/lab2-core/input.txt")
	if err != nil {
		return
	}
	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numb, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		sum += numb
	}
	fmt.Println(sum)
}
