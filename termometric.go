package main

import "fmt"

const ebu_k float64 = 373.15

func main() {
	conv := ebu_k - 273
	fmt.Printf("Temperature in Kelvin: %.2f K\n Temperature in Celsius: %.2f Cº\n", ebu_k, conv)
}
