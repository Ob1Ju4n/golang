package golang

import "fmt"

// PrintFibonacciNthTerm - Prints the fibonacci sequence to the specified
// nth term
func PrintFibonacciNthTerm(n int) {
	f := fibonacci()
	for i := 0; i < n; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	n, n1, n2 := 0, 0, 1
	return func() int {
		f := n
		if n > 1 {
			f = n1 + n2
			n1 = n2
			n2 = f
		}
		n++

		return f
	}
}
