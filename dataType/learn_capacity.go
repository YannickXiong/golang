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
	// It was futrher found that all types of pointers have size of 8.
	fmt.Println("------------------ int = CPU Nums------------------")
	var iVar int
	var ipVar *int
	fmt.Printf("var iVar int \t\tSizeof(iVar) = > %d\n", unsafe.Sizeof(iVar))
	fmt.Printf("var ipVar *int \t\tSizeof(ipVar) = > %d\n", unsafe.Sizeof(ipVar))
	var iVar1 = 1
	var ipVar1 = &iVar1
	fmt.Printf("var iVar1 = 1 \t\tSizeof(iVar1) = > %d\n", unsafe.Sizeof(iVar1))
	fmt.Printf("var ipVar1 = &iVar1 \tSizeof(ipVar1) = > %d\n", unsafe.Sizeof(ipVar1))

	fmt.Println("------------------ uint = CPU Nums ------------------")
	var uiVar uint
	var uipVar *uint
	fmt.Printf("var uiVar uint \t\tSizeof(uiVar) = > %d\n", unsafe.Sizeof(uiVar))
	fmt.Printf("var uipVar *uint \t\tSizeof(uipVar) = > %d\n", unsafe.Sizeof(uipVar))
	var uiVar1 = 1
	var uipVar1 = &uiVar1
	fmt.Printf("var uiVar1 = 1 \t\tSizeof(uiVar1) = > %d\n", unsafe.Sizeof(uiVar1))
	fmt.Printf("var uipVar1 = &uiVar1 \tSizeof(uipVar1) = > %d\n", unsafe.Sizeof(uipVar1))

	fmt.Println("------------------ byte = uint8 ------------------")
	var byVar byte
	var bypVar *byte
	fmt.Printf("var byVar byte \tSizeof(byVar) = > %d\n", unsafe.Sizeof(byVar))
	fmt.Printf("var bypVar *byte \tSizeof(bypVar) = > %d\n", unsafe.Sizeof(bypVar))
	var byVar1 = byte(1)
	var bypVar1 = &byVar1
	fmt.Printf("var byVar1 = byte(1) \tSizeof(byVar1) = > %d\n", unsafe.Sizeof(byVar1))
	fmt.Printf("var bypVar1 = &byVar1 Sizeof(bypVar1) = > %d\n", unsafe.Sizeof(bypVar1))

	fmt.Println("------------------ int8 ------------------")
	var i8Var int8
	var i8pVar *int8
	fmt.Printf("var i8Var int8 \tSizeof(i8Var) = > %d\n", unsafe.Sizeof(i8Var))
	fmt.Printf("var i8pVar *int8 \tSizeof(i8pVar) = > %d\n", unsafe.Sizeof(i8pVar))
	var i8Var1 = int8(1)
	var i8pVar1 = &i8Var1
	fmt.Printf("var i8Var1 = int8(1) \tSizeof(i8Var1) = > %d\n", unsafe.Sizeof(i8Var1))
	fmt.Printf("var i8pVar1 = &i8Var1 Sizeof(i8pVar1) = > %d\n", unsafe.Sizeof(i8pVar1))

	fmt.Println("------------------ uint8 ------------------")
	var ui8Var uint8
	var ui8pVar *uint8
	fmt.Printf("var ui8Var uint8 \tSizeof(ui8Var) = > %d\n", unsafe.Sizeof(ui8Var))
	fmt.Printf("var ui8pVar *uint8 \tSizeof(ui8pVar) = > %d\n", unsafe.Sizeof(ui8pVar))
	var ui8Var1 = uint8(1)
	var ui8pVar1 = &ui8Var1
	fmt.Printf("var ui8Var1 = uint8(1) \tSizeof(ui8Var1) = > %d\n", unsafe.Sizeof(ui8Var1))
	fmt.Printf("var ui8pVar1 = &ui8Var1 Sizeof(ui8pVar1) = > %d\n", unsafe.Sizeof(ui8pVar1))

	fmt.Println("------------------ int16 ------------------")
	var i16Var int16
	var i16pVar *int16
	fmt.Printf("var i16Var int16 \tSizeof(i16Var) = > %d\n", unsafe.Sizeof(i16Var))
	fmt.Printf("var i16pVar *int16 \tSizeof(i16pVar) = > %d\n", unsafe.Sizeof(i16pVar))
	var i16Var1 = int16(1)
	var i16pVar1 = &i16Var1
	fmt.Printf("var i16Var1 = int16(1) \tSizeof(i16Var1) = > %d\n", unsafe.Sizeof(i16Var1))
	fmt.Printf("var i16pVar1 = &i16Var1 Sizeof(i16pVar1) = > %d\n", unsafe.Sizeof(i16pVar1))

	fmt.Println("------------------ uint16 ------------------")
	var ui16Var uint16
	var ui16pVar *uint16
	fmt.Printf("var ui16Var uint16 \tSizeof(ui16Var) = > %d\n", unsafe.Sizeof(ui16Var))
	fmt.Printf("var ui16pVar *uint16 \tSizeof(ui16pVar) = > %d\n", unsafe.Sizeof(ui16pVar))
	var ui16Var1 = uint16(1)
	var ui16pVar1 = &ui16Var1
	fmt.Printf("var ui16Var1 = uint16(1) \tSizeof(ui16Var1) = > %d\n", unsafe.Sizeof(ui16Var1))
	fmt.Printf("var ui16pVar1 = &ui16Var1 Sizeof(ui16pVar1) = > %d\n", unsafe.Sizeof(ui16pVar1))

	fmt.Println("------------------ int32 ------------------")
	var i32Var int32
	var i32pVar *int32
	fmt.Printf("var i32Var int32 \tSizeof(i32Var) = > %d\n", unsafe.Sizeof(i32Var))
	fmt.Printf("var i32pVar *int32 \tSizeof(i32pVar) = > %d\n", unsafe.Sizeof(i32pVar))
	var i32Var1 = int32(1)
	var i32pVar1 = &i32Var1
	fmt.Printf("var i32Var1 = int32(1) \tSizeof(i32Var1) = > %d\n", unsafe.Sizeof(i32Var1))
	fmt.Printf("var i32pVar1 = &i32Var1 Sizeof(i32pVar1) = > %d\n", unsafe.Sizeof(i32pVar1))

	fmt.Println("------------------ rune = int32 ------------------")
	var ruVar rune
	var rupVar *rune
	fmt.Printf("var ruVar rune \tSizeof(ruVar) = > %d\n", unsafe.Sizeof(ruVar))
	fmt.Printf("var rupVar *rune \tSizeof(rupVar) = > %d\n", unsafe.Sizeof(rupVar))
	var ruVar1 = rune(1)
	var rupVar1 = &ruVar1
	fmt.Printf("var ruVar1 = rune(1) \tSizeof(ruVar1) = > %d\n", unsafe.Sizeof(ruVar1))
	fmt.Printf("var rupVar1 = &ruVar1 Sizeof(rupVar1) = > %d\n", unsafe.Sizeof(rupVar1))

	fmt.Println("------------------ uint32 ------------------")
	var ui32Var uint32
	var ui32pVar *uint32
	fmt.Printf("var ui32Var uint32 \tSizeof(ui32Var) = > %d\n", unsafe.Sizeof(ui32Var))
	fmt.Printf("var ui32pVar *uint32 \tSizeof(ui32pVar) = > %d\n", unsafe.Sizeof(ui32pVar))
	var ui32Var1 = uint32(1)
	var ui32pVar1 = &ui32Var1
	fmt.Printf("var ui32Var1 = uint32(1) \tSizeof(ui32Var1) = > %d\n", unsafe.Sizeof(ui32Var1))
	fmt.Printf("var ui32pVar1 = &ui32Var1 Sizeof(ui32pVar1) = > %d\n", unsafe.Sizeof(ui32pVar1))

	fmt.Println("------------------ int64 ------------------")
	var i64Var int64
	var i64pVar *int64
	fmt.Printf("var i64Var int64 \tSizeof(i64Var) = > %d\n", unsafe.Sizeof(i64Var))
	fmt.Printf("var i64pVar *int64 \tSizeof(i64pVar) = > %d\n", unsafe.Sizeof(i64pVar))
	var i64Var1 = int64(1)
	var i64pVar1 = &i64Var1
	fmt.Printf("var i64Var1 = int64(1) \tSizeof(i64Var1) = > %d\n", unsafe.Sizeof(i64Var1))
	fmt.Printf("var i64pVar1 = &i64Var1 Sizeof(i64pVar1) = > %d\n", unsafe.Sizeof(i64pVar1))

	fmt.Println("------------------ uint64 ------------------")
	var ui64Var uint64
	var ui64pVar *uint64
	fmt.Printf("var ui64Var uint64 \tSizeof(ui64Var) = > %d\n", unsafe.Sizeof(ui64Var))
	fmt.Printf("var ui64pVar *uint64 \tSizeof(ui64pVar) = > %d\n", unsafe.Sizeof(ui64pVar))
	var ui64Var1 = uint64(1)
	var ui64pVar1 = &ui64Var1
	fmt.Printf("var ui64Var1 = uint64(1) \tSizeof(ui64Var1) = > %d\n", unsafe.Sizeof(ui64Var1))
	fmt.Printf("var ui64pVar1 = &ui64Var1 Sizeof(ui64pVar1) = > %d\n", unsafe.Sizeof(ui64pVar1))

	fmt.Println("------------------ float32 ------------------")
	var f32Var float32
	var f32pVar *float32
	fmt.Printf("var f32Var float32 \tSizeof(f32Var) = > %d\n", unsafe.Sizeof(f32Var))
	fmt.Printf("var f32pVar *float32 \tSizeof(f32pVar) = > %d\n", unsafe.Sizeof(f32pVar))
	var f32Var1 = float32(1.0)
	var f32pVar1 = &f32Var1
	fmt.Printf("var f32Var1 = float32(1) \tSizeof(f32Var1) = > %d\n", unsafe.Sizeof(f32Var1))
	fmt.Printf("var f32pVar1 = &f32Var1 Sizeof(f32pVar1) = > %d\n", unsafe.Sizeof(f32pVar1))

	fmt.Println("------------------ float64 ------------------")
	var f64Var float64
	var f64pVar *float64
	fmt.Printf("var f64Var float64 \tSizeof(f64Var) = > %d\n", unsafe.Sizeof(f64Var))
	fmt.Printf("var f64pVar *float64 \tSizeof(f64pVar) = > %d\n", unsafe.Sizeof(f64pVar))
	var f64Var1 = float64(1.0)
	var f64pVar1 = &f64Var1
	fmt.Printf("var f64Var1 = float64(1) \tSizeof(f64Var1) = > %d\n", unsafe.Sizeof(f64Var1))
	fmt.Printf("var f64pVar1 = &f64Var1 Sizeof(f64pVar1) = > %d\n", unsafe.Sizeof(f64pVar1))

	fmt.Println("------------------ string ------------------")
	var sVar string
	var spVar *string
	fmt.Printf("var sVar string \t\t\tSizeof(sVar) = > %d\n", unsafe.Sizeof(sVar))
	fmt.Printf("var spVar *string \t\t\tSizeof(spVar) = > %d\n", unsafe.Sizeof(spVar))
	// Sizeof(sVar1) is 16, while len(sVar1) is 35
	var sVar1 = "hello world, young man"
	var spVar1 = &sVar1
	fmt.Printf("var sVar1 = \"hello world, young man\" \tSizeof(sVar1) => %d len(sVar1) = > %d\n", unsafe.Sizeof(sVar1), len(sVar1))
	fmt.Printf("var spVar1 = &sVar1 \t\t\tSizeof(spVar1) = > %d\n", unsafe.Sizeof(spVar1))

	fmt.Println("------------------ bool ------------------")
	var bVar bool
	var bpVar *bool
	fmt.Printf("var bVar bool \t\tSizeof(bVar) = > %d\n", unsafe.Sizeof(bVar))
	fmt.Printf("var bpVar *bool \tSizeof(bpVar) = > %d\n", unsafe.Sizeof(bpVar))
	// bool type does not support len()
	var bVar1 = false
	var bpVar1 = &bVar1
	fmt.Printf("var bVar1 = false \tSizeof(bVar1) => %d\n", unsafe.Sizeof(bVar1))
	fmt.Printf("var bpVar1 =&bVar1 \tSizeof(bpVar1) = > %d\n", unsafe.Sizeof(bpVar1))

	fmt.Println("------------------ struct ------------------")
	type Phone struct {
		name   string //size 16
		romCap uint8  // size 1
	}

	var phone Phone
	var pphone *Phone
	// Sizeof(phone) is 24, not 16 + 1 = 17, just because of memory padding.
	fmt.Printf("var phone Phone \tSizeof(phone) = > %d\n", unsafe.Sizeof(phone))
	fmt.Printf("phone.name \t\tSizeof(phone.name) = > %d\n", unsafe.Sizeof(phone.name))
	fmt.Printf("phone.romCap \t\tSizeof(phone.romCap) = > %d\n", unsafe.Sizeof(phone.romCap))
	fmt.Printf("var pphone * phone \tSizeof(pphone) = > %d\n", unsafe.Sizeof(pphone))

}
