package RecursionPackage

import "fmt"

func TreeRecursion(n int) {
	if n > 0 {
		fmt.Printf("%d ", n)
		TreeRecursion(n - 1)
		TreeRecursion(n - 1)
	}
}
