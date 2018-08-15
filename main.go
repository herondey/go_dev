package main

import (
	"fmt"
	//↓vscodeだと警告がでるので、full pathでの指定が良い
	//"./animals"
	//"github.com/herondey/go_dev/animals"
	"math"
)

const ONE = 1

func one() (int, interface{}) {

	return ONE, nil
}

func main() {
	fmt.Printf("%d\n", math.MaxUint32)
	/*
		fmt.Println(AppName())
		fmt.Println(animals.MonkeyFeed())
		fmt.Println(animals.TigerFeed())
		fmt.Println(animals.RabbitFeed())
	*/

	x, y := one()
	fmt.Printf("%d %d\n", x, y)

	var a, b int
	a = 2
	b = 3

	if n := a * b; n > 10 {
		fmt.Println("10より大きい")
	} else {
		fmt.Println("10より小さい")
	}

	//fmt.Printf("%d\n", n)

	if _, err := one(); err != nil {
		fmt.Println("error !!")
	} else {
		fmt.Println("no error !!")
	}

}
