// You can think of it as an counterpart of try and catch in python


package main

import ("time"
        "fmt"
	    "sync")

var wg sync.WaitGroup

func cleanup(){
	defer wg.Done()
	if r:= recover(); r!=nil{// We assign recover() to r and check if it is nil or not
		fmt.Println("Recovered in cleanup", r)
	}
}

func say(s string){

	
    defer cleanup()
	for i:=0; i<3; i++{
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
		if i == 2 {
           panic("Oh dear a 2")
		}
	}
	
}

func main() {

	wg.Add(1)//Add a waitgroup before the go routine 
	go say("Hey")
	wg.Add(1)
	go say("There")
	wg.Wait()

	wg.Add(1)
	go say("Hi")
	wg.Wait()
}