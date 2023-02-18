package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("This is to demonstrate command line args")
	fmt.Printf("Arguments passed through command line are as below:\n%v\n%v", args, args[1:])
	//we can pass args through command line while executing the code
	//Or we can add configuration file launch.json and add args in that json file
	//fmt.Println(len(args))
	
}
