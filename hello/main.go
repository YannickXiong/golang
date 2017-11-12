/*
 * @Author: Yannick
 * @Date: 2017-11-12 15:32:15
 * @Last Modified by: Yannick
 * @Last Modified time: 2017-11-12 21:15:34
 */

package main

import (
	"fmt"
	"learn_golang/hello/test1"
	"learn_golang/hello/test2"
	"learn_golang/simpleSet"
)

func main() {
	fmt.Println("Hello world!")
	test2.T2()
	test1.T1()
	s := set.New()
	s.Add(1)
	s.Add("xyang")
	s1 := set.New()
	s1.Add(1)
	s1.Add("xyang")
	fmt.Printf("simpleSet Info: => %v\n", s.Element)
	fmt.Printf("s1 is the same s?: => %v\n", s.Same(s1))
	s1.Add(7.34)
	s1.Add(true)
	s.Clear()
	fmt.Printf("after Clear() :%v\n", s)
	snap := s1.Snapshot()
	fmt.Printf("snap => %v\n", snap)
	//fmt.Println(s1.String())
	// both %s and %v call called s1.String()
	fmt.Printf("%s =>, %v\n", s1, s1)
	set.Destroy(s1)
	fmt.Printf("before => %p\n", s.Element)
	s.Clear()
	fmt.Printf("after => %p\n", s.Element)
	fmt.Printf("%v => %d => %v\n", s.Element, s.Len(), s.IsEmpty())
	fmt.Printf("%v, %v", s1, nil == s1)

}
