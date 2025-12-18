package main

import "fmt"

func main() {
	Alphabet()
}

func Alphabet() {
	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("%c", i)
	}
	fmt.Printf("\n")
}
