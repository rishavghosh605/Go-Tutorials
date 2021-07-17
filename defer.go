//Used when wg.Done() cannot be reached by the function, one example would be that the function  hits an error and the code flow leaves the function before executing wg.Done()- thus to handle go routines at that point we need the defer statement

package main

import ("time"
        "fmt"
	    "sync")

var wg sync.WaitGroup

func foo() {
	defer fmt.Println("Done") // This is going to be evaluated immediately but it will run after foo function is finished
	defer fmt.Println("Are we Done?")// Defer statement works in a LIFO manner
	fmt.Println("Doing stuff who knows what")
}

func say(s string){

	defer wg.Done()

	for i:=0; i<3; i++{
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
	}
	
}

func main() {
    foo()

	wg.Add(1)//Add a waitgroup before the go routine 
	go say("Hey")
	wg.Add(1)
	go say("There")
	wg.Wait()

	wg.Add(1)
	go say("Hi")
	wg.Wait()
}
