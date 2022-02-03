package RecursionPackage

import "fmt"

func HeadRecursion(n int) {
	if n > 0 {
		HeadRecursion(n - 1)
		fmt.Printf("%d ", n)
	}
}
