package main

import (
	"fmt"
	//↓vscodeだと警告がでるので、full pathでの指定が良い
	//"./animals"
	//"github.com/herondey/go_dev/animals"
	_ "math"
	"runtime"
)

func subLoop(){
	defer func(){
		fmt.Println(" sub loop end")
	}()

	for i:=0;i<20;i++{
		fmt.Println(" sub loop", i)
	}
}

func main() {

	/*
		fmt.Println(AppName())
		fmt.Println(animals.MonkeyFeed())
		fmt.Println(animals.TigerFeed())
		fmt.Println(animals.RabbitFeed())
	*/

	go subLoop()

	for i := 0; i < 20; i++ {
		fmt.Println("main loop", i)
	}

	fmt.Printf("Num CPU %d\n", runtime.NumCPU())
	fmt.Printf("Num Goroutine %d\n", runtime.NumGoroutine())
	fmt.Printf("Version %d\n", runtime.Version())
}
