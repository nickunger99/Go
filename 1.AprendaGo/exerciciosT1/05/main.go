package main

import "fmt"

type hotdog int

var x hotdog

var y int

func main() {
	fmt.Print("x: ")
	fmt.Print(x, " ")
	fmt.Printf("%T\t", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Print("y: ")
	fmt.Printf("%v\t%T", y, y)
}
