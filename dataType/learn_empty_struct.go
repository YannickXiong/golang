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
	name     string
	dataType IAVPType
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

		0: {"username", AVPType{}}, //ok
		1: {"username", avpTypeT},  //ok, is the same to {"username", AVPType{}}, avpDesc => [{username {}} {username {}}]
		// 0: {"username", AVPType}, // not ok
		// 0: {"username", new(AVPType)}, //ok,  
		avpDesc => [{username 0x116ccf0}]
		// 0: {"username", *new(AVPType)}, //ok, avpDesc => [{username {}} {username {}}]

	}

	fmt.Printf("avpDesc => %v\n", avpDesc)
	fmt.Printf("size : => %d\n", unsafe.Sizeof(avpDesc))
}
