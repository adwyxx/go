package main

import (
	"fmt"
	"demo/init/utils"
)

//3.其次，执行main.go的全局变量定义
var num1 int= 10
var num2 int

//4.再次，执行main.go的init()函数
func init() {
	num2=20
	fmt.Printf("mian.go init():num1=%v,num2=%v \n",num1,num2)
}

//5.最后，执行main()函数
func main() {
	num1 = 30
	num2 = 40
	fmt.Printf("mian.go main():num1=%v,num2=%v \n",num1,num2)
	utils.PrintValue()
}