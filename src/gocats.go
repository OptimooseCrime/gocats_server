package main

import (
	"fmt"
	"net/http"
)


func index_handler(w http.ResponseWriter, r *http.Request){
	//Fprint formats based on what you specified
	fmt.Fprint(w, "GO GO GO GO Rangers!")
}

func about_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Whoa dude.... what's this about page all about")
}


func main(){
	fmt.Println("Sup Earth, this is my server")
	//handlers: takes URL and take that path and what fuctiuon corresponds to path
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler) //the index
 //the index
	http.ListenAndServe(":8000", nil)

}