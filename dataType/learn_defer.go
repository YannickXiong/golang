/*
 * @Author: Yannick
 * @Date: 2017-11-18 19:52:26
 * @Last Modified by: Yannick
 * @Last Modified time: 2017-11-18 22:13:28
 */

package main

import (
	"fmt"
)

// 所有含有defer的函数都可以改写：
// 返回值 = xxx
// 调用defer函数
// 空的return

// deferF1 :
func deferF1() (result int) {
	defer func() { // An anonymous function
		result++
	}() // a closure

	return 0
}

// deferF2 :
func deferF2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()

	return t
}

// 根据规则改写deferF2
// func f() (r int) {
// 	t := 5
// 	r = t //赋值指令
// 	func() {        //defer被插入到赋值与返回之间执行，这个例子中返回值r没被修改过
// 		t = t + 5
// 	}
// 	return        //空的return指令
// }

// deferF3 :
func deferF3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r) // 注意这里要传一个实参呀！

	return 1
}

// 根据规则改写deferF3
// func f() (r int) {
// 	r = 1  //给返回值赋值
// 	func(r int) {        //这里改的r是传值传进去的r，不会改变要返回的那个r值
// 		 r = r + 5
// 	}(r)
// 	return        //空的return
// }

// LearnDefer :
func LearnDefer() {
	// defer is executed by FILO order
	for i := 0; i < 5; i++ {
		defer fmt.Println("Executed ", i)
	}

	// First in last out, defer fmt.Println("Executed ", i) will be
	// executed after deferF1()
	r1 := deferF1()
	fmt.Printf("r1 => %d\n", r1)

	r2 := deferF2()
	fmt.Printf("r2 => %d\n", r2)

	r3 := deferF3()
	fmt.Printf("r3 => %d\n", r3)
}
