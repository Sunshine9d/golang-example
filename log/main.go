package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()
	log.SetOutput(file)
	log.Println("Logging to a file in Go!")

	vars := []int{1, 3, 4, 2, 34, 43, 232, 23}
	sort.Ints(vars)
	fmt.Println(vars)

}
