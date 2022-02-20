package main

import (
	"fmt"
)

// typeで構造体の型を指定する。
// typeを指定することで同じような構造体を複数宣言することが簡単になる。
type Mydata struct {
	Name string
	Data []int
}

func main() {
	taro := Mydata{"Taro", []int{10, 20, 30}}
	hanako := Mydata{
		Name: "Hanako",
		Data: []int{90, 80, 70},
	}
	fmt.Println(taro)
	fmt.Println(hanako)
}
