package utils

import "fmt"

//1.先执行引入包的全局变量定义
var MaxValue int = 100
var MinValue int

//2.再执行引入包的init()方法
func init() {
		MinValue = 1
		fmt.Printf("utils.go init():MaxValue=%v ,MinValue=%v \n",MaxValue,MinValue)
}

func PrintValue() {
    fmt.Printf("utils.PrintValue() : MaxValue=%v ,MinValue=%v \n",MaxValue,MinValue)
}