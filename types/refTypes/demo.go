package main

import (
	"fmt"
	_ "unsafe"
	_ "strconv"
)

func main(){
	i := 10
	//var pv *int = &i //在变量前加“&”来获取其内存地址
	pv := &i
	fmt.Printf("pv type is %T \n",pv) //pv type is *int
	fmt.Printf("pv=$%v &pv=%v,i=%v,&i=%v \n",pv,&pv,i,&i)

	fmt.Println("变量i的内存地址：",&i)				//变量i的内存地址： 0xc0000120a0
	fmt.Println("指针pv指向的内存地址：",pv)	//指针pv指向的内存地址： 0xc0000120a0
	fmt.Println("指针pv本身的内存地址：",&pv) //指针pv本身的内存地址： 0xc000006028
	fmt.Println("指针pv指向的值：",*pv) //指针pv指向的值：10

}