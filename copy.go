package main

import "fmt"

func main()  {
	//a := []int{1, 2}
	//fmt.Printf("%p\n",a)
	//b := []int{3, 4}
	//fmt.Printf("%p\n",b)
	//
	//check := a
	//copy(a, b)
	//fmt.Printf("%p\n",a)
	//fmt.Printf("%p\n",check)
	//fmt.Printf("%p\n",b)
	// Output: [3 4] [3 4] [3 4]


	a := []int{1, 2}
	fmt.Printf("1:%p\n",a)
	b := []int{3, 4}
	fmt.Printf("2:%p\n",b)
	check := a
	a = b
	fmt.Printf("3:%p\n",a)
	fmt.Printf("4:%p\n",check)
	fmt.Printf("5:%p\n",b)
	fmt.Println(a, b, check)
	// Output: [3 4] [3 4] [1 2]
}