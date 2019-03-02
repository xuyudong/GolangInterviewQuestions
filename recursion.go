package main

import "fmt"

func main(){
	f(1)
}
func f(n int)  {
	fmt.Printf("Level %d : n location %p \n",n, &n)
	if n<4 {
		f(n+1)
	}
	//1-2-3-4
	fmt.Printf("Level_ %d : n location %p \n",n, &n)

}