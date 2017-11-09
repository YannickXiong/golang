/**
 * @Author : Yannick
 * @File   : learn_iota.go
 * @Date   : 2017-11-08
 * @Desc   : This is a demo which I could learn golang through it.
 */

package main

import (
	"fmt"
)

// LearnIota : to verify some features of an iota data struct.
func LearnIota() {

	fmt.Println("## LearnIota() called begin ..")

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
	Test(c)
	// x := 1
	// test(x): lead to a compile error: cannot use x (type int) as type MonthType in argument to test

	fmt.Println("## LearnIota() called end ..")
	fmt.Println()

}
