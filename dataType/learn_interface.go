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
	Car        // anonymous ttribute
	seatNum    int
	wheelsSize float32
}

// VWCar : has price, carName, seatNum and wheelsSize
type VWCar struct {
	Car     // anonymous ttribute
	seatNum int
	color   string
}

// CarInterface : has 1 method
type CarInterface interface {
	ShowPrice()
}

// AudiInterface : has 4 methods
type AudiInterface interface {
	ShowColor()
	ShowPrice()
	ShowDiscountPrice(discount float32)
	ShowWheelsSize()
}

// BenZInterface : has 4 methods
type BenZInterface interface {
	ShowSeatNum()
	ShowPrice()
	ShowDiscountPrice(discount float32)
	ShowWheelsSize()
}

// VWInterface : has 4 methods
type VWInterface interface {
	ShowColor()
	// ShowPrice()
	ShowDiscountPrice(discount float32)
	ShowSeatNum()
}

// ShowPrice : type Car implements ShowPrice
func (car *Car) ShowPrice() {
	fmt.Printf("Car::ShowPrice => %.2f\n", car.price)
}

// ShowColor : type AudiCar implements ShowColor
func (audiCar *AudiCar) ShowColor() {
	fmt.Printf("AudiCar::color => %s\n", audiCar.color)
}

// ShowPrice : type AudiCar implements ShowPrice
// Also overwirte Car.ShowPrice. Attention that in go func cannot be overwritten.
func (audiCar *AudiCar) ShowPrice() {
	// audiCar.Car.price == audiCar.price
	// fmt.Printf("AudiCar::price => %.2f\n", audiCar.Car.price)
	fmt.Printf("AudiCar::price => %.2f\n", audiCar.price)
}

// ShowDiscountPrice : type AudiCar implements ShowDiscountPrice
func (audiCar *AudiCar) ShowDiscountPrice(discount float32) {
	//fmt.Printf("AudiCar::discount Price => %.2f\n", audiCar.Car.price*discount)
	fmt.Printf("AudiCar::discount Price => %.2f(discount:%.2f)\n", audiCar.price*discount, discount)
}

// ShowWheelsSize : type AudiCar implements ShowWheelsSize
func (audiCar *AudiCar) ShowWheelsSize() {
	fmt.Printf("AudiCar::wheelsSize => %.1f\n", audiCar.wheelsSize)
}

// ShowSeatNum : type BenZCar implements ShowSeatNum
func (benzCar *BenZCar) ShowSeatNum() {
	fmt.Printf("BenZCar::seatNum => %d\n", benzCar.seatNum)
}

// // BenZInterface defined interface ShowPrice, while type BenZCar does not implement it.
// // So it will usd the Car.ShowPrice or not?
// // The answer is Yes! benzCar.ShowPrice() => result: Car::ShowPrice => 688888.00
// // ShowPrice : type BenZCar implements ShowPrice
// func (benzCar *BenZCar)ShowPrice(){
// 	fmt.Printf("BenZCar::price => %.2f\n", benzCar.price)
// }

// ShowDiscountPrice : type AudiCar implements ShowDiscountPrice
func (benzCar *BenZCar) ShowDiscountPrice(discount float32) {
	//fmt.Printf("BenZCar::discount Price => %.2f\n", BenZCar.Car.price*discount)
	fmt.Printf("BenZCar::discount Price => %.2f(discount:%.2f)\n", benzCar.price*discount, discount)
}

// ShowWheelsSize : type AudiCar implements ShowWheelsSize
func (benzCar *BenZCar) ShowWheelsSize() {
	fmt.Printf("BenZCar::wheelsSize => %.1f\n", benzCar.wheelsSize)
}

// ShowColor : type VWCar implements ShowColor
func (vwCar VWCar) ShowColor() {
	fmt.Printf("VWCar::color => %s\n", vwCar.color)
}

// ShowPrice : type VWCar implements ShowPrice
// will leat to warming(not error): receiver name vmCar should be consistent with
// previous receiver name vwCar for VWCar
// Reason: does not define ShowPrice in VMCarInterface.
// Another question: VWInterface does not define interface ShowPrice, and type VWCar
// does not implement it.So it will usd the Car.ShowPrice or not?
// The answer is Yes! vmCar.ShowPrice() => result: Car::ShowPrice => 151788.00
// func (vwCar VWCar) ShowPrice() {
// 	fmt.Printf("VWCar::price => %.2f\n", vwCar.price)
// }

// ShowDiscountPrice : type VMCar implements ShowDiscountPrice
func (vwCar VWCar) ShowDiscountPrice(discount float32) {
	//fmt.Printf("VWCar::discount Price => %.2f\n", vwCar.Car.price*discount)
	fmt.Printf("VWCar::discount Price => %.2f(discount:%.2f)\n", vwCar.price*discount, discount)
}

// ShowSeatNum : type VMCar implements ShowSeatNum
func (vwCar VWCar) ShowSeatNum() {
	fmt.Printf("VWCar::seatNum => %d\n", vwCar.seatNum)
}

// LearnInterface :
func LearnInterface() {
	fmt.Println("## LearnInterface() called begin ..")

	// the fellow way of init audiCar will lead compile error.
	// audiCar := AudiCar{789998.00, "Audo Q7-Vx", 4, "White"}
	fmt.Println("***** AudiCar Info ***** ")
	audiCar := &AudiCar{Car{786999.00, "Audi Q7 VVx-honor"}, 17.1, "White"}
	audiCar.ShowPrice()
	audiCar.ShowDiscountPrice(0.88)
	audiCar.ShowWheelsSize()
	audiCar.ShowColor()

	fmt.Println("***** BenZCar Info ***** ")
	// benzCar is value type, not reference of BenZCar{Car{688888.00, "BenZ GLK-4000"}, 6, 15.8}
	// however, it can be used benzCar.ShowPrice() whose receiver is a reference type.
	// Attention that normal function can not be used like this.
	// About this, details see learn_method.go
	benzCar := &BenZCar{Car{688888.00, "BenZ GLK-4000"}, 6, 15.8}
	benzCar.ShowPrice()
	benzCar.ShowDiscountPrice(0.68)
	benzCar.ShowWheelsSize()
	benzCar.ShowSeatNum()

	fmt.Println("***** VWCar Info ***** ")
	vwCar := &VWCar{Car{151788.00, "Golf-1.2T"}, 5, "Red"}
	vwCar.ShowPrice()
	vwCar.ShowDiscountPrice(0.95)
	vwCar.ShowSeatNum()
	vwCar.ShowColor()

	// Up to now, you think you've used the GO language interface, my answer is no, you just used method.
	// You look back at your definition of AudiInterface, BenZInterface, VWInterface, do you use it?
	// In GO language, if you define a variable of interface, then the variable can store
	// any type of object that implements the interface

	var ai AudiInterface // ai is AudiInterface type
	// var bi *BenZInterface // bi is *BenZInterface
	// bi will lead to compile error" bi.ShowPrice undefined, type *BenZInterface is pointer to interface, not interface)
	// Explains why there's never a pointer to interface
	var bi BenZInterface // bi is BenZInterface
	var vi VWInterface   // vi is VWInterface

	// audiCar implemented AudiInterface, so ai can story audiCar
	ai = audiCar
	fmt.Println("***** ai AudiCar Info ***** ")
	ai.ShowColor()
	ai.ShowPrice()
	ai.ShowDiscountPrice(0.68)
	ai.ShowWheelsSize()

	// benzCar implemented BenzInterface, so bi can story benzCar
	bi = benzCar
	fmt.Println("***** bi BenZCar Info ***** ")
	bi.ShowPrice()
	bi.ShowDiscountPrice(0.68)
	bi.ShowWheelsSize()
	bi.ShowSeatNum()

	vi = vwCar
	fmt.Println("***** vi VWCar Info ***** ")
	vi.ShowColor()
	vi.ShowDiscountPrice(0.95)
	vi.ShowSeatNum()
	// error: vi.ShowPrice undefined (type VWInterface has no field or method ShowPrice)
	// vi is a VWInterface, has no ShowPrice() method. While vwCar is VWCar type, has ShowPrice from anonymous member Car,
	// is it a bit like the inheritance concept in PYTHON?
	// vi.ShowPrice()

	// ci is a CarInterface, and all audiCar, benzCar, vwCar implemented CarInterface.
	// So ci can store all of them.
	ci := make([]CarInterface, 3)
	ci[0], ci[1], ci[2] = audiCar, benzCar, vwCar
	fmt.Println("***** CarInterface::audiCar Info ***** ")
	// audiCar.ShowPrice()
	ci[0].ShowPrice()
	fmt.Println("***** CarInterface::benzCar Info ***** ")
	// BenZCar did not implement ShowPrice, call Car.ShowPrice()
	ci[1].ShowPrice()
	fmt.Println("***** CarInterface::vwCar Info ***** ")
	// VWCar did not implement ShowPrice, call Car.ShowPrice()
	ci[2].ShowPrice()

	// empty interface can store any type!
	type EmptyInterface interface{}

	// var ei, epi EmptyInterface = audiCar, &audiCar
	// which will lead error: interface conversion: main.EmptyInterface is **main.AudiCar, not *main.AudiCar
	// think why?

	var ei, epi EmptyInterface = *audiCar, &*audiCar
	fmt.Println("***** Empty Interface::AudiCar Info ***** ")
	// ei can store AudiCar, but cannot call ei.ShowColor(),
	// ei.ShowColor undefined (type EmptyInterface is interface with no methods)
	//ei.ShowColor()
	// ei.name = "Audi Q3" : error, ei.name undefined (type EmptyInterface is interface with no methods)
	// ei.(audiCar).carName = "Audi Q3" : error, audiCar is not a type
	// ei.(AudiCar).carName = "Audi Q3" : error  cannot assign to ei.(AudiCar).Car.carName
	// OK, because interface type returned a temporary object, only the pointer can modify its state
	epi.(*AudiCar).carName = "Audi Q3"
	fmt.Printf("ei.(AudiCar).carName=>%s\n", ei.(AudiCar).carName)
	fmt.Printf("epi.(*AudiCar).carName=>%s\n", epi.(*AudiCar).carName)

	fmt.Println("## LearnInterface() called end ..")
}
