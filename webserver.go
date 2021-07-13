package main

import ("fmt" 
		"net/http")//short form for format

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa go is neat")
}
func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa about is neat")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.ListenAndServe(":8000",nil)
}