// package main

// import ("fmt")
// func main(){
// 	var(name string
// 	num int
//      price float64)

// 	 fmt.Print("名称： ")
// 	 fmt.Scanln(&name)
// 	 fmt.Print("数量： ")
// 	 fmt.Scanln(&num)
// 	 fmt.Print("单价:  ")
// 	 fmt.Scanln(&price)
// 	 total:=float64(num)*price
// 	 fmt.Println("总计： ",total)

// 	 fmt.Println("结算明细")
// 	 fmt.Println("名称： ",name)
// 	 fmt.Println("数量： ",num)
// 	 fmt.Println("总计： ",total)

//     //ret:=`
// 	结算清单
// 	名称： %v
// 	数量： %v
// 	单价： %v
// 	总价： %v
// 	`
// 	fmt.Printf(ret, name, num, price, total)
// }

// package main

// import "fmt"

// func main() {
//     // 定义商品信息
//     var productName string = "玩具"
//     var quantity int = 200
//     var unitPrice float64 = 10

//     // 计算总价
//     totalPrice := float64(quantity) * unitPrice

//     // 输出结果
//     fmt.Printf("商品：%s\n", productName)
//     fmt.Printf("数量：%d\n", quantity)
//     fmt.Printf("单价：%.2f\n", unitPrice)
//     fmt.Printf("总价：%.2f\n", totalPrice)
// }

package main

import "fmt"

func main() {
	var (
		name  string
		num   int
		price float64
	)
	// fmt.Print("名称: ")
	// fmt.Scanln(&name)
	// fmt.Print("数量: ")
	// fmt.Scanln(&num)
	// fmt.Print("价格: ")
	// fmt.Scanln(&price)
	// total :=float64(num)*price
	// fmt.Println("总价: ",total)

	// // fmt.Println("商品清单")
	// // fmt.Println("名称： ",name)
	// // fmt.Println("数量",num)
	// // fmt.Println("价格: ",price)
	// // fmt.Println("总价: ",total)
	// ret := `
	// 商品清单
	// 名称: %v
	// 数量：%v
	// 价格：%v
	// 总价：%v
	// `
	// fmt.Printf(ret, name, num, price, total)
	fmt.Print("名称： ")
	fmt.Scanln(&name)
	fmt.Print("数量： ")
	fmt.Scanln(&num)
	fmt.Print("价格： ")
	fmt.Scanln(&price)
	total := float64(num) * price
	fmt.Println("总价： ", total)

	// fmt.Println("商品清单")
	// fmt.Println("名称： ",name)
	// fmt.Println("数量： ",num)
	// fmt.Println("价格：",price)
	// fmt.Println("总价： ",total)
	ret := `
	结算清单 
	名称： %v
	数量： %v
	价格： %v
	总价： %v
	`
	fmt.Printf(ret, name, num, price, total)

}
