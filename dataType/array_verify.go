package main

import (
	"fmt"
)

//ArrayVerify : to verify some features of an array data struct.
func ArrayVerify() {

	fmt.Println("## ArrayVerify() called begin ..")

	// ways of init an array
	// create an array with length 5, and its element values are: 1, 2, 3, 4, 5.
	arr1 := [5]int{1, 2, 3, 4, 5}
	// way1 to loop an array
	for index, value := range arr1 {
		fmt.Printf("arr1[%d] = > %d\n", index, value)
	}
	fmt.Printf("the size of arr1 => %d\n\n", len(arr1))

	/*
		create an array with length 5, and its element values are: 1, 2, 0, 0, 0.
		In initialization, the element without specifying the initial value will be assigned to the default
		value of int of its element type, and the default value of the string is""
	*/
	arr2 := [5]int{1, 2}
	// way2 to loop an array
	for index := 0; index < len(arr2); index++ {
		fmt.Printf("arr2[%d] = > %d\n", index, arr2[index])
	}
	fmt.Printf("the size of arr2 => %d\n\n", len(arr2))

	/*
		create an array with length 4, and its element values are: 1, 2, 3, 4.
		Its length is determined by the number of elements specified at initialization
	*/
	arr3 := [...]int{1, 2, 3, 4}
	for index := 0; index < len(arr3); index++ {
		fmt.Printf("arr3[%d] = > %d\n", index, arr3[index])
	}
	fmt.Printf("the size of arr3 => %d\n\n", len(arr3))

	/*
		create an array with length 6, and its element values are key-values: 0:0, 0:0, 2:1, 3:2, 4:3, 0:0.
		Its length is determined by the given size, and the unrefered key-value is 0:0.
	*/
	arr4 := [6]int{2: 1, 3: 2, 4: 3}
	for index := 0; index < len(arr4); index++ {
		fmt.Printf("arr4[%d] = > %d\n", index, arr4[index])
	}
	fmt.Printf("the size of arr4 => %d\n\n", len(arr4))

	/*
		create an array with length 5, and its element values are key-values: 0:0, 0:0, 2:1, 3:2, 4:3.
		Its length is determined by the max index of the key.
	*/
	arr5 := [...]int{2: 1, 3: 2, 4: 3}
	for index := 0; index < len(arr5); index++ {
		fmt.Printf("arr5[%d] = > %d\n", index, arr5[index])
	}
	fmt.Printf("the size of arr5 => %d\n\n", len(arr5))

	// An array is a value type. When an array is assigned to another array, a new element is copied.
	arr6 := [4]int{1, 56, 0, -19}
	fmt.Printf("arr6 => %d, and address of arr6 is %v\n", arr6, &arr6[0]) // can not use &arr6, why?
	arr7 := arr6
	fmt.Printf("arr7 => %d, and address of arr7 is %v\n", arr7, &arr7[0])
	arr7[0] = arr7[1] + arr7[3]
	fmt.Printf("after assign & calc, arr6 => %d, and address of arr6 is %v\n", arr6, &arr6[0])
	fmt.Printf("after assign & calc, arr7 => %d, and address of arr7 is %v\n\n", arr7, &arr7[0])

	// Two dimensional array
	arr8 := [4][3]int{{10, 11, 17}, {0, 0, 0}, {20, -9, 4}}
	fmt.Printf("arr8 is =>%d, len of arr is => %d\n\n", arr8, len(arr8))

	// loop 1
	for i := 0; i < len(arr8); i++ {
		for j := 0; j < len(arr8[0]); j++ {
			fmt.Printf("arr8[%d][%d] => %d, ", i, j, arr8[i][j])
		}
		fmt.Printf("\n")

	}
	fmt.Printf("\n")
	// loop 2

	for _, subArry := range arr8 {
		for _, v := range subArry {
			fmt.Printf("%d ", v)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
	fmt.Println("## ArrayVerify() called end ..")
	fmt.Println()
}
