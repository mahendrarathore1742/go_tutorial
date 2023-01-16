package main;

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
);

func main(){
	fmt.Println("hello mod from golang");
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome)
	log.Fatal(http.ListenAndServe(":8000",r))
}


func serverHome( w http.ResponseWriter,r *http.Request){

	 w.Write([]byte("<h1> hello from home </h1>"));
	 
}