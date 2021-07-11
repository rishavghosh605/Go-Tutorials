package main

import ("fmt"
		"math/rand")//math.rand is imported like this in go

//Functions in the package all start with a capital letter
func foo() {
	fmt.Println("A random number from 1-100",rand.Intn(100))
}
func main (){
	foo()
}