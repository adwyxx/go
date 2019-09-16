package main

import "fmt"

func main(){
	//fmt.Println("Hello go!")
	//PrintInfo()
	VariableDefin1()
}
// 打印表格
func PrintInfo(){
	fmt.Println("姓名\t年龄\t籍贯\t住址\nJon\t20\tBeijing\t朝阳")
}

func VariableDefin1(){
	var name string
	var age int = 20
	name = "Tom"
	fmt.Println("Name : ",name)
	fmt.Println("Age : " ,age)

	address := "朝阳，北京"
	fmt.Println("Address : " ,address)


}

