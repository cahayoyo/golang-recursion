package main

import (
	"fmt"
	"golang_recursion/RecursionPackage"
)

func main() {
	var number int
	fmt.Print("Enter Recursion : ")
	fmt.Scan(&number)

	fmt.Println("This is Head Recursion")
	RecursionPackage.HeadRecursion(number)

	fmt.Println("This is Tail Recursion")
	RecursionPackage.TailRecursion(number)
}