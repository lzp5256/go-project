package main

import (
	"fmt"
	"unsafe"
)

/**
	 Go 支持的基本类型:
   	 bool
	 数字类型:
		 int8, int16, int32, int64, int
		 uint8, uint16, uint32, uint64, uint
		 float32, float64
		 complex64, complex128
		 byte
		 rune
	string
*/

func main() {
	// bool
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)

	// 有符号整型
	var int_a int = 89
	int_b := 95
	fmt.Println("value of a is ", int_a, ", int_b is ", int_b)
	fmt.Printf("type of int_a is %T , size of int_a is %d", int_a, unsafe.Sizeof(int_a))   // a的类型和大小
	fmt.Printf("\ntype of int_b is %T , size of int_b is %d", int_b, unsafe.Sizeof(int_b)) // b的类型和大小
	fmt.Printf("\n")

	// 浮点数
	float_a, float_b := 5.67, 5.78
	fmt.Printf("type of float_a %T float_b %T\n", float_a, float_b)
	float_sum := float_a + float_b
	diff := float_a - float_b
	fmt.Println("sum:", float_sum, "diff:", diff)
	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)

	// string
	first := "Naveen"
	last := "Ramanathan"
	z := first + " " + last
	fmt.Println("My name is", z)

	// 类型转换
	i := 55
	j := 55.5
	// sum := i + j  // invalid operation: i + j (mismatched types int and float64)
	sum := i + int(j)
	fmt.Println("sum", sum)

}
