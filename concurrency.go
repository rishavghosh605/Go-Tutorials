//Concurrency is not parrallelism which denotes that one an do different jobs 
// one at a time constantly changing according to needs rather than doing everything at once
//Basically meaning of async
// SO here we use concurrency using goroutine
package main

import ("fmt"
        "time")

func say(s string){
	for i:=0; i<3; i++{
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
	}
}

func main(){
	say("Hey")
	say("There")
	fmt.Println("--------------")
	go say("Hey")
	say("There")
	fmt.Println("--------------")
	go say("Hey")
	go say("There")
	// Note: When both are made as go routines then they essentially become non blocking code thus resulting in the unique situation of the program ending before the goroutines themselves thus nothing is returned in the output
	//Quick solve for such issues is to simply ass a time.Sleep
}