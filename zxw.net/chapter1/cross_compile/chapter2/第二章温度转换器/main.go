package main
import "fmt"
import "time"
func main (){
	// var c float64
	// fmt.Print("请输入一个温度：")
    // fmt.Scanln(&c)
	// k := c+273.15
	// f := float64(5/9)*c+float64(320)
	// fmt.Print("正在转换.....")
	// time.Sleep(3*time.Second)
	// fmt.Printf("转换完成\n开氏度: %.2f\n摄氏度: %.2f\n",k,f)
	// time.Sleep(30*time.Second)












	// var c float64
	// fmt.Print("请输入一个温度: ")
	// fmt.Scanln(&c)
	// k:= c+ 273.15
	// f:=float64(5/9)*c+float64(320)
	// fmt.Print("温度转换中....")
	// time.Sleep(3*time.Second)
	// fmt.Printf("转换完成\n开氏度: %.2f\n华氏度: %.2f\n",k,f)















	var c  float64
	fmt.Print("请输入一个温度： ")
	fmt.Scanln(&c)
	K := c +273.15
	F := float64(5/9)*float64(c)+320
	fmt.Println("正在转换中请稍等.....")
	time.Sleep(5*time.Second)
	fmt.Printf("转换已完成\n 开氏度：%.2f\n 摄氏度： %.2f\n",K,F)








}