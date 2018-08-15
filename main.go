package main

import (
	"fmt"
	//vscodeだと警告がでる
	//"./animals"
	"github.com/herondey/go_dev/animals"
	"math"
)

func main() {
	fmt.Println("hello")
	fmt.Println(AppName())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.TigerFeed())
	fmt.Println(animals.RabbitFeed())
	fmt.Printf("%d\n", math.MaxUint32)
}
