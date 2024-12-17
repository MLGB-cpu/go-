package main
import "fmt"
import "time"
func main (){
	fmt.Println(1234)
	fmt.Println(5678)
	time.Sleep(5* time.Second)
	now := time.Now ()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println( now.Hour ())


}