/**
 * @Author : Yannick
 * @File   : learn_reflect.go
 * @Date   : 2017-11-21
 * @Desc   : learn usages of reflect.
 */

package main

import (
	"encoding/binary"
	"fmt"
	"reflect"
)

func LearnReflect() {
	a := 12
	// annot convert a (type int) to type []byte
	// b := []byte(a)

	// invalid type assertion: a.([]byte) (non-interface type int on left)
	// b := a.([]byte)
	b := []byte("2")
	fmt.Printf("a => %v \t b => %v\n", a, b)
	// reflect.New() returns a Value representing a pointer to a new zero value for the specified type.
	// That is, the returned Value's Type is PtrTo(typ).
	newPtr1 := reflect.New(reflect.TypeOf(a))
	newPtr2 := reflect.New(reflect.TypeOf(b))
	fmt.Printf("reflect.New(reflect.TypeOf(a) => %v \t reflect.New(reflect.TypeOf(b) => %v\n", newPtr1, newPtr2)

	// func (v Value) Elem() Value
	// returns the value that the interface v contains or that the pointer v points to.
	// It panics if v's Kind is not Interface or Ptr. It returns the zero Value if v is nil.
	value1 := newPtr1.Elem()
	value2 := newPtr2.Elem()
	fmt.Printf("newPtr1.Elem() => %v \t rnewPtr2.Elem() => %v\n", value1, value2)

	fmt.Printf("value1.Interface() => %v \t value2.Interface() => %v\n", value1.Interface(), value2.Interface())
	value2.SetUint(uint64(binary.BigEndian.Uint32(b)))
	fmt.Printf("newPtr1.Elem() => %v \t rnewPtr2.Elem() => %v\n", value1, value2)
	fmt.Printf("value1.Interface() => %v \t value2.Interface() => %v\n", value1.Interface(), value2.Interface())

}
