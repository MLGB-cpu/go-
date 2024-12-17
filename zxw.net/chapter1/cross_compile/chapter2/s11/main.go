package main

import ("fmt" 
       //"strings"
	)

 

func main(){
	//s1 :="hello world"
	
	//fmt.Println(strings.Contains(s1,"llo"))//判断一个子串是否在一个字符当中
	//存在就等会得到一个true，相反就是false
       
	//s1 := "hello world" 
          //012345678910 索引值
		  //Index返回是第一次出现子串的位置或者索引值
	// fmt.Println(strings.Index(s1,"llo"))
	// fmt.Println(strings.Index(s1,"l"))
	// //得出位置为2,
	// fmt.Println(strings.LastIndex(s1,"l"))
	// //LastIndex 是从字符右边向左搜索的。
	
	
	//fmt.Println(strings.Trim(s1,"hd")) //得出的结果是" ello worl"
	//fmt.Println(strings.Trim(s1,"hd "))//得出的结果是"ello worl"
	//Trim 这个函数只能把首尾的字符移除，但是中间的字符无法移除
	//fmt.Println(strings.Trim(s1,"he "))//得出的结果是"llo world"
    //fmt.Println(strings.TrimLeft(s1,"he"))
	//TrimLeft这个函数是移除左边的字符。
	fmt.Println(strings.TrimRight(s1,"ld"))
    //TrimRight这个函数是移除右边的字符。

}
	
