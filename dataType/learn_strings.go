/**
 * @Author : Yannick
 * @File   : learn_strings.go
 * @Date   : 2017-12-078
 * @Desc   : Learning and practicing the usage of the strings library in GO.
 */

package main

import (
	"bytes"
	"fmt"
	"strings"
	"unsafe"
)

func LearnStrings() {
	fmt.Println("----strings.NewReader :")
	var str = "Good good study, day day up!"
	// func NewReader(s string) *Reader { return &Reader{s, 0, -1} }
	// 通过字符串 s 创建 strings.Reader 对象
	// 类似于 bytes.NewBufferString, 但比 bytes.NewBufferString 更有效率，而且只读
	strReader := strings.NewReader(str)
	fmt.Printf("strReader Len => %d\n", strReader.Len())
	p := (*struct {
		s       string
		i       int
		preRune int
	})(unsafe.Pointer(strReader))
	fmt.Printf("strReader underlying info : %+v\n", p)

	fmt.Println("\n----strings.NewReader.Read :")
	// Read 将 r.i 之后的所有数据写入到 b 中（如果 b 的容量足够大）
	// 返回读取的字节数和读取过程中遇到的错误
	// 如果无可读数据，则返回 io.EOF
	// func (r *Reader) Read(b []byte) (n int, err error)
	buf1 := make([]byte, 5)
	for n, _ := strReader.Read(buf1); n > 0; n, _ = strReader.Read(buf1) {
		fmt.Printf("%q ", buf1[:n])
	}
	fmt.Printf("\nstrReader underlying info : %+v\n", p)

	fmt.Println("\n\n----strings.NewReader.ReadAt :")
	// ReadAt 将 off 之后的所有数据写入到 b 中（如果 b 的容量足够大）
	// 返回读取的字节数和读取过程中遇到的错误
	// 如果无可读数据，则返回 io.EOF
	// 如果数据被一次性读取完毕，则返回 io.EOF
	// func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
	buf2 := make([]byte, 23)
	n, err := strReader.ReadAt(buf2, 5)
	fmt.Printf("read %d bytes to buf2, data of buf2 => %q, err =>%v\n", n, buf2[:], err)
	fmt.Printf("strReader underlying info : %+v\n", p)

	fmt.Println("\n----strings.NewReader.ReadByte :")
	// ReadByte 将 r.i 之后的一个字节写入到返回值 b 中
	// 返回读取的字节和读取过程中遇到的错误
	// 如果无可读数据，则返回 io.EOF
	// func (r *Reader) ReadByte() (b byte, err error)
	for i := 0; i < 3; i++ {
		b, _ := strReader.ReadByte()
		fmt.Printf("%q, ", b) //strReader.Read(buf1)使用r.i变化了(28)，所以这里打印出来都是\x00
	}

	fmt.Println("\n\n----strings.NewReader.ReadByte :")
	fmt.Printf("strReader underlying info : %+v\n", p)

	// UnreadByte 撤消前一次的 ReadByte 操作，即 r.i--
	// func (r *Reader) UnreadByte() error
	fmt.Println("----strings.NewReader.UnreadByte :")
	for i := 0; i < 5; i++ {
		_ = strReader.UnreadByte()
	}
	fmt.Printf("strReader underlying info : %+v\n", p) // r.i将从28回退5，即r.i = 23

	fmt.Println("\n----strings.NewReader.ReadRune :")
	// ReadRune 将 r.i 之后的一个字符写入到返回值 ch 中
	// ch： 读取的字符
	// size：ch 的编码长度
	// err： 读取过程中遇到的错误
	// 如果无可读数据，则返回 io.EOF
	// 如果 r.i 之后不是一个合法的 UTF-8 字符编码，则返回 utf8.RuneError 字符
	// func (r *Reader) ReadRune() (ch rune, size int, err error)

	var str1 = "你好世界！(hello,world)" // 在GO中，一个汉字占用3个字节
	strReader1 := strings.NewReader(str1)
	// for i := 0; i < strReader1.Len()+10; i++ {
	for i := 0; i < 18; i++ {
		b, n, _ := strReader1.ReadRune()
		fmt.Printf(`"%c : %v", `, b, n)
	}
	p1 := (*struct {
		s       string
		i       int
		preRune int
	})(unsafe.Pointer(strReader1))
	fmt.Printf("\nstrReader underlying info : %+v\n", p1)

	fmt.Println("\n----strings.NewReader.seek :")
	// Seek 用来移动 r 中的索引位置
	// offset：要移动的偏移量，负数表示反向移动
	// whence：从那里开始移动，0：起始位置，1：当前位置，2：结尾位置
	// 如果 whence 不是 0、1、2，则返回错误信息
	// 如果目标索引位置超出字符串范围，则返回错误信息
	// 目标索引位置不能超出 1 << 31，否则返回错误信息
	// func (r *Reader) Seek(offset int64, whence int) (int64, error)
	fmt.Printf("origin : strReader underlying info : %+v\n", p1) // r.i = 28
	strReader1.Seek(-5, 1)                                       // 当前位置向后移动5，r.i应该为r.i = 23
	fmt.Printf("Seek(-5, 1) : strReader underlying info : %+v\n", p1)
	strReader1.Seek(3, 0) // 开始位置向前移动3，r.i应该为r.i = 3
	fmt.Printf("Seek(3, 0) : strReader underlying info : %+v\n", p1)

	fmt.Println("\n----strings.NewReader.WriteTo :")
	// WriteTo 将 r.i 之后的数据写入接口 w 中
	// func (r *Reader) WriteTo(w io.Writer) (n int64, err error)

	buf3 := bytes.NewBuffer(nil) // bytes.NewBuffer实现了io.Writer接口
	strReader1.WriteTo(buf3)
	fmt.Printf("%q\n", buf3)
	strReader1.Seek(0, 0)
	buf3.Reset() // 清空buf3
	strReader1.WriteTo(buf3)
	fmt.Printf("after seek(0,0) %q\n", buf3)

	fmt.Println("\n----strings.NewReplacer :")
	// NewReplacer 通过“替换列表”创建一个 Replacer 对象。
	// 按照“替换列表”中的顺序进行替换，只替换非重叠部分。
	// 如果参数的个数不是偶数，则抛出异常。
	// 如果在“替换列表”中有相同的“查找项”，则后面重复的“查找项”会被忽略
	// func NewReplacer(oldnew ...string) *Replacer

	srp := strings.NewReplacer("a", "A", "x", "xx", ".", "-") //将a替换成A，x替换成xx, "."替换成"-"
	s4 := "bar.sexy banana.six "
	rst := srp.Replace(s4)
	fmt.Printf("after replace, rst => %s, s4 => %s\n", rst, s4) // 返回新的字符串，原来的字符串不会被修改
	rst1 := strings.Replace(s4, "a", "A", 1)                    // 跟NewReplacer相比，一次只能替换一个，n为替换的个数，<0则替换全部
	fmt.Printf("after replace, rst1 => %s, s4 => %s\n", rst1, s4)
}
