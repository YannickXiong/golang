/**
 * @Author : Yannick
 * @File   : learn_capacity.go
 * @Date   : 2017-11-24
 * @Desc   : To explore who much size of capacity of different data type in go.
 */

package main

import (
	"fmt"
	"unsafe"
)

// LearnCapacity :
func LearnCapacity() {
	// var int & pointer int both have size of 8.
	fmt.Println("------------------ int ------------------")
	var iVar int
	var ipVar *int
	fmt.Printf("var iVar int \t\tSizeof(iVar) = > %d\n", unsafe.Sizeof(iVar))
	fmt.Printf("var ipVar *int \t\tSizeof(ipVar) = > %d\n", unsafe.Sizeof(ipVar))
	var iVar1 = 1
	var ipVar1 = &iVar1
	fmt.Printf("var iVar1 = 1 \t\tSizeof(iVar1) = > %d\n", unsafe.Sizeof(iVar1))
	fmt.Printf("var ipVar1 = &iVar1 \tSizeof(ipVar1) = > %d\n", unsafe.Sizeof(ipVar1))

	fmt.Println("------------------ int16 ------------------")
	var i16Var int16
	var i16pVar *int16
	fmt.Printf("var i16Var int16 \tSizeof(i16Var) = > %d\n", unsafe.Sizeof(i16Var))
	fmt.Printf("var i16pVar *int16 \tSizeof(i16pVar) = > %d\n", unsafe.Sizeof(i16pVar))
	var i16Var1 = int16(1)
	var i16pVar1 = &i16Var1
	fmt.Printf("var i16Var1 = int16(1) \tSizeof(i16Var1) = > %d\n", unsafe.Sizeof(i16Var1))
	fmt.Printf("var i16pVar1 = &i16Var1 Sizeof(i16pVar1) = > %d\n", unsafe.Sizeof(i16pVar1))

	fmt.Println("------------------ int32 ------------------")
	var i32Var int32
	var i32pVar *int32
	fmt.Printf("var i32Var int32 \tSizeof(i32Var) = > %d\n", unsafe.Sizeof(i32Var))
	fmt.Printf("var i32pVar *int32 \tSizeof(i32pVar) = > %d\n", unsafe.Sizeof(i32pVar))
	var i32Var1 = int32(1)
	var i32pVar1 = &i32Var1
	fmt.Printf("var i32Var1 = int32(1) \tSizeof(i32Var1) = > %d\n", unsafe.Sizeof(i32Var1))
	fmt.Printf("var i32pVar1 = &i32Var1 Sizeof(i32pVar1) = > %d\n", unsafe.Sizeof(i32pVar1))

	fmt.Println("------------------ int64 ------------------")
	var i64Var int64
	var i64pVar *int64
	fmt.Printf("var i64Var int64 \tSizeof(i64Var) = > %d\n", unsafe.Sizeof(i64Var))
	fmt.Printf("var i64pVar *int64 \tSizeof(i64pVar) = > %d\n", unsafe.Sizeof(i64pVar))
	var i64Var1 = int64(1)
	var i64pVar1 = &i64Var1
	fmt.Printf("var i64Var1 = int64(1) \tSizeof(i64Var1) = > %d\n", unsafe.Sizeof(i64Var1))
	fmt.Printf("var i64pVar1 = &i64Var1 Sizeof(i64pVar1) = > %d\n", unsafe.Sizeof(i64pVar1))

}
