//Here we solve the issue of the program ending before the goroutine itself using synchoronization
package main

import ("fmt"
        "time"
	    "sync")

var wg sync.WaitGroup

func say(s string){
	for i:=0; i<3; i++{
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
	}
	wg.Done()//We use this to notify the wg that the function is done thus no need to wait
}

func main(){
	wg.Add(1)//Add a waitgroup before the go routine 
	go say("Hey")
	wg.Add(1)
	go say("There")
	wg.Wait()

	wg.Add(1)
	go say("Hi")
	wg.Wait()
}