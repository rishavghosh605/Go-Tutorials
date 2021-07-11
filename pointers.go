package main

import "fmt"//short form for format

func main() {
	x:=15
	a:=&x // a storing the memory address of x
	fmt.Println(a) // Prints the address of a
	fmt.Println(*a) // Prints the value of a
	*a = 5
	fmt.Println(x)
}