/**
 * @Author : Yannick
 * @File   : learn_array.go
 * @Date   : 2017-11-24
 * @Desc   : This is a demo which I could learn golang channel through it.
 */

package main

import (
	"fmt"
)

func cf1(in chan int) {
	fmt.Println("start in cf1 ..")
	fmt.Println(<-in)
}

func cf2(msg *string, in chan int) {
	in <- 1
	fmt.Println("start in cf2 ..", *msg)
}

func LearnChannel() {

	fmt.Println()
	fmt.Println("## 无缓冲channel：example 1，看上去准备好接收者和发送者，同样会dead lock ##")

	// 注意：
	// 首先我们知道无缓冲channel，无论读写都会阻塞，直到有人准备好写入或者接收。但是也要分场景，如下
	// go cf1执行被阻塞，它阻塞的是函数cf1，此时跳出goroutine运行main()中的out <- 0，
	// 发现写入者，立即恢复阻塞的cf1
	out := make(chan int)
	go cf1(out)
	out <- 0
	fmt.Println("start outof cf1")

	// out <- 0执行被阻塞，它阻塞的是函数main()，此时跳出goroutine没有东西可运行
	// 看上去下一步go cf1(out)会准备好接收者，应该会立即恢复out <- 0并执行，但实际上go cf1(out)根本没执行main()就阻塞了
	// 下面的代码会panic deak lock

	// out := make(chan int)
	// out <- 0
	// go cf1(out)
	// fmt.Println("start outof cf1")

	fmt.Println()
	fmt.Println("## 无缓冲channel：example 2，基本演示 ##")

	chan1 := make(chan int)
	exit := make(chan bool)
	// Setp :
	// 	1) 对于无缓冲的channel，不管是写数据(channel<-data)，还是取数据（value<-channel），都会阻塞。
	//     即channel<-data阻塞，直到有人从channel中取数据；value<-channel也阻塞，直到有人向channel中写数据
	go func() {
		for {
			nTemp := <-chan1 // 写里写成chan1<-，用写来阻塞，下面改成读是一样的。
			fmt.Println(nTemp)

			if nTemp == 4 {
				exit <- true
				break
			}
		}
	}()
	chan1 <- 1 // 上面改成chan1<-，则这里改成<-chan1是一样的
	chan1 <- 2
	chan1 <- 3
	chan1 <- 4

	fmt.Println()
	fmt.Println("## 无缓冲channel：example 3，可以从关闭的channel取数据，不会阻塞，返回类型T的默认值，但是写入会dead lock ##")
	// exit已经被写了，但是没被取数据，再次写入会panic
	// exit <- false

	var bRet bool
	// 注意不能写成bRet <-exit
	bRet = <-exit
	fmt.Printf("bRet => %v\n", bRet)

	// 可以从关闭的channel中取数据，不会阻塞，取出来的数据为类型T的默认值（int =>0, bool => false)
	// 向一个关闭的channel写数据会panic
	close(chan1)
	var iRet int = <-chan1
	fmt.Printf("iRet => %d\n", iRet)

	fmt.Println()
	fmt.Println("## 无缓冲channel：example 4，阻塞函数的调用顺序：FILO（不同于从channel中取数据，channel中数据为FIFO ##")
	in1 := make(chan int)
	str1 := "hello "
	str2 := "world"
	// 如果将go cf2(&str1, in1)写成cf2(&str1, in1)，怎么样都是死锁
	// 输出如下，跟调用顺序是相反的，函数阻塞类似于defer调用
	// start in cf2 .. world
	// start in cf2 .. hello
	go cf2(&str1, in1)
	go cf2(&str2, in1)
	<-in1
	<-in1

	fmt.Println()
	fmt.Println("## 带缓冲channel：example 1，for range 或者for i, ok遍历时，怎么关闭channel ##")
	// 带缓冲的channel，在写满前是不阻塞的。
	chanRan := make(chan int, 4)
	go func() {
		// 注意这里range只返回一个参数
		for v := range chanRan {
			fmt.Println(v)
		}

		// for {
		// 	i, isClose := <-chanRan
		// 	if !isClose {
		// 		fmt.Printf("chanRan is closed!\n")
		// 		break // 这里必须break，否则当chanRan关闭的时候，这里死循环了
		// 	} else {
		// 		fmt.Printf("%d in chanRan \n", i)
		// 		// break // 这里如果break，则只会从chanRan中取一次数据，即5（满足FIFO）。
		// 	}
		// }

		exit <- true
	}()

	chanRan <- 5
	chanRan <- 6
	chanRan <- 7
	chanRan <- 8

	// 这里如果不关闭，则for循环中i, isClose := <-chanRan继续向一个空的chanRan取数据，会deadlock。
	// for range 同样如此。
	// 如果在生产者-消息费模式中，请一定要生产者处关闭channel，在消费者处关闭channel会引起panic
	close(chanRan)

	bRet = <-exit

	fmt.Println(" ##")
	var c = make(chan int, 1)

	go func() {
		c <- 'c'
		fmt.Println("在goroutine内")
	}()
	c <- 'c'
	fmt.Println("在goroutine外")

	// 	chanTimeOut := make(chan bool, 1)
	// 	chanTest := make(chan int, 1)
	// 	fmt.Println("下面是测试select相对于channel的用法 以及利用 select 来设置超时:")
	// 	go func() {
	// 		for i := 0; i < 3; i++ {
	// 			time.Sleep(1e9)
	// 			chanTimeOut <- true
	// 		}
	// 	}()

	// 	for i := 0; i < 3; i++ {
	// 		select {
	// 		case chanTest <- i:
	// 			fmt.Println("success write %d to channel", i)
	// 		case chanTest <- i + 1:
	// 			fmt.Println("success write %d to channel", i+1)
	// 		case chanTest <- i + 2:
	// 			fmt.Println("success write %d to channel", i+2)
	// 		case chanTest <- i + 3:
	// 			fmt.Println("success write %d to channel", i+3)
	// 		case chanTest <- i + 4:
	// 			fmt.Println("success write %d to channel", i+4)
	// 		case chanTest <- i + 5:
	// 			fmt.Println("success write %d to channel", i+5)
	// 		case ok := <-chanTimeOut:
	// 			fmt.Println("chanTimeOut is ", ok, "time out !")
	// 			//default:
	// 			//	fmt.Println("this channel is blocked")
	// 		}
	// 		//n := <-chanTest
	// 		//fmt.Println("this is %d :  value : %d", i, n)
	// 	}

	// 	for i := 0; i < len(chanTest); i++ {
	// 		fmt.Println(<-chanTest)
	// 	}

	// 	fmt.Println("下面是测试 chanel 的 range用法 :")
	// 	data := make(chan int, 3)
	// 	data <- 1
	// 	data <- 2
	// 	data <- 3
	// 	go func() {
	// 		for v := range data {
	// 			fmt.Println(v)
	// 		}
	// 		exit <- true
	// 	}()

	// 	data <- 4
	// 	data <- 5

	// 	close(data)
	// 	<-exit

	// 	fmt.Println("下面是判断 chanel 是否关闭的用法 :")
	// 	chanClose := make(chan string)

	// 	go func() {
	// 		for {
	// 			if v, ok := <-chanClose; ok {
	// 				fmt.Println(v)
	// 			} else {
	// 				fmt.Println("close this channel")
	// 				break
	// 			}

	// 		}
	// 		exit <- true
	// 	}()

	// 	time.Sleep(2 * 1e9)

	// 	chanClose <- "ready to close"
	// 	close(chanClose)

	// 	<-exit
	// 	fmt.Println("下面是定义单项 chanel 的用法 :")

	// 	channel := make(chan int)
	// 	var cSend chan<- int = channel
	// 	var cReveice <-chan int = channel
	// 	go func() {
	// 		nRet := <-cReveice
	// 		//cReveice <- 2   invalid operation: cReveice <- 2 (send to receive-only type <-chan int)
	// 		fmt.Println(nRet)
	// 		exit <- true
	// 	}()
	// 	cSend <- 1
	// 	//nRet := <-cSend   invalid operation: <-cSend (receive from send-only type chan<- int)
	// 	<-exit
	// 	fmt.Println("下面是多核分布式计算总数总和 :")

	// 	runtime.GOMAXPROCS(runtime.NumCPU())
	// 	fmt.Println("count is ", AllCount(13245678321, int64(runtime.NumCPU()), exit))

	// }
	// func SomeCount(nStart int64, nEnd int64, c chan int64) {

	// 	var nCount int64 = 0
	// 	for i := nStart; i < nEnd+1; i++ {
	// 		nCount = nCount + i
	// 	}
	// 	fmt.Println("start - ", nStart, "end - ", nEnd, ":count is ", nCount)
	// 	c <- nCount

	// }
	// func AllCount(nCount int64, nCPU int64, exit chan bool) int64 {
	// 	c := make(chan int64)
	// 	var i int64
	// 	for i = 0; i < nCPU; i++ {
	// 		go SomeCount((i*nCount)/nCPU, ((i+1)*nCount)/nCPU, c)
	// 	}
	// 	var nInt64 int64 = 0
	// 	for i = 0; i < nCPU; i++ {
	// 		nInt64 = nInt64 + <-c
	// 	}

	// return nInt64
}
