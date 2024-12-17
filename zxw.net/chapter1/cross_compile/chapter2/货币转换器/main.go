package main

import "fmt"
import "time"

func main() {
	var rmb float64
	fmt.Print("请输入兑换金额: ")
	fmt.Scanln(&rmb)
	sgd := rmb / 5.8
	dollar := rmb / 7.59
	php := rmb * 8.4
	fmt.Println("正在兑换.....")
	time.Sleep(5*time.Second)
    fmt.Printf("兑换已完成\n sgd: %.2f\n dollar:%.2f\n php: %.2f", sgd, dollar, php)



}
