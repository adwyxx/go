package main

import (
	"fmt"
	"unsafe"
	"strconv"
)

func main(){

	//测试 int默认类型
	var n1 int = 100
	fmt.Printf("n1的类型为:%T ,n1占用的字节数:%d \n",n1,unsafe.Sizeof(n1))

	//var n2 int8 = 256

	// 将字符赋值给整型变量存储
	var char int='A'
	var char1 int = '中'
	//输出： char value is 65,type is int; char1 is 20013,type is int
	fmt.Printf("char value is %v,type is %T; char1 is %v,type is %T \n",char,char,char1,char1)
	fmt.Printf("char=%c, char1=%c \n",char,char1)
	//编译报错：constant 22269 overflows byte ，byte数据范围为0-255，超出了范围
	//var char3 byte = '国' 

	char2 := char+5
	//输出： char2=F , char2 value is 70
	fmt.Printf("char2=%c , char2 value is %v \n",char2,char2)

	//循环输出A-Z
	for i := 0; i < 26; i++ {
		newChar := char + i
		fmt.Printf("%c ",newChar)
	}
	
	var num1 float32 = -123.0000901
	var num2 float64 = -123.0000901
	//输出：num1= -123.00009 num2= -123.0000901 。num1的精度丢失了
	fmt.Println("num1=",num1,"num2=",num2)

	str := "hello"
	//str[0] = 'a' //cannot assign to str[0]
	fmt.Println(str)

	var num int32 = 100
	var f1 float32 = float32(num) // 低精度--转-->高精度
	var f2 int64 = int64(num) // 低精度--转-->高精度
	var b1 int8 = int8(num) // 高精度--转-->低精度
	fmt.Printf("num=%v f1=%f f2=%v b1=%v \n",num,f1,f2,b1)

	//高精度--转-->低精度 溢出情况
	var num4 int64 = 25542
	var b2 int8 = int8(num4) 
	fmt.Printf("num4=%v b2=%v \n",num4,b2) 
	//输出：num4=25542 b2=-58

	pi := 3.14159
	// 按数值本身的格式输出
	variant := fmt.Sprintf("%v %v %v", "月球基地", pi, true)
	fmt.Println(variant)

	
	// 匿名结构体声明, 并赋予初值
	profile := &struct {
		Name string
		HP   int
	}{
		Name: "rat",
		HP:   150,
	}
	fmt.Printf("使用'%%+v' %+v\n", profile)
	fmt.Printf("使用'%%#v' %#v\n", profile)
	fmt.Printf("使用'%%T' %T\n", profile)

	//---------------strcnv包--------------
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatBool(false))
	//参数：(浮点数值3.1415926,输出格式：浮点数据格式,保留小数位数,64位精度)
	fmt.Println(strconv.FormatFloat(3.1415926,'f',4,64))
	//参数：(浮点数值3.1415926,输出格式：科学技术发,保留小数位数,32位精度)
	fmt.Println(strconv.FormatFloat(3.1415926,'e',2,32))
	fmt.Println(strconv.FormatFloat(3.1415926,'E',3,64))
	fmt.Println(strconv.FormatFloat(3.1415926,'g',4,64))
	fmt.Println(strconv.FormatFloat(3.1415926,'G',6,64))

	fmt.Println("--------------strconv.ParseBool-----------------")
	bv1,_:=	strconv.ParseBool("true")
	fmt.Printf("bv1=%v , type is %T \n",bv1,bv1) //bv1=true , type is bool

	fmt.Println("--------------strconv.ParseFloat-----------------")
	fv1,_ := strconv.ParseFloat("3.1415926",64);
	fmt.Printf("fv1=%f , type is %T \n",fv1,fv1) //v1=3.141593 , type is float64

	fv2,err := strconv.ParseFloat("3.14159263AAA",64);
	fmt.Printf("fv2=%f , type is %T ; ERR:%v \n",fv2,fv2,err)
	//fv2=0.000000 , type is float64 ; ERR:strconv.ParseFloat: parsing "3.14159263AAA": invalid syntax

	fmt.Println("--------------strconv.ParseInt-----------------")
	//第二个参数：指定进制（2到36），如果base为0，则会从字符串前置判断，"0x"是16进制，"0"是8进制，否则是10进制；
	iv1,err := strconv.ParseInt("01101",2,64) // 二进制字符串转Int64
	fmt.Printf("iv1=%v , type is %T ; ERR:%v \n",iv1,iv1,err) //iv1=13 , type is int64 ; ERR:<nil>

	iv1,err = strconv.ParseInt("157",8,64) // 八进制字符串转Int64
	fmt.Printf("iv1=%v , type is %T ; ERR:%v \n",iv1,iv1,err) //iv1=111 , type is int64 ; ERR:<nil>

	iv1,err = strconv.ParseInt("157",10,64) // 十进制字符串转Int64
	fmt.Printf("iv1=%v , type is %T ; ERR:%v \n",iv1,iv1,err) //iv1=157 , type is int64 ; ERR:<nil>

	iv1,err = strconv.ParseInt("1578A",16,64) // 十六进制字符串转Int64，注意:不要加0x
	fmt.Printf("iv1=%v , type is %T ; ERR:%v \n",iv1,iv1,err) //iv1=5496 , type is int64 ; ERR:<nil>

	iv1,err = strconv.ParseInt("128",10,8) // 十进制字符串转Int64，精度为int8（-128~127）,此时报错out of range
	fmt.Printf("iv1=%v , type is %T ; ERR:%v \n",iv1,iv1,err) //iv1=7 , type is int64 ; ERR:strconv.ParseInt: parsing "157": value out of range

	fmt.Println("--------------strconv.ParseUint-----------------")
	uiv1,err := strconv.ParseUint("111",2,8)
	fmt.Printf("uiv1=%v , type is %T ; ERR:%v \n",uiv1,uiv1,err)
	//uiv1=7 , type is uint64 ; ERR:<nil>

	uiv1,err = strconv.ParseUint("127",8,8)
	fmt.Printf("uiv1=%v , type is %T ; ERR:%v \n",uiv1,uiv1,err)
	//uiv1=87 , type is uint64 ; ERR:<nil>

	uiv1,err = strconv.ParseUint("128",10,8)
	fmt.Printf("uiv1=%v , type is %T ; ERR:%v \n",uiv1,uiv1,err)
	//uiv1=128 , type is uint64 ; ERR:<nil>

	uiv1,err = strconv.ParseUint("128",16,16)
	fmt.Printf("uiv1=%v , type is %T ; ERR:%v \n",uiv1,uiv1,err)
	//uiv1=296 , type is uint64 ; ERR:<nil>

	uiv1,err = strconv.ParseUint("128",32,32)
	fmt.Printf("uiv1=%v , type is %T ; ERR:%v \n",uiv1,uiv1,err)
	//uiv1=1096 , type is uint64 ; ERR:<nil>

	uiv1,err = strconv.ParseUint("0x128",0,64)
	fmt.Printf("uiv1=%v , type is %T ; ERR:%v \n",uiv1,uiv1,err) 
	//uiv1=296 , type is uint64 ; ERR:<nil>

}