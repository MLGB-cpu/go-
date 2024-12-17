package main
import "fmt"
import "time"
func main(){

	// fmt.Println(123)
	// //让程序暂停一会睡个5秒钟
	// time.Sleep(5 * time.Second)
	// fmt.Println(456)
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())

	fmt.Println(now.Hour())




}