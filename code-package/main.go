package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func process(i int) error {
	if i%2 == 0 {
		return fmt.Errorf("Error processing %v", i)
	}
	return nil
}
func checkError(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println("Operation successful")
}

func main() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 4)
	//only add 8 bytes to the slice
	//r.Read(b)
	//fmt.Println(b)
	//fmt.Println(string(b))
	// read the string in 8 byte chunks
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err != nil {
			break
		}
	}

	path := filepath.Join("/Users/a01/golang-example/code-package", "note.txt")
	fmt.Println(path)
	pathNote := "/Users/a01/golang-example/code-package/note.txt"
	fileInfo, err := os.Stat(pathNote)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("info", fileInfo.Name())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.ModTime())
	data, err := os.ReadFile(pathNote)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("data: ", string(data))
	fmt.Println("************")
	data1, err := os.Open(pathNote)
	note := make([]byte, 8)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		n, err := data1.Read(note)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, note)
		fmt.Printf("b[:n] = %q\n", note[:n])
		if err != nil {
			break
		}
	}

	checkError(process(2))
	checkError(process(3))
}
