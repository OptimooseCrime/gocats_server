package main

import (
	"fmt"
	"net/http"
)


func index_handler(w http.ResponseWriter, r *http.Request){
	//Fprint formats based on what you specified
	fmt.Fprint(w, "<h1>GO GO GO GO Rangers!</h1>")
	fmt.Fprint(w, "<p>GO HERE, Go There</p>")
	fmt.Fprint(w, "<p>You can go anywhere!</p>")
	fmt.Fprint(w, "<p>This is %s passing a %s</p>","me","variable")

}

func about_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Whoa dude.... what's this page all about", "about")
}


func main(){
	 var x = "Adam"
	fmt.Println("my name is %s", x)
	fmt.Println("Sup Earth, this is my server")
	//handlers: takes URL and take that path and what fuctiuon corresponds to path
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler) //the index
 //the index
	http.ListenAndServe(":8000", nil)//ListenAndServeTPS for https

}