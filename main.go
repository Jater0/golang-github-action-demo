package main

import "fmt"

func Cat() string {
	return "Fuck me"
}

func main() {
	saying := Cat()
	fmt.Print(saying)
}