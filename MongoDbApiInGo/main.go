package main;

import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/mahendrarathore1742/router"
) 
 

func main(){

	fmt.Println("Bulding API using mongoDB");
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":8000", r))

}






















