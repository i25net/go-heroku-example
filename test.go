package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", indexpage)
	http.HandleFunc("/test", hello)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, time.Now())
}

func indexpage(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello heroku")
}