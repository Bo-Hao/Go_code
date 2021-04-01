package main

import (
	"fmt"
	"unsafe"
)

/*
The example of unsafe pointer:

pointer, uintptr, unsafe.Pointer
2.1 *Pointer
普通指针，用于传递对象地址，不能进行指针运算。

2.2 unsafe.Pointer
通用指针类型，用于转换不同类型的指针，不能进行指针运算。

2.3 uintptr
用于指针运算，GC 不把 uintptr 当指针，uintptr 无法持有对象。uintptr 类型的目标会被回收

2.4 结论
unsafe.Pointer 可以和 普通指针 进行相互转换。

unsafe.Pointer 可以和 uintptr 进行相互转换。

也就是说 unsafe.Pointer 是桥梁，可以让任意类型的指针实现相互转换，也可以将任意类型的指针转换为 uintptr 进行指针运算
*/

func main() {
	//define a sequence of int8 with length 3.
	a := [3]int8{6, 8, 9}

	//取出数组第一个位置的地址
	a_first_point := &a[0]
	a_first_unsafe_point := unsafe.Pointer(a_first_point)
	fmt.Println("a[0]的地址为：", a_first_unsafe_point, a_first_point)

	//指针只能一个字节字节取，int8占一个字节，所以看到值只加了1
	fmt.Println("a[1]的地址为：", unsafe.Pointer(&a[1]))

	//把a_first_unsafe_point转成uintptr类型，就可以指针运算了
	a_uintptr_first_unsafe_point := uintptr(a_first_unsafe_point)
	fmt.Println("unitptr:", a_uintptr_first_unsafe_point)

	//指针+1 表示到了数组的第二个位置
	a_uintptr_first_unsafe_point++
	fmt.Println("a[0]位置指针自增1后，的指针位置：", a_uintptr_first_unsafe_point)

	//打印出来可以看到跟&a[1]的地址是一样的
	a_uintptr_second_unsafe_point := unsafe.Pointer(a_uintptr_first_unsafe_point)
	fmt.Println("a[0]位置指针自增1后，的指针位置，转成unsafe_Pointer类型：", a_uintptr_second_unsafe_point)

	//将该指针转换成 *int8类型（因为它本身就是*int8类型）
	int8_point := (*int8)(a_uintptr_second_unsafe_point)

	//解引用，得到指针对应的结果，就是数组的第二个值，8
	fmt.Println(*int8_point)
}

//可以使用unsafe pointer 去移動位址
