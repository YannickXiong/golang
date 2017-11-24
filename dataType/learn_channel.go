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

func LearnChannel() {

	chan1 := make(chan int)
	exit := make(chan bool)

	//

	go func() {
		for {
			nTemp := <-chan1
			fmt.Println(nTemp)

			if nTemp == 4 {
				exit <- true
				break
			}
		}
	}()
	chan1 <- 1
	chan1 <- 2
	chan1 <- 3
	chan1 <- 4

	var bRet bool = <-exit
	fmt.Println(bRet)
	close(chan1)

	// 	chanRan := make(chan int, 4)
	// 	go func() {
	// 		for v := range chanRan {
	// 			fmt.Println(v)
	// 		}

	// 		exit <- true
	// 	}()

	// 	chanRan <- 5
	// 	chanRan <- 6
	// 	chanRan <- 7
	// 	chanRan <- 8

	// 	close(chanRan)

	// 	bRet = <-exit

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
