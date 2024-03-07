package main

import (
	"fmt"
	"log"
	"net/http"
	"golang-react-todo/router"
)

func main(){
	r := router.Router();

	fmt.Println("Starting server at port 9000")
	
	log.Fatal(http.ListenAndServe(":9000",r))
}