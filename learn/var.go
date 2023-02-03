package main

import "fmt"

// 全局变量
var (
	v9  = 100
	v10 = 101
)

func main() {
	// 局部变量
	//var age int
	//fmt.Println("int init value =", age)
	//age = 18
	//fmt.Println("age =", age)
	//intPrint()
	//StringPrint()
	//floatPrint()
}

func intPrint() {
	var age = 19
	fmt.Println("age =", age)
}

func StringPrint() {
	name := "男"
	fmt.Println(name)
}

func floatPrint() {
	var num1 float32 = 3.14
	num2 := -3.14
	num3 := 314e-2
	num4 := 314e+2

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)

	// 浮点类型默认为 float64
	var num5 = 256.000000916
	fmt.Printf("%T \n", num5)
	fmt.Printf("%v \n", num5)

	// float32 可能会有精度损失
	var num6 float32 = 256.000000916
	fmt.Printf("%T \n", num6)
	fmt.Printf("%v \n", num6)
}

// 没有char类型，一般用byte代替
func bytePrint() {
	var c1 byte = 'a'
	var c2 = 'b'
	c3 := 'c'
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)

}
