package main

import ("fmt"
	"net/http")

func index_handler (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Hey there</h1>")
}

func main(){
	http.HandlerFunc("/", index_handler)
	http.ListenAndServe(":8000",nil)
}