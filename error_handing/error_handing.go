/**
* @Author  : Yannick
* @File    : learn_anonymous_func.go
* @Date    : 2017-11-14
* @Desc    : Pracise to use error handing in go.
**/

package main

import (
	"errors"
	"fmt"
)

/*
* In golang, error type is actually a interface with an Error() string method.append
* type error interface{
	Error() string
}
*/

// errors, always used to be as the last return para in a function, it's service outside.
// MyDev : Hands error with errors.New()
func MyDev(a int, b int) (result int, err error) {
	if b == 0 {
		return -1, errors.New("the dividend cannot be zero")
	}

	// 10 / 3 => 3
	return a / b, nil
}

// ErrMsg :
type ErrMsg struct {
	ErrNo int
	Msg   string
}

// Error: ErrMsg type implements Error interface.
func (e ErrMsg) Error() string {
	return fmt.Sprintf("ErrNo is %d, msg is %s", e.ErrNo, e.Msg)
}

// MyDev1 : Hand error With a custom parameter error output
func MyDev1(a int, b int) (result int, err error) {
	if b == 0 {
		return -1, ErrMsg{
			ErrNo: 3001,
			Msg:   "the dividend cannot be zero",
		}
	}

	// 10 / 3 => 3
	return a / b, nil
}

// LearnErrHanding :
func LearnErrHanding() {

	if result1, err := MyDev(10, 3); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("result := %d\n", result1)
	}

	if result2, err := MyDev(10, 0); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("result := %d\n", result2)
	}

	if result3, err := MyDev1(10, 3); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("result := %d\n", result3)
	}

	if result4, err := MyDev1(10, 0); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("result := %d\n", result4)
	}
}
