package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r http.Request) {
	if r.URL.Path != "/form"{
		http.Error(w,"Not Found",http.StatusNotFound)
		return
	}
	if r.Method !="POST"{
		http.Error(w,"Unsupported Request Method",http.StatusMethodNotAllowed)
		return
	}
	if err:=r.ParseForm(); err!=nil{
		fmt.Fprintf(w, "ParseForm error: %v",err)
	}
	name,address:=r.FormValue("Name"),r.FormValue("Address")
	fmt.Fprintf(w,"Your Name: %v\n",name)
	fmt.Fprintf(w,"Your Address: %v",address)

}

func helloHandler(w http.ResponseWriter, r http.Request) {
	if r.URL.Path != "/hello"{
		http.Error(w,"Not Found",http.StatusNotFound)
		return
	}
	if r.Method !="PGET"{
		http.Error(w,"Unsupported Request Method",http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w,"hello")
	return
}


func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)
	fmt.Printf("starting at 8080\n")
	if err:= http.ListenAndServe("8080",nil); err!=nil{
		log.Fatal(err)
	}
}
