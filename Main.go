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

	fmt.Println("\nThis is Tail Recursion")
	RecursionPackage.TailRecursion(number)

	fmt.Println("\nThis is Tree Recursion")
	RecursionPackage.TreeRecursion(number)

	fmt.Println("\nThis is Indirect Recursion")
	RecursionPackage.IndirectA(number)

	fmt.Println("\nThis is Nested Recursion")
	fmt.Print(RecursionPackage.NestedRecursion(number))
	fmt.Print(" <- Result Return")
}
