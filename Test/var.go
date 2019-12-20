package main

import (
	"fmt"
	"math"
)

func main() {
	// var name type 是声明单个变量的语法。
	var age int
	fmt.Println("my age is ", age)

	// var name type = initialvalue 的语法用于声明变量并初始化。
	var name = "test_one"
	fmt.Println("my name is ", name)

	// 自动推断具有初始值的变量的类型
	var big_age = 19
	fmt.Println("my big_age is ", big_age)

	var name_str = "test_two"
	fmt.Println("my chinse name is ", name_str)

	// 声明多个变量的语法是 var name1, name2 type = initialvalue1, initialvalue2。
	var width, height int = 60, 180
	fmt.Println("my width is", width, ",height is ", height)

	// 声明不同类型的变量
	var (
		l_name  = "test_three"
		l_age   = 20
		l_width = "60kg"
	)
	fmt.Println("my name is", l_name, ",age is", l_age, "width is", l_width)

	// 简短语法 name := initialvalue
	// 简短语法要求 := 操作符左边的所有变量都有初始值
	// 简短声明的语法要求 := 操作符的左边至少有一个变量是尚未声明的
	z_name, z_age := "test_four", 18 // 简短说明
	fmt.Println("my name is ", z_name, ",my age is", z_age)

	/** error */
	//err_name,err_age := "err_test" // error:assignment mismatch: 2 variable but 1 values
	//fmt.Println("my name is",err_name,",my age is",err_age)

	//a,b := 1,2
	//fmt.Println("a is", a, "b is", b)
	//a,b := 2,3 // error no new variables on left side of := , [因为 a 和 b 的变量已经声明过了，:= 的左边并没有尚未声明的变量。]

	// 变量也可以在运行时进行赋值
	test_a, test_b := 148.22, 248.78
	test_c := math.Max(test_a, test_b)
	fmt.Println("minimun values is", test_c)

	// 由于 Go 是强类型（Strongly Typed）语言，因此不允许某一类型的变量赋值为其他类型的值
	// error:cannot use "haha" (type string) as type int in assignment,是因为 age 本来声明为 int 类型，而我们却尝试给它赋字符串类型的值。
	//test_d := 10
	//test_d = "haha"

}
