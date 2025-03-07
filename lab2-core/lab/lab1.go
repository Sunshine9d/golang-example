package lab

import (
	"fmt"
	"strings"
)

func WordCount(s string, word string) int {
	return strings.Count(s, word)
}

func Lab1() {
	count := WordCount("hello, Hello how have you been in helloworld", "hello")
	fmt.Println(count)
}
