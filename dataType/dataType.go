package main

import (
	"fmt"
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

}
