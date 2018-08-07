package main

import (
	"fmt"
	//vscodeだと警告がでる
	//"./animals"
	"github.com/herondey/go_dev/animals"
)

func main() {
	fmt.Println("hello")
	fmt.Println(AppName())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.TigerFeed())
	fmt.Println(animals.RabbitFeed())
}
