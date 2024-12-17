package main

import (
	"fmt"
	"strings"
	//"strings"
	//"unicode/utf8"
)

func main() {
// s1:= ("hello world") //11个字节  一个字符占一个字节
// s2:=("hello 世界")   // 12个字节 中文一个文字占3个字节所以是12个字节
// 	fmt.Println(len(s1))

// 	fmt.Println(len(s2))
// 	fmt.Println(utf8.RuneCountInString(s2))
	//"utf8.RuneCountInString"这个函数是统计字符的个数。
	//字符是指（）内的内容。
    
	       //拼接字符
	// s1 :="hello"
	// s2 :="world"
	// fmt.Println(s1 +  s2) 
	// //方式一:直接使用逻辑运算中的加法打印，打印出"helloworld"
	// fmt.Println(s1 +" "+ s2)
	// //在字符中间加上"  " 这个符号打印结果是"hello world"
	// fmt.Printf("%s %s\n",s1 ,s2)
	// //方式二：打印两个字符换行，后跟上两个变量名称 s1 , s2
	
	
	
	// s1 := "heLLo"
	// fmt.Println(strings.ToUpper(s1))
	
	// //"ToUpper"这个函数是改变字符的大小写得出结果是HELLO
    
	// fmt.Println(strings.Title(s1))
	
	// //"title"这个函数是将目标字符开头改成大写得出结果是Hello
	
	// fmt.Println(strings.ToLower(s1))
    
	// //"ToLower" 这个函数是将目标字符更改为小写（字符里包含大小写不收影响）
    
	s1 := "hello world"
	fmt.Println(strings.HasPrefix(s1 , "hello"))
	//HasPrefix这个函数是用来判断字符串的开头 
	//正确的情况系会返回一个布尔值true，反之会得到一个false
    
	fmt.Println(strings.HasSuffix(s1 , "world"))
	
	//HasSUffix这个函数是用来判断字符串的结尾 
	//正确的情况系会返回一个布尔值true，反之会得到一个false
	//注意使用这个函数需要导入一个包
	//"strings" fmt.Println(strings.HasPrefis(一个字符串+开头也可以叫子字符串))
	
    }