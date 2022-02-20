package main

import (
	"fmt"
)

type Mydata struct {
	Name string
	Data []int
}

// Mydataにメソッド(構造体に組み込む関数のこと)を追加する
func (md Mydata) PrintData() {
	fmt.Println("*** Mydata ***")
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
	fmt.Println("*** end ***")
}

func main() {
	taro := Mydata{
		"Hanako", []int{98, 76, 54, 32, 10},
	}
	taro.PrintData()
}
