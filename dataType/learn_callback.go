/**
 * @Author : Yannick
 * @File   : learn_calback.go
 * @Date   : 2017-11-28
 * @Desc   : This is a demo which I could learn golang call back function through it.
 */

package main

import (
	"fmt"
	"time"
)

// 回调函数：一个通过指针调用的函数就叫回调函数，回调函数通常作为另外一个函数的参数来使用。
// 加深理解：回调函数就是一个通过函数指针调用的函数。如果你把函数的指针（地址）作为参数传递给另一个函数，
//          当这个指针被用来调用其所指向的函数时，我们就说这是回调函数。回调函数不是由该函数的实现方直接调用，
//			而是在特定的事件或条件发生时由另外的一方调用的，用于对该事件或条件进行响应
// 存在意义：抽象画了一个函数，简洁代码。异步执行。

// iSOdd ：
func iSOdd(in int) bool {
	if in%2 != 0 {
		return true
	}

	return false
}

// iSEven ：
func iSEven(in int) bool {
	if in%2 == 0 {
		return true
	}

	return false
}

type filterHandle func(int) bool

// Filter : 一个参数是f func(in int) bool，这就是个回调函数
// 具体实现由上面iSOdd和iSEven实现
// 也可以直接定义func Filter(s []int, f func(int) bool) []int {}
func Filter(s []int, f filterHandle) []int {
	var result []int

	for _, value := range s {
		if f(value) {
			result = append(result, value)
		}
	}

	return result
}

// Call的基本定义，对外部使用者的请求、返回以及异步使用进行封装。
type Call struct {
	Request interface{}
	Reply   interface{}
	Done    chan *Call //用于结果返回时,指向自己的指针
}

// 非常重要的异步调用结果返回，供框架内部使用。
func (call *Call) done() {
	select {
	case call.Done <- call:
	// ok
	default:
		// 阻塞情况处理,这里忽略
	}
}

// 供业务调用的异步计算函数封装，用户只需要了解对应参数。
func GO(req int, reply *int, done chan *Call) *Call {
	if done == nil {
		done = make(chan *Call, 10)
	} else {
		if cap(done) == 0 {
			fmt.Println("chan容量为0,无法返回结果,退出此次计算!")
			return nil
		}
	}
	call := &Call{
		Request: req,
		Reply:   reply,
		Done:    done,
	}
	//调用一个可能比较耗时的计算，注意用"go"
	go caculate(call)
	return call
}

//真正的业务处理代码
//简单示意,其实存在读写竞争。run -race 就会出现提示
func caculate(call *Call) {
	//假定运算一次需要耗时1秒
	time.Sleep(time.Second)
	tmp := call.Request.(int) * 5
	call.Reply = &tmp
	call.done()
}

// LearnCallBack :
func LearnCallBack() {
	fmt.Println("----回调函数基本例子----")
	s := []int{-9, -2, 0, 1, 2, 3, 4, 5, 7, 8, 9, 10, 11}
	oddSlice := Filter(s, iSOdd) // 函数当做值来传递了
	fmt.Printf("oddSlice is %v \n", oddSlice)
	evenSlice := Filter(s, iSEven) // 函数当做值来传递了
	fmt.Printf("evenSlice is %v \n", evenSlice)

	fmt.Println("----回调函数第二个例子，结合信道通知事件完成----")
	for i := 0; i < 100; i++ {
		var reply *int
		call := GO(i, reply, nil) //获取到了call，但此时call.Reply还不是运算结果
		//先打印结果还没有计算出来的情况
		fmt.Printf("i=%d,运算前：call.Reply=%v \n", i, call.Reply.(*int))

		result := <-call.Done //等待Done的通知，此时call.Reply发生了变化。
		fmt.Printf("i=%d,运算后：call.Reply=%v,result=%+v \n", i, *(call.Reply.(*int)), *(result.Reply.(*int)))
	}
}
