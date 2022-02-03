package RecursionPackage

import "fmt"

func IndirectA(n int) {
	if n > 0 {
		fmt.Printf("indirect in function A : %d \n", n)
		IndirectB(n - 1)
	}
}

func IndirectB(n int) {
	if n > 1 {
		fmt.Printf("indirect in function B : %d \n", n)
		IndirectA(n / 2)
	}
}
