package gomoduleb

import "fmt"

func init() {
	fmt.Println("Loading module B")
}

func FctFromModuleB() {
	fmt.Println(modB())
}

func modB() string {
	return "Calling modB() v3"
}
