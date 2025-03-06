package main

import (
	"fmt"
	"github.com/Sunshine9d/golang-example/cryptit/decrypt"
	"github.com/Sunshine9d/golang-example/cryptit/encrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("Kodekloud")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
}
