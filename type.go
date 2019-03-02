package main

import "fmt"

type Ptr *int
type Map map[int]string
type MapMap Map

func main() {
	var p *int
	var mm Map
	var mmm MapMap
	var m1 map[int]string = mm
	var m2 map[int]string = mmm
	var ptr Ptr = p
	fmt.Println(ptr)
	fmt.Println(m1)
	fmt.Println(m2)

	fmt.Println(1)
	println(2)
}