package main

import "fmt"
//全局匿名函数定义
var (
	Func1 = func(n1 int,n2 int) int {
			return n1+n2
	}
)

func main() {
	//函数只是用一次，可以在使用时直接定义
	f := func(n1 int,n2 int) int {
		return n1+n2
	}
	fmt.Println(f(1,2))

	//匿名函数声明后直接调用
	res := func(n1 int,n2 int) int {
		return n1 - n2
	}(10,2) //注意，直接在匿名函数结尾处传入参数列表进行调用
	fmt.Println(res)

	//调用全局匿名含数
	fmt.Println(Func1(1,2))
}