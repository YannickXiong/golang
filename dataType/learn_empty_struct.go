/**
 * @Author : Yannick
 * @File   : learn_interface.go
 * @Date   : 2017-11-30
 * @Desc   : This is a demo which I could learn golang empty struct through it.
 */

package main

import (
	"fmt"
	"unsafe"
)

type IAVPType interface {
	Value() interface{}
	String() string
}

type AVPTypeDesc struct {
	name     string   // size 16
	dataType IAVPType // size 16
}

type AVPType struct{}

var avpTypeT AVPType

func (a AVPType) Value() interface{} {
	ret := []byte("hello")
	return ret
}

func (a AVPType) String() string {
	return "world"
}

// LearnEmptyStruct :
func LearnEmptyStruct() {
	avpDesc := [2]AVPTypeDesc{

		// if AVPType does not implement IAVPType interface, then cannot use AVPType as an interface
		// in field value. Errors is:
		// cannot use AVPType literal (type AVPType) as type IAVPType in field value:
		// AVPType does not implement IAVPType (missing String method)
		0: {"id", AVPType{}}, //ok
		1: {"id", avpTypeT},  //ok, is the same to {"username", AVPType{}}, avpDesc => [{username {}} {username {}}]
		// 0: {"username", AVPType}, // not ok
		// 0: {"username", new(AVPType)}, //ok,  avpDesc => [{username 0x116ccf0}]
		// 0: {"username", *new(AVPType)}, //ok, avpDesc => [{username {}} {username {}}]

	}

	fmt.Printf("avpDesc => %v\n", avpDesc)
	fmt.Printf("size : => %d\n", unsafe.Sizeof(avpDesc))
	fmt.Printf("size : avpDesc[0]=> %d\n", unsafe.Sizeof(avpDesc[0]))
	fmt.Printf("size : avpDesc[0].name => %d\n", unsafe.Sizeof(avpDesc[0].name))
	fmt.Printf("size : avpDesc[0].dataType => %d\n", unsafe.Sizeof(avpDesc[0].dataType))
	fmt.Printf("size : dataType.String => %d\n", unsafe.Sizeof(avpDesc[0].dataType.String))
	fmt.Printf("size : avpDesc[0].dataType.String() => %d\n", unsafe.Sizeof(avpDesc[0].dataType.String()))
	fmt.Printf("size : avpDesc[0].dataType.Value => %d\n", unsafe.Sizeof(avpDesc[0].dataType.Value))
	fmt.Printf("size : avpDesc[0].dataType.Value() => %d\n", unsafe.Sizeof(avpDesc[0].dataType.Value()))

	var a struct{}
	var b struct{}

	// empty struct has 0 size, and the same address!
	fmt.Printf("a size : => %d, addr => %p\n", unsafe.Sizeof(a), &a)
	fmt.Printf("b size : => %d, addr => %p\n", unsafe.Sizeof(b), &b)
	fmt.Println("a == b ? => ", a == b)
}
