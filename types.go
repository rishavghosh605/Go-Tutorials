package main

import ("fmt")//short form for format

func add(x,y float64) float64 {
	return x+y
}

const a int = 4// random declaration of a constant type

func multiple (a,b string) (string,string){
	return a,b
}
func main() {
	// var num1 float64 = 5.6
	// var num2 float64 = 9.5
	num1,num2 := 5.6,9.5 // Go uses this type of assigment operator to dynamically assign the type
	// var num1,num2 float64 = 5.6,9.5

	fmt.Println(add(num1,num2))

	w1,w2 := "hey", "there"

	fmt.Println(multiple(w1,w2))

	// Some type casting stuff
	var s int = 84
	var d float64 = float64(a)

	x := a // x will be type int
}