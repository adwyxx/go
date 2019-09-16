package main

import (
	"fmt"
	"strconv"
)

func main(){
	sum,plus := SumAndPlus(10,20)
	fmt.Printf("sum=%v, plus=%v \n",sum,plus)


	sum1,plus1 := SumAndPlus1(10,20)
	fmt.Printf("sum=%v, plus=%v \n",sum1,plus1)

		//函数赋值给一个变量
	fun1 := SumAndPlus
	sum1,plus1 = fun1(10,5)
	fmt.Printf("fun1 sum=%v, plus=%v \n",sum1,plus1)

	sum1,plus1 = getSumAndPlus(fun1,15,10)
	fmt.Printf("getSumAndPlus sum=%v, plus=%v \n",sum1,plus1)

	fmt.Println(myFunc("输出数字：",1,2,3,4,5))
}

//1.返回多个值
func SumAndPlus(num1 int ,num2 int) (int,int) {
		sum := num1 + num2
		plus := num1 - num2
		return sum,plus
}

//2. 也可以直接在返回值类型类表中直接指定返回值变量
func SumAndPlus1(num1 int,num2 int) (sum int, plus int) {
	//注意：此处sum、plus直接使用了返回值类型列表中的定义
	sum = num1 + num2		//只有赋值，不用声明
	plus = num1 - num2	//只有赋值，不用声明
		//注意：返回值得顺序一定得跟函数定义的返回值类型列表一直
		return sum,plus
}

//3.如果只有一个返回值的话，返回值列表的括号可以去掉
func Sum(num1 int,num2 int) int {
	return num1 + num2
}

//函数作为形参,注意：形参类型是函数的定义：func(int,int)(int,int)
func getSumAndPlus(sumAndPlus func(int,int)(int,int),num1 int ,num2 int) (int , int) {
    
    return sumAndPlus(num1,num2)
}

func myFunc(str string,args... int64) string {
	res := str
	for i := 0; i < len(args); i++ {
		res = res + strconv.FormatInt(args[i],10) + "\t"
	}

	for index, val := range args {
			fmt.Printf("args[%v]=%v ;",index,val)
	}
	return res
}