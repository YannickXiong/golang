/**
 * @Author : Yannick
 * @File   : learn_switch.go
 * @Date   : 2017-12-04
 * @Desc   : To learn usage of swich.
 */

package main

import (
	"fmt"
)

func LearnSwitch() {
	for i := 0; i < 10; i++ {
		switch i {
		case 0:
			fmt.Printf("i = 0\n") // break is no need
		case 1, 2, 3: // can match multiple value
			i = i + 2
			fmt.Printf("i = 1, or 2, or 3\n")
		case 4:
			fmt.Printf("i = 4\n")
			fallthrough // case 100 cannot be reached, but with fallthrough, they can be reached
			// callthrough does only affect the neartest one, which means that 101 case cannot be reached.
		case 100:
			fmt.Printf("i = 100, it cannot be reached, but with fallthrough, this can be reached.\n")
		case 101:
			fmt.Printf("i = 101, it cannot be reached, but with fallthrough, this can be reached.\n")
		default:
			i = i + 5
			fmt.Printf("i = 4, or 6, or 7\n")
		}
	}
}
