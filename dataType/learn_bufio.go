/**
 * @Author : Yannick
 * @File   : learn_bufio.go
 * @Date   : 2017-12-078
 * @Desc   : Learning and practicing the usage of the bufio library in GO.
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unsafe"
)

// bufio 包实现了带缓存的 I/O 操作
// 它封装一个 io.Reader 或 io.Writer 对象
// 使其具有缓存和一些文本读写功能

// type Reader struct {
// 	buf          []byte
// 	rd           io.Reader // reader provided by the client
// 	r, w         int       // buf read and write positions
// 	err          error
// 	lastByte     int
// 	lastRuneSize int
// }

func LearnBufio() {
	// NewReaderSize 将 rd 封装成一个拥有 size 大小缓存的 bufio.Reader 对象
	// 如果 rd 的基类型就是 bufio.Reader 类型，而且拥有足够的缓存
	// 则直接将 rd 转换为基类型并返回
	// func NewReaderSize(rd io.Reader, size int) *Reader
	fmt.Printf("----NewReaderSize : \n")
	str := "Good good study, day day up!"
	strReader := strings.NewReader(str)
	bufReader := bufio.NewReaderSize(strReader, 5)
	p := (*struct {
		buf          []byte
		rd           io.Reader
		r, w         int
		err          error
		lastByte     int
		lastRuneSize int
	})(unsafe.Pointer(bufReader))

	b, _ := bufReader.Peek(5) // Peek 返回缓存的一个切片，该切片引用缓存中前 n 字节数据

	// 返回Good，返回前5不错，注意Good后有一个空格
	fmt.Printf("b = bufReader.Peek(5) => %s\n", b)
	// 返回"Good good study,"即返回了16个字节，明明指定bufio.NewReaderSize(strReader, 5)缓冲大小为5？

	// 原因：bufio中定义了最小缓冲区大小const minReadBufferSize = 16， 不信打印出bufReader底层信息看下
	// buf:[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]，果然是16
	fmt.Printf("bufReader underlying info : %+v\n", p)
	b, _ = bufReader.Peek(len(str))

	// 返回：28 12
	// 问题：为什么strReader.Len()返回的不是len(str)即28？
	// 原因：strReader.Len()返回的是未len(r.s) - r.i的值，即字符串s中未被读取的字符长度，不信打印底层信息看看
	fmt.Printf("%d %d\n", len(str), strReader.Len())
	pSReader := (*struct {
		s       string
		i       int
		preRune int
	})(unsafe.Pointer(strReader))
	fmt.Printf("strReader underlying info : %+v\n", pSReader)
	fmt.Printf("b bufReader.Peek(len(str)) => %s\n", b)

}
