package main

import "fmt"

func Cat() string {
	return "Fuck Jason"
}

func main() {
	saying := Cat()
	fmt.Print(saying)
}