package RecursionPackage

import "fmt"

func TailRecursion(n int) {
	if n > 0 {
		fmt.Println(n)
		TailRecursion(n - 1)
	}
}
