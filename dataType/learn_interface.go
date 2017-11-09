/**
 * @Author : Yannick
 * @File   : learn_interface.go
 * @Date   : 2017-11-09
 * @Desc   : This is a demo which I could learn golang interface through it.
 */

package main

import (
	"fmt"
)

/**
* Basic concepts of function, method, interface.
* function:  func foo(in string) (out int) {}
* method:    func (p *Perple) foo(in string) (out int){}
*            compare to function, method has a receiver more than a function.a
* interface: a set of method, for example as below
*            type Reader interface{
*				  method1
*				  method2
*				  methodn
*            }
 */

// Car : has price and carName
type Car struct {
	price   float32
	carName string
}

// AudiCar : has price, carName, wheelsSize and color
type AudiCar struct {
	Car        // anonymous ttribute
	wheelsSize float32
	color      string
}

// BenZCar : has price, carName, seatNum and color
type BenZCar struct {
	Car     // anonymous ttribute
	seatNum int
	color   string
}

// VWCar : has price, carName, seatNum and wheelsSize
type VWCar struct {
	Car        // anonymous ttribute
	seatNum    int
	wheelsSize float32
}

// AudiInterface : has 4 methods
type AudiInterface interface {
	ShowColor()
	GetPrice()
	GetDiscountPrice(discount float32)
	ShowWheelsSize()
}

// BenZInterface : has 4 methods
type BenZInterface interface {
	ShowSeatNum()
	GetPrice()
	GetDiscountPrice(discount float32)
	ShowWheelsSize()
}

// VWInterface : has 4 methods
type VWInterface interface {
	ShowColor()
	GetPrice()
	GetDiscountPrice(discount float32)
	ShowWheelsSize()
}

// LearnInterface :
func LearnInterface() {
	fmt.Println("## LearnInterface() called begin ..")
	fmt.Println("## LearnInterface() called end ..")
}
