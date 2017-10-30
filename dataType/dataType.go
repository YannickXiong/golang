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

}
