package main

import "fmt"

func main() {
	// 金額
	c := 1314
	// 貨幣の種別
	cc := []int{100, 50, 10, 1}

	// 大きい貨幣から順に両替していく
	for _, v := range cc {
		var q int = c % v
		fmt.Println(v, "円が", q, "枚")
	}
}
