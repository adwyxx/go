package main

import "fmt"

func main ()  {
	var age uint 
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)

	if age < 18 {
    fmt.Println("少年")
	} else if age>18 && age<40 {
			fmt.Println("青年")
	} else if age>40 && age<50 {
			fmt.Println("中年")
	} else {
			fmt.Println("老年")
	}

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

	//判断运年
	var year uint
	fmt.Println("请输入年份")
	fmt.Scanln(&year)
	//被4整除但不能被100整除，或者 能被400整除
	if (year % 4 == 0 && year % 100 !=0) || year % 400 == 0 {
			fmt.Printf("%v 是闰年 \n",year)
			
	} else {
			fmt.Printf("%v 不是闰年 \n",year)
			fmt.Scanln(&year)
	}
}

