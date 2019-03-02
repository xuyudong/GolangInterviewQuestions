package main

import "fmt"

func main()  {
	fmt.Println(fibonacci(10))
}

func fibonacci(n int) int {
	if (n==1||n==2) {
		return 1
	}
	return fibonacci(n-1)+fibonacci(n-2)
}
