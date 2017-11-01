package main

import (
	"fmt"
	"unsafe"
)

// the default value of var int is 0
var x1 int

/*	not recommended
	given a warming:
	should omit type int from declaration of var x2; it will be inferred from the right-hand side
*/
var x2 int = 10
var x3 = 10
var x4 = 11

var f1 float32
var f2 float32 = 1.2345
var f3 = 1.2345

// the default value of var string is none.
var s1 string
var s2 string = "test abc"
var s3 = "test abc"

// the default value of var bool is false.
var b1 bool
var b2 bool = true
var b3 = false

/*
   var b4 bool = 5
   compile error: cannot use 5 (type int) as type bool in assignment
*/

// PI = 3.141516926, a const, changeable.
const PI float32 = 3.141516926
const y1, y2 = 0, 1
const logLevel = "DEBUG"

// const can be a result of operating
const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(b)
)

// enum, use iota to define a constant group with index starts from 0 to an enumeration value
const (
	Sunday    = iota // 0
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

const (
	red    = iota //0
	green         //1
	black  = iota //2
	yellow        //3
	pink          //4
)

/*
MonthType denifition
use custom type to achieve the enumeration type limit.
*/
type MonthType int

const (
	// May MonthType: MonthType is data type, actually is int.
	May  MonthType = iota // 0
	Jone                  // 1
	July                  // 2
)

func test(m MonthType) {}

// array

/*
	compile error, := assignment method can only used inner a function.
	data, i := [3]int{0, 1, 2}
*/

func main() {
	/*
		fmt.Println("var x1 int => %d", x1) work normally, but vet will generate a warming as below.
		possible formatting directive in Println call
	*/
	fmt.Printf("var x1 int => %d\n", x1)
	fmt.Printf("var x2 int = 10 => %d\n", x2)
	fmt.Printf("var x3 = 10 => %d\n\n", x3)

	fmt.Printf("var f1 float32 => %f\n", f1)
	fmt.Printf("var f2 float32 = 1.2345 => %f\n", f2)
	fmt.Printf("var f3 = 1.2345 => %G\n\n", f3)

	fmt.Printf("var s1 string => %s\n", s1)
	// notice the difference between %s and %q.
	fmt.Printf("var s2 string = 'test abc' => %s\n", s2)
	fmt.Printf("var s3 = 'test abc' => %q\n\n", s3)

	fmt.Printf("var b1 bool => %t\n", b1)
	fmt.Printf("var b2 string = 'true' => %t\n", b2)
	fmt.Printf("var b3 = 'false' => %t\n\n", b3)
	/*
		fmt.Printf("the address, value of PI => %v, %s, \n", &PI, PI)
		compile error: cannot take the address of PI, this shows that can not get address of a const?
	*/

	/*
		this assignment method can only be used inner a function.
		and more important is that, un unused local var in a function will lead to a compile error,
		while that will not happen outside of a function, like var x4(var x4 = 11)
	*/
	p1 := 3.14
	fmt.Printf("local var p1 => %f\n", p1)
	fmt.Printf("const var PI => %f\n\n", PI)
	// notice that const var can be changed in go.
	PI := PI - 0.001516926
	fmt.Printf("after PI = PI - 0.001516926, const var PI => %f\n", PI)
	fmt.Printf("after 'PI := PI - 0.001516926', the address, value of PI => %v, %f, \n", &PI, PI)
	// notice: PI = PI - 0.001516926 will lead to compile error, while PI = 5 will work Ok.
	PI = 5
	fmt.Printf("PI = 5, const var PI => %f\n\n", PI)
	fmt.Printf("after 'PI = 5', the address, value of PI => %v, %f, \n", &PI, PI)

	data, i := [3]int{0, 1, 2}, 0
	fmt.Printf("first 'data, i := [3]int{0, 1, 2}, 0', data => %d, i=> %d\n", data, i)
	i, data[i] = 2, 100
	/*
		notice: data[0] = 100, not data[2] = 100 !
		because When multiple variables are assigned, all the relevant values are computed first,
		and then the values are assigned from left to right.
	*/
	fmt.Printf("after 'i, data[i] = 2, 100', data => %d, i=> %d\n\n", data, i)

	c1 := "abc"
	/*
		notice:
		c1 := "efg" will lead to a compile error: no new variables on left side of :=
		c1 := 10 will lead to 2 compile errors: one is the same to above, and another is
		cannot use 10 (type int) as type string in assignment
	*/
	fmt.Printf("the address of c1 => %v\n", &c1)
	c1, c2 := "efg", 10
	/*
		notice: while c1, c2 := "efg", 10 will not lead to a compile error?
		we come back to the compile error: no new variables on left side of :=
		and for c1, c2 := "efg", 10, there is new variables on left side of :=, so this will not
		lead to a compile error.
	*/
	fmt.Printf("the address, value of c1 => %v, %s, the address, value of c2 => %v, %d\n", &c1, c1, &c2, c2)
	c1 = "500"
	// address of c1 will not change, while the value changed.
	fmt.Printf("after 'c1 = '500'', the address, value of c1 => %v, %s, \n\n", &c1, c1)

	/*
		notice:
			1) x2 is outside of funtion, and := inner function will define a new var, adreess, value
			   of x2 will both change.
			2) x2 and := are both outside of function, or both inner a functioin, will re-assigne new
			   value to x2, only value of x2 change, the address of x2 will keep the same.
	*/
	fmt.Printf("the address, value of x2 => %v, %d, \n", &x2, x2)
	x2 := 30
	fmt.Printf("after 'x2 := 30' in a function, he address, value of x2 => %v, %d, \n\n", &x2, x2)

	// iota
	// fmt.Println(iota) will lead to a compile error, because iota can be only used with const.
	fmt.Printf("Sunday => %d\n", Sunday)
	fmt.Printf("Monday => %d\n", Monday)
	fmt.Printf("Tuesday => %d\n", Tuesday)
	fmt.Printf("Wednesday => %d\n", Wednesday)
	fmt.Printf("Thursday => %d\n", Thursday)
	fmt.Printf("Friday => %d\n", Friday)
	fmt.Printf("Saturday => %d\n\n", Saturday)

	// use iota to reset am enum starting from 0.
	fmt.Printf("red => %d\n", red)
	fmt.Printf("green => %d\n", green)
	fmt.Printf("black => %d\n", black)
	fmt.Printf("yellow => %d\n", yellow)
	fmt.Printf("pink => %d\n\n", pink)

	/*
		People usually think that iota is always 0, but that's not true.
		To be exact, when iota follows the first line of the keyword const, iota is 0,
		when the second line appears, iota is 1, and so on;
		when iota meets const again, it is reset to 0.
	*/
	const (
		i1 = iota // 0
		j1 = iota // 1
		k1 = iota // 2
	)
	fmt.Printf("i1 => %d\n", i1)
	fmt.Printf("j1 => %d\n", j1)
	fmt.Printf("k1 => %d\n\n", k1)

	// when iota meets const again, it is reset to 0.
	const i2 = iota // 0
	const j2 = iota // 0
	const k2 = iota // 0
	fmt.Printf("i2 => %d\n", i2)
	fmt.Printf("j2 => %d\n", j2)
	fmt.Printf("k2 => %d\n\n", k2)

	// iota in the first line of const, so i3 = 0, j3 = 0.
	// k3, l3 in the second line, so they are 1.
	// m3, n3 in the third line, so they are 2.
	const (
		i3, j3 = iota, iota // 0
		// only k3 will lead to a compile error, can only k3, l3
		k3, l3 // 1
		m3, n3 // 2
	)
	fmt.Printf("i3 => %d\n", i3)
	fmt.Printf("j3 => %d\n", j3)
	fmt.Printf("k3 => %d\n", k3)
	fmt.Printf("l3 => %d\n", l3)
	fmt.Printf("m3 => %d\n", m3)
	fmt.Printf("n3 => %d\n\n", n3)

	const (
		i4, j4 = iota, iota // 0
		k4, l4              // 1
		m4, n4 = iota, iota // 2
	)
	fmt.Printf("i4 => %d\n", i4)
	fmt.Printf("j4 => %d\n", j4)
	fmt.Printf("k4 => %d\n", k4)
	fmt.Printf("l4 => %d\n", l4)
	fmt.Printf("m4 => %d\n", m4)
	fmt.Printf("n4 => %d\n\n", n4)

	const (
		a4, b4 = iota, iota + 7 // 0, 7
		c4, d4                  // 1, 8
		e4, f4 = iota, iota * 3 // 2, 6 think: why is 6(2 * 3), not 9(2 + 7)?
	)
	fmt.Printf("a4 => %d\n", a4)
	fmt.Printf("b4 => %d\n", b4)
	fmt.Printf("c4 => %d\n", c4)
	fmt.Printf("d4 => %d\n", d4)
	fmt.Printf("e4 => %d\n", e4)
	fmt.Printf("f4 => %d\n\n", f4)

	const (
		i5 = -1   // -1
		j5 = iota // 1, not zero, because iota appears in the second line of keyword const, j5 = 1, not 0.
		k5        // 2
	)
	fmt.Printf("i5 => %d\n", i5)
	fmt.Printf("j5 => %d\n", j5)
	fmt.Printf("k5 => %d\n\n", k5)

	const (
		_        = iota             // iota = 0
		KB int64 = 1 << (10 * iota) // iota = 1
		MB                          // 1<< (10 * 2)
		GB                          // 1<< (10 * 3)
		TB                          // 1<< (10 * 4)
	)
	fmt.Printf("KB => %d\n", KB)
	fmt.Printf("MB => %d\n", MB)
	fmt.Printf("GB => %d\n", GB)
	fmt.Printf("TB => %d\n\n", TB)

	const (
		A = iota // 0
		B        // 1
		C = "c"  // c, interrupt the iota.
		D        // c，the same to the previous line
		E = iota // 4，recover iota, and include the line of C and B.
		F        // 5
	)

	fmt.Printf("A => %d\n", A)
	fmt.Printf("B => %d\n", B)
	fmt.Printf("C => %v\n", C)
	fmt.Printf("D => %v\n", D)
	fmt.Printf("E => %v\n", E)
	fmt.Printf("F => %v\n\n", F)

	c := Jone
	test(c)
	// x := 1
	// test(x): lead to a compile error: cannot use x (type int) as type MonthType in argument to test

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

	/*
	   Slice: slice is a reference type.
	   There are 2 ways to create a slice, one is by array, another is by make()
	*/
	// by array, [] means that this is a slice, len == cap == 3
	s1 := []int{1, 2, 3}
	fmt.Printf("len of s1 => %d, cap of s1 => %d\n", len(s1), cap(s1))

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

}
