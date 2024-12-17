package main

import "fmt"

func main() {
	var n1 int
	var s1 string
	n, err :=fmt.Scan(&n1 , &s1)
	 fmt.Println(n,err) 
	// //打印出来n是数字2，表示成功传入两个值，成功传入几个就是数字字
	// //nil，表示已经成功传入的数据值，如果有错误的情况就是err
	 //fmt.Println(n1, s1)
	 fmt.Printf("n1:%v,s1:%v",n1 , s1)
	// fmt.Scanln(&n1 , &s1)
	// fmt.Printf("n1:%v, s1:%v",n1, s1)
	//Scanln 这个函数输入多个值的时候，输入完一个不能按回车键否则第二个值无法输入，
	//正确的方式是输入空格然后回车
	//fmt.Scanf("n1: %v",&n1)
	//fmt.Printf("n1: %v",n1)

	
}

