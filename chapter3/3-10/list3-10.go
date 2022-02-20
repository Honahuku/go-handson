package main

import (
	"fmt"
)

type Mydata struct {
	Name string
	Data []int
}

func main() {
	taro := new(Mydata) //newで構造体を作成する。構造体はゼロ値になり、返り値はポインタになる。
	fmt.Println(taro)
	taro.Name = "Taro"
	taro.Data = make([]int, 5)
	fmt.Println(taro)
}
