package main

import (
	"fmt"
)

type Mydata struct {
	Name string
	Data []int
}

func main() {
	taro := Mydata{
		"Taro",
		[]int{10, 20, 30},
	}
	fmt.Println(taro)
	rev(&taro) //構造体へ値を参照渡しする。構造体はコピーされない。
	fmt.Println(taro)
}

func rev(md *Mydata) Mydata {
	od := (*md).Data
	nd := []int{}
	for i := len(od) - 1; i >= 0; i-- {
		nd = append(nd, od[i])
	}
	md.Data = nd
	return *md
}
