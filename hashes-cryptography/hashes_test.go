package main

import "testing"

func TestCheckEven(t *testing.T) {
	i := 10
	expected := false
	res := checkEven(i)
	if res != expected {
		t.Errorf("Expected %v but got %v", true, expected)
	}

}
