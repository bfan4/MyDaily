package main

import (
	//"github.com/astaxie/beego"
	"io"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world, this is version 1.0")

}

func homePage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is Home Page! Welcome!!")
}

func main() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/hello", sayHello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
