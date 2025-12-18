package main

import "fmt"

func main() {
	ReverseAlphabet()
}

func ReverseAlphabet() {
	for i := 'z'; i >= 'a'; i-- {
		fmt.Printf("%c", i)
	}
	fmt.Printf("\n")
}
