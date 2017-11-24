/**
 * @Author : Yannick
 * @File   : learn_concurrent.go
 * @Date   : 2017-11-23
 * @Desc   : This is a demo which I could learn concurrent(syc, go, channel) in golang through it.
 */

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Sum(list ...int) {
	total := 0
	for _, num := range list {
		total += num
	}
	fmt.Printf("Sum of %v => %d\n", list, total)
	wg.Done() // 从队列中移除一个任务，相当于wg.Add(-1)
}

func LearnConcurrent() {
	a := []int{1, 2, 3, 4, 5, 6}
	// Sum(a) // error: cannot use a (type []int) as type int in argument to Sum
	// Sum(a...) // Sum(a...) is ok

	for i := 0; i < 10; i++ {
		wg.Add(1) // 增加一个任务到队列
		fmt.Printf("The %d job runs ..\n", i+1)
		// a = append(a, a...)

		// 思考如果a := append(a, i)，会怎么样?
		a = append(a, i)
		go Sum(a...)
		// wg.Done() // 如果Sum中的wg.Done()在这里，则Sum中的打印不会出现。只打印了fmt.Printf("The %d job runs ..", i+1)
	}
	wg.Wait() //阻塞，直接列队中任务完成
	fmt.Println("All jobs done !")
}
