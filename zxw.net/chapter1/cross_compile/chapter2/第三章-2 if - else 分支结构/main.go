package main

import "fmt"

func main(){
// 	a :=85 
// 	if a > 5 {
// 		//当表达式大于5 为真true的的时候就会执行这段代码
// 		fmt.Println("a > 5")

// }else {//当表达式小于5 为假 false的的时候就会执行这段代码}
// 	fmt.Println("a<=5")

// }

score :=80  //以这个值为参考标准去判断，满足这个值会打印出来
if score >=90 {   // 已设定的值去判断如果达到条件，就会打印出来相应的结果。
fmt.Println("等级为: A")//不满足的情况下会继续判断查找判断，如果最后还是没有找到判断
} else if score >=80 { 
fmt.Println("等级为: B")
} else if score >= 70 {
fmt.Println("等级为: C")
} else if score >= 60{
fmt.Println("等级为: D")
} else {               //如果最后还是没有找到判断
fmt.Println("等级为: E")
} 

}








}