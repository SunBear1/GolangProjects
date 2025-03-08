package main

import (
	"Singletons/furniture"
	"fmt"
)

func main() {
	t1 := furniture.GetInstance()
	fmt.Println("t1: " + t1.GetName())
	t1.SetName("fancy table")
	fmt.Println("t1: (after name change) " + t1.GetName())
	t2 := furniture.GetInstance()
	fmt.Println("t2: " + t2.GetName())
	t2.SetName("even more fancy table")
	fmt.Println("t2: (after name change) " + t2.GetName())
	fmt.Println("t1: (after name change in t2) " + t1.GetName())
}
