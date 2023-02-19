package main

import "fmt"

type Employee struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	fmt.Println("This is a Struct program")
	emp1 := Employee{
		Name:   "Jasmin",
		Age:    23,
		Salary: 23000,
	}
	fmt.Printf("%v",emp1)
}
