package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("main.go")
	if err!=nil {
		// 文件不存在、权限等原因
		fmt.Println("open file failed reason:" + err.Error())
		return
	}
	defer f.Close()
	
}
