package main

import "fmt"

func main() {
	var a int = 46
	var b int = 6784
	var c int = 614
	var V int = (a * b * c)
	var S int = (2 * (a*b + b*c + c*a))
	fmt.Println(V, S)
}
