package main

import "fmt"

func main() {
	floatStack := NewStack[float32]()
	_ = floatStack.push(2.5)
	fTop, _ := floatStack.pop()
	fmt.Printf("%f", fTop)

	stringStack := NewStack[string]()
	_ = stringStack.push("2.5")
	sTop, _ := stringStack.pop()
	fmt.Printf("%s", sTop)
}
