/**
 * @Author : Yannick
 * @File   : learn_slice_array_addr.go
 * @Date   : 2017-11-08
 * @Desc   : This is a demo which I could learn golang through it.
 */

package main

import (
	"fmt"
)

/*
Verify :
Inorder to study the reltaionship of that what's the address of an array, and what's the address of s alice?
	- what's &array refers to?
	- what's &array[0] refers to?
	- what's the address of an array?
	- what's &slice refers to?
	- what's &slice[0] refers to?
	- what's the address of a slice itself?
	- what's the address which a slice refers to an actual array?
*/
func LearnSliceArrayAddress() {

	fmt.Println("## LearnSliceArrayAddress() called begin ..")

	array1 := []int{10, 11, 12, 13, 14}
	fmt.Printf("%p\n", &array1)    // what's this?
	fmt.Printf("%p\n", &array1[0]) // address of array1
	fmt.Printf("%p\n\n", array1)   // address of array1

	fmt.Printf("%d\n", *&array1)    // value of array1
	fmt.Printf("%d\n", *&array1[0]) // value of array1[0]
	fmt.Printf("%d\n\n", array1)    // value of array1
	// fmt.Printf("%d\n\n", *array1)   // runntime error, not compile error

	slice1 := array1[0:3]
	fmt.Printf("%p\n", &slice1)    // address of the struct of slice1 itself
	fmt.Printf("%p\n", &slice1[0]) // address of the array which slice1 refers to(array1, also array1[0])
	fmt.Printf("%p\n\n", slice1)   // address of the array which slice1 refers to(array1, also array1[0])

	fmt.Printf("%d\n", *&slice1)    // value of slice1, actually is array1
	fmt.Printf("%d\n", *&slice1[0]) // value of slice1, actually is array1[0]
	fmt.Printf("%d\n\n", slice1)    // value of slice1, actually is array1
	// fmt.Printf("%d\n\n", *slice1)   // runntime error, not compile error

	fmt.Println("## LearnSliceArrayAddress() called end ..")
	fmt.Println()

}
