package main

import "fmt"

const ebu_k float64 = 373.15

func main() {
	conv := ebu_k - 273
	fmt.Printf("Temperature conversion = %.2f to = %.2f Cº\n", ebu_k, conv)
}
