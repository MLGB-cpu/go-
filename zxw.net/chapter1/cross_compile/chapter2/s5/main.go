package main
import "fmt"
func main(){
	n1 := 663 //int 类型
	n2 :=3.66 //float64 类型
	 //total := n1 + n2
	
	//T()
	//第一种打印方式 int-> float64
	//total := float64(n1) +n2 666.66
	
	//第二种 ：float64-> int
	total := int(n2) +n1 //666
	fmt.Println(int(n2)) // 3 
	//将浮点数转为整数的时候会将后边的小数舍弃
	fmt.Println(total)

}