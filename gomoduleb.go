package gomoduleb

import "fmt"

func init() {
	fmt.Println("Loading module B")
}

func FctFromModuleB() {
	fmt.Println("Calling FctFromModuleB()")
}
