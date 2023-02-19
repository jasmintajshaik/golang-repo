package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

/*
This code requires a test server or any simple application that should be up and running on some port.
We first give the URL that app is running on. In this case I have given http:localhost:8080
I have configured a simple web app that runs on port 8080 and passing its URL through command line args while executing.
This code then reads the data of that app through http API call (GET)
*/
func main() {
	args := os.Args
	//Conditional to check whether url is given or not
	if len(args) < 2 {
		fmt.Printf("Please enter 1 or more Arguments\n")
		os.Exit(1)
	}
	fmt.Printf("First argument is:%s\n", args[1])
	//checking whether URL format is correct or not
	_, err := url.ParseRequestURI(args[1])
	if err != nil {
		fmt.Printf("Invalid URL format entered\n%s\n", err)
		os.Exit(1)
	}
	//checking the response code
	response, err := http.Get(args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("HTTP Status code:%d\n", response.StatusCode)
	defer response.Body.Close()            //closing the stream
	body, err := io.ReadAll(response.Body) //reading and printing body
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Body of the response:%s", body)
}
