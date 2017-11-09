package main

import (
	"fmt"
)

// AddTest1 : a function,  which is no explicitly specified receiver.
// But when called: p = AddTest(a), p is the receiver of AddTest.
func AddTest1(a int) int {
	return a + 10
}

// AddTest : try to overwrite AddTest, will lead error: cannot use a (type int) as type *int in argument to AddTest
// Note: Golang does not support function overwrite, but support interface-method overwrite.
// func AddTest1(a *int) int {
// 	return *a + 10
// }

// AddTest2 :
func AddTest2(a *int) int {
	return *a + 10
}

// Persion :
type Persion struct {
	id   int
	name string
}

// ShowName1 : a method, which is explicitly specified receiver(p Persion)
// the receiver is a value type
func (p Persion) ShowName1() {
	fmt.Printf("ShowName1::Name => %s\n", p.name)
}

// ShowName2 : a method, which is explicitly specified receiver(p Persion)
// the receiver is a reference type
func (p *Persion) ShowName2() {
	fmt.Printf("ShowName2::Name => %s\n", p.name)
}

// LearnMethod : test the difference between a function and a method when faced
// a value type and a preference type arguments.
// For a function, if the receiver is value type(AddTest1), the arguments can only be value type,
// so is the preference type(AddTest2)
// While for a method, if the receiver is value type(AddTest1), the arguments can only be value type,
// so is the preference type(AddTest2)
func LearnMethod() {
	fmt.Println("## LearnMethod() called begin ..")

	a := 2
	fmt.Printf("a => %d\n", a)
	// OK
	fmt.Printf("AddTest1(a) => %d\n", AddTest1(a))
	// OK
	fmt.Printf("AddTest2(&a) => %d\n", AddTest2(&a))
	// NOK, cannot use &a (type *int) as type int in argument to AddTest1
	// fmt.Printf("AddTest1(&a) => %d\n", AddTest1(&a))
	// NOK, cannot use a (type int) as type *int in argument to AddTest2
	// fmt.Printf("AddTest2(a) => %d\n", AddTest2(a))

	// Note: not p1 := Persion(10001, "Bom Smith")
	p1 := Persion{10001, "Bom Smith"}
	// OK
	p1.ShowName1()
	// OK
	p1.ShowName2()

	p2 := &Persion{10002, "Devid Jone"}
	// OK
	p2.ShowName1()
	// OK
	p2.ShowName2()

	fmt.Println("## LearnMethod() called end ..")

}
