package main

import "fmt"

func main() {
    //    a:=10==10 //true a=true
	//    b:=10!=10 //false b=false
	//    fmt.Printf("a:%v,type:%T\n",a,a)
	//    fmt.Printf("b:%v,type:%T\n",b,b)
	   a:=10==10 //true a=true
	   b:=10!=10 //false b=false
	//    fmt.Println(a && a) && and这个字符两数值都为真的或者true则得出的结论为 true
	//    fmt.Println(a && b) 其中一边的数值类型为假或者数值类型为false，则得出的结论为false

	//      fmt.Println(a || b) or 逻辑或的这个字符当两边的数值类型其中一个为真(true)
	//得出的结论则为true
    //     fmt.Println(b || b) || or 逻辑或的这个字符当两边的数值类型其中一个为真(true)
	//得出的结论则为true
	 fmt.Println(!a) //！这个运算符是取反的意思，当原始数值为true 
	 //得出的结果则是false 反之是得到true
     fmt.Println(!b)

	 //赋值运算符，+= -= *= /= %=
	 

}
	