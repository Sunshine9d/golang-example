package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func checkEven(i int) bool {
	return i%2 == 0
}

func main() {
	var password string
	fmt.Scanln(&password)
	fmt.Println("Password encrypted to: ", encodeWithMD5(password))

	s := "Hello World, 世界"
	fmt.Println(s, "ASCII: ", isASCII(s))
	fmt.Println(len(s))
	fmt.Println([]byte(s))
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ { // ASCII strings are single-byte
		if s[i] > 127 { // ASCII range: 0-127
			return false
		}
	}
	return true
}

func encodeWithMD5(password string) any {
	var hashCode = md5.Sum([]byte(password))
	fmt.Println(hashCode)
	return hex.EncodeToString(hashCode[:])
}
