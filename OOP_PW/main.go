package main

import (
	"OOP_PW/pixel"
	"fmt"
)

func main() {
	myPixel := pixel.CreatePixel(10, 2.4, 6.2)
	myPixel.ToString()
	//fmt.Printf("%d\n", myPixel.luminance)
	fmt.Println(myPixel.GetLuminance())
	myPixel.SetLuminance(20)
	fmt.Println(myPixel.GetLuminance())

	mySecondPixel := myPixel
	mySecondPixel.SetLuminance(2137)
	mySecondPixel.ToString()
	myPixel.ToString()
}
