/**
 * @Author : Yannick
 * @File   : learn_interface.go
 * @Date   : 2017-11-11
 * @Desc   : This is a demo which I could learn golang struct through it.
 */

package main

import (
	"fmt"
)

// Student :
type Student struct {
	Persion // anonymous attribute, define in learn_method.go
	school  string
	phone   string
}

// LearnStruct :
func LearnStruct() {
	fmt.Println("## LearnSliceArrayAddress() called begin ..")

	// Initialize way1, note that don't miss Persion type.
	// s1 := Student{Persion{100001, "Bob"}, "HUST", "138-xxxx-7611"}
	// Compile err: missing type in composite literal
	s1 := Student{Persion{100001, "Bob"}, "HUST", "138-xxxx-7611"}
	fmt.Printf("s1 => %v\n", s1)

	// Initialize way2, note that dont's miss Persion:
	// s2 := Student{Persion{id: 100002, name: "Deve"}, school: "oakland college", phone: "139-xxxx-0884"}
	// mixture of field:value and value initializers
	s2 := Student{
		Persion: Persion{id: 100002, name: "Deve"},
		school:  "oakland college",
		phone:   "139-xxxx-0884", // 这里如果}在phone这一行的后面，则phone这一行的逗号不是必须。但是换行后逗号是必须的。
	}
	fmt.Printf("s2 => %v\n", s2)

	// s3 is an anonymous struct.
	// Anonymous struct declarations must be initialized and written together with initialization
	s3 := struct {
		name string
		age  int
	}{
		name: "Ted",
		age:  24,
	}
	fmt.Printf("s3 => %v\n", s3)

	fmt.Println("## LearnSliceArrayAddress() called end ..")
}
