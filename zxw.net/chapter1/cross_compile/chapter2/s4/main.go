package main

import "fmt"

func main() {
	//fmt.Println(1 == 1) 打印出true类型
	//fmt.Println(1 == 2) 打印出false类型
	//fmt.Printf("%T\n", 1==1)  //打印出true类型 并且换行
	// const(
		q1 := 1==1 
	 	q2 := 1==2
	// ) 常量赋值
	 fmt.Println(q1 , q2 ) //打印出来这个两个常量的类型
	//fmt.Printf(q1+q2) 布尔值不能参与数字运算
	//也不能和其他类型转换

	//a1 :="世界那么大\n我想去看看"
	
	//a2 :=`钱包那么小,那也去不了`
	//fmt.Println(a1, a2)
}
