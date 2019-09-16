package main

import "fmt"

func main ()  {
	n1 := 5
	n2 := 20
	switch n1 {
			case n2,10,15 :
					fmt.Println("Ok")
			case 5 :
					fmt.Println("Ok")
			default : 
					fmt.Println("Not Ok")
	}

	var age uint = 35
	//switch后表达式也可不写，但case后的表达式返回值必须是bool型
	//等同于if-else
	switch {
		case age < 18:
			fmt.Println("少年")
		case age>18 && age<40 :
			fmt.Println("青年")
		case  age>40 && age<50 :
			fmt.Println("中年")
		default :
			fmt.Println("老年")
	}

	//switch后可直接声明变量并赋值 不推荐这么写
	//注意d := 1;后面有个分号
	switch d := 1; { 
		case d==1,d==2,d==3,d==4,d==5 : //注意case后必须是表达式
			fmt.Println("工作日")
		default :
			fmt.Println("休息日")
	}

	//判断运年
	var year uint=1998
	switch {
		case IsLeapYear(year) : //表达式可以是一个函数
			fmt.Println("闰年")
		default :
		fmt.Println("非闰年")
	}

	//switch穿透--fallthrough
	num := 10
	switch num {
			case 10 :
					fmt.Println("OK 10")
					fallthrough
			case 20 :
					fmt.Println("OK 20")
			case 30 :
					fmt.Println("OK 30")
			default :
					fmt.Println("NOT OK")
	}


	arr := [5]int{1,15,20,11,12}
	for index,val := range arr {
			fmt.Printf("索引值：%v ,value: %v \n",index,val)
	}
}

//判断是否是闰年
func IsLeapYear(year uint) bool {
	if (year % 4 == 0 && year % 100 !=0) || year % 400 == 0 {
		return true
	} else {
		return false
	}
}