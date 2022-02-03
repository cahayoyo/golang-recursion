package RecursionPackage

import "fmt"

func NestedRecursion(n int) int {
	if n > 10 {
		fmt.Printf("Recursion in inner function : %d\n", n)
		return n - 4
	} else {
		fmt.Printf("Recursion in outer function : %d\n", n)
		return NestedRecursion(NestedRecursion(n + 5))
	}
}
