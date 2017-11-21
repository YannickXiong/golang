/**
 * @Author : Yannick
 * @File   : learn_special.go
 * @Date   : 2017-11-21
 * @Desc   : some special useage.
 */

package main

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"strconv"
)

type AttributeType uint8

const (
	//start rfc2865
	_                          = iota //drop the zero
	UserName     AttributeType = iota //1
	UserPassword AttributeType = iota //2
	ServiceType  AttributeType = 6
)

type AVP struct {
	Type  AttributeType
	Value []byte
}

type Packet struct {
	Secret        string
	Identifier    uint8
	Authenticator [16]byte
	AVPs          []AVP
}

type avpDataType interface {
	Value(p *Packet, a AVP) interface{}
	String(p *Packet, a AVP) string
}

type attributeTypeDesc struct {
	name     string
	dataType avpDataType
}

type avpUint32Enum struct {
	t interface{} // t should from a uint32 type like AcctStatusTypeEnum
}

var avpString avpStringt

type avpStringt struct{}

type ServiceTypeEnum uint32

var attributeTypeMap = [7]attributeTypeDesc{
	UserName:    {"UserName", avpString},
	ServiceType: {"ServiceType", avpUint32Enum{ServiceTypeEnum(0)}},
}

// avpUint32Enum implements method Value of interface avpDataType
func (s avpUint32Enum) Value(p *Packet, a AVP) interface{} {
	value := reflect.New(reflect.TypeOf(s.t)).Elem()
	value.SetUint(uint64(binary.BigEndian.Uint32(a.Value)))
	return value.Interface()
}

// avpUint32Enum implements String Value of interface avpDataType
func (s avpUint32Enum) String(p *Packet, a AVP) string {
	number := binary.BigEndian.Uint32(a.Value)
	value := reflect.New(reflect.TypeOf(s.t)).Elem()
	value.SetUint(uint64(number))
	method := value.MethodByName("String")
	if !method.IsValid() {
		return strconv.Itoa(int(number))
	}
	out := method.Call(nil)
	return out[0].Interface().(string)
}

// avpStringt implements method Value of interface avpDataType
func (s avpStringt) Value(p *Packet, a AVP) interface{} {
	return string(a.Value)
}

// avpStringt implements String Value of interface avpDataType
func (s avpStringt) String(p *Packet, a AVP) string {
	return string(a.Value)
}

func LearnSpecial() {
	fmt.Printf("attributeTypeMap => %v\n", attributeTypeMap)
}
