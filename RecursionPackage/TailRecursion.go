package RecursionPackage

import "fmt"

func TailRecursion(n int) {
	if n > 0 {
		fmt.Printf("%d ", n)
		TailRecursion(n - 1)
	}
}
