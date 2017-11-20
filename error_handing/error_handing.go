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

func funcA() error {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("panic recover! p: %v\n", p)
			// debug.PrintStack()
		}
	}()
	return funcB()
}

func funcD() (err error) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("panic recover! p:", p)
			// fmt.Printf("str => %v, ok => %v\n", str, ok)
			// str => foo, ok => true

			// 跟funcA相比，panic异常处理机制不会自动将错误信息传递给error，所以要在funcD函数中进行显式的传递，代码如下所示：
			str, ok := p.(string)
			if ok {
				err = errors.New(str)
			} else {
				err = errors.New("panic")
			}
			// debug.PrintStack()
		}
	}()
	return funcB()
}

func funcB() error {
	// simulation
	panic("foo")
	return errors.New("success")
}

func test1() {
	err := funcA()
	if err == nil {
		fmt.Printf("err is nil \n")
	} else {
		fmt.Printf("err is %v \n", err)
	}
}

func test2() {
	err := funcD()
	if err == nil {
		fmt.Printf("err is nil \n")
	} else {
		fmt.Printf("err is %v \n", err)
	}
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

	fmt.Println("execute a ..")
	fmt.Println("execute b ..")
	// program will exit when meet panic.
	// panic("panic here ..")
	// warming: unreachable code.
	fmt.Println("execute c ..")
	fmt.Println("execute d ..")

	test1()
	test2()

}
