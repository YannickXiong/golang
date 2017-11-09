/**
 * @Author : Yannick
 * @File   : learn_slice.go
 * @Date   : 2017-11-08
 * @Desc   : This is a demo which I could learn golang through it.
 */

package main

import (
	"fmt"
)

// LearnSlice : to verify some feature of a slice data struct.
func LearnSlice() {

	fmt.Println("## LearnSlice() called begin ..")

	/*
	   Slice: slice is a reference type.
	   There are 2 ways to create a slice, one is by array, another is by make()
	*/
	// by array, [] means that this is a slice, len == cap == 3

	// nil slice
	var snil []int
	fmt.Printf("nil slice: snil =>%v, len() => %d, cap() => %d, self address => %p, refers to array address => %p\n", snil, len(snil), cap(snil), &snil, snil)

	// empty slice
	sempty := make([]int, 0, 0)
	fmt.Printf("empty slice:sempty =>%v, len() => %d, cap() => %d, self address => %p, refers to array address => %p\n", sempty, len(sempty), cap(sempty), &sempty, sempty)

	s1 := []int{1, 2, 3}
	fmt.Printf("len of s1 => %d, cap of s1 => %d\n", len(s1), cap(s1))

	// An array is a value type. When an array is assigned to another array, a new element is copied.
	arr6 := [4]int{1, 56, 0, -19}
	fmt.Printf("arr6 => %d, and address of arr6 is %v\n", arr6, &arr6[0]) // can not use &arr6, why?
	arr7 := arr6
	fmt.Printf("arr7 => %d, and address of arr7 is %v\n", arr7, &arr7[0])
	arr7[0] = arr7[1] + arr7[3]
	fmt.Printf("after assign & calc, arr6 => %d, and address of arr6 is %v\n", arr6, &arr6[0])
	fmt.Printf("after assign & calc, arr7 => %d, and address of arr7 is %v\n\n", arr7, &arr7[0])

	// s2 is reference of arr7
	fmt.Printf("arr7 => %d\n", arr7)
	s2 := arr7[:]
	fmt.Printf("s2 := arr7[:], len of s2 => %d, cap of s2 => %d, s2 =>%d\n", len(s2), cap(s2), s2)

	// [start:end], includes start .. end -1, not includes end.
	s3 := arr7[1:2]
	fmt.Printf("s3 := arr7[1:2], len of s3 => %d, cap of s3 => %d, s3 =>%d\n", len(s3), cap(s3), s3)

	// [:end], the same to [0:end]
	s4 := arr7[:3]
	fmt.Printf("s4 := arr7[:3], len of s4 => %d, cap of s4 => %d, s4 =>%d\n", len(s4), cap(s4), s4)

	// [start:], the same to [start:end+1], notice that this will include the last index.
	s5 := arr7[2:]
	fmt.Printf("s5 := arr7[2:], len of s5 => %d, cap of s5 => %d, s5 =>%d\n", len(s5), cap(s5), s5)

	// s :=make([]int,len,cap)
	s6 := make([]int, 3, 7)
	fmt.Printf("len of s6 => %d, cap of s6 => %d, s6 =>%d\n", len(s6), cap(s6), s6)

	/*
		Dilatancy principle: refer to source code of append
			newcap := old.cap
		    doublecap := newcap + newcap
		    if cap > doublecap {
		        newcap = cap	// how to calc cap ? Refer to s8, s9, s10, s11 among the fellowing test cases.
		    } else {
		        if old.len < 1024 {
		            newcap = doublecap // increase by 2 times
		        } else {
		            for newcap < cap {
		                newcap += newcap / 4  // increase by 1.25 times
		            }
		        }
		    }

	*/
	// s7 := make([]int{4, 8, 12}, 3, 7), try to init the data when create a slice, will lead to a compile error.
	s7 := make([]int, 5, 7)
	fmt.Printf("len of s7 => %d, cap of s7 => %d, s7 =>%d\n", len(s7), cap(s7), s7)
	// s7 := append(s7, 11, 12, 13, 14, 15) will lead to a compile error. Because no new value at left of :=
	// while s7 = append(s7, 11, 12, 13, 14, 15) is ok

	// cap of s8 is 14, < 2 * 7(cap s6), new cap is 14
	s8 := append(s6, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21)
	fmt.Printf("len of s8 => %d, cap of s8 => %d, s8 =>%d\n", len(s8), cap(s8), s8)

	// cap of s9 is 15, > 2 * 7(cap s6), so new cap is 15 + 1, new cap is always even number, not odd number.
	s9 := append(s6, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22)
	fmt.Printf("len of s9 => %d, cap of s9 => %d, s9 =>%d\n", len(s9), cap(s9), s9)

	// the same, new cap is 16
	s10 := append(s7, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21)
	fmt.Printf("len of s10 => %d, cap of s10 => %d, s11 =>%d\n", len(s10), cap(s10), s10)

	// the same, new cap is 18
	s11 := append(s7, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22)
	fmt.Printf("len of s11 => %d, cap of s11 => %d, s11 =>%d\n\n", len(s11), cap(s11), s11)

	// trap of append
	array1 := []int{10, 11, 12, 13, 14}
	slice1 := array1[1:4]
	fmt.Printf("address of array1 => %p, address of array1[0] => %p, array is =>%d\n", &array1, &array1[0], array1)
	// slice1 refers to : also can using %p => &slice1[0]
	fmt.Printf("slice1 refers to(address of actual array, slice[0] now refers to array1[1]) => %p, slice1 is => %d\n", &slice1[0], slice1)
	fmt.Printf("address of slice 1 (address of slice struct) => %p\n", &slice1)

	array1[1] += 100 // change array1[1], slice1[0], because slice1 refers to array1
	slice1[1] += 200 // change array1[2], slice1[1], because slice1[1] refers to array1[2]
	fmt.Println("\nafter array1[0] += 100 and slice1[1] += 200")
	fmt.Printf("address of array1 => %p, array is => %d\n", array1, array1)
	// slice1 refers to : also can using %p => &slice1[0]
	fmt.Printf("slice1 refers to(address of actual array, slice[0] now refers to array1[1]) => %p, slice1 is =>%d\n", slice1, slice1)
	fmt.Printf("address of slice 1 (address of slice struct) => %p\n", &slice1)

	/*
		2 ways to call append:
		one is append(s, element)
		another is append(s1, s2...), notice ... is every important.
	*/
	slice2 := append(slice1, s11...)
	fmt.Println("\nafter 'slice2 := append(slice1, s11...)'")
	fmt.Printf("address of array1 => %p, array is => %d\n", array1, array1)
	// slice1 refers to : also can using %p => &slice1[0]
	fmt.Printf("slice1 refers to(address of actual array, slice[0] now refers to array1[1]) => %p, slice1 is => %d\n", slice1, slice1)
	fmt.Printf("address of slice 1 (address of slice struct) => %p\n", &slice1)
	fmt.Printf("slice2 refers to(address of new actual array) => %p, slice2 is => %d\n", &slice2[0], slice2)
	fmt.Printf("slice2 refers to(address of new actual array) => %p, slice2 is => %d\n", slice2, slice2)
	fmt.Printf("address of slice 2 (address of slice struct) => %p\n", &slice2)
	// verify the Dilatancy principle again, cap of slice2 should be 20.
	fmt.Printf("len of slice2 => %d, cap of slice2 => %d, slice2 => %d\n\n", len(slice2), cap(slice2), slice2)

	fmt.Println("## LearnSlice() called end ..")
	fmt.Println()

}
