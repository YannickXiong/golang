/**
 * @Author : Yannick
 * @File   : learn_anonymous_func.go
 * @Date   : 2017-11-09
 * @Desc   : This is a demo which I could learn golang anonymous function through it.
 */

package main

import (
	"fmt"
)

/**
* Squares1: to test anonymous func feature in golang.
* @parameter: none
* @return: a anonymous func() int
**/
func Squares1() func() int {
	var x int // 0 default
	x = x + 1
	return func() int {
		return x * x
	}
}

/**
* Squares2: to test anonymous func feature in golang.
* @parameter: none
* @return: a anonymous func() int
**/
func Squares2() func() int {
	var x int // 0 default
	x = x + 1
	return func() int {
		x++
		return x * x
	}
}

// LearnAnonymousFunc :
func LearnAnonymousFunc() {
	f1 := Squares1() // returns a pointer which refers to Squares().
	fmt.Printf("The address of f1 => %v\n", f1)

	for i := 0; i < 5; i++ {
		fmt.Printf("The %d times call f1, the addr of f1 =>%v, the value f1() =>%d\n", i, f1, f1())
	}

	f2 := Squares2() // returns a pointer which refers to Squares().
	fmt.Printf("The address of f2 => %v\n", f2)

	for i := 0; i < 5; i++ {
		fmt.Printf("The %d times call f2, the addr of f2 =>%v, the value f2() =>%d\n", i, f2, f2())
	}

	/*
		Conclusion: We see that the life cycle of a variable is not determined by
		its scope: after squares returns, the variable x still implicitly exists in f2.

		This is almost like closure in python language.a

		In addition, compare Squares1 and Squares2, we find that variable assignment in outer
		function (x = x + 1) can affect the inner func one tiems, when f1 := Squares1() called.a
		And the variable assignment inner a function (x++) will affect the inner function every
		time it's been called.

	*/

}
