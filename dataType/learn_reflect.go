/**
 * @Author : Yannick
 * @File   : learn_reflect.go
 * @Date   : 2017-11-21
 * @Desc   : learn usAges of reflect.
 */

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type myInt int
type myString string

// Foo :
type Foo struct {
	Name myString
	Age  myInt
}

func (foo Foo) String() string {
	return fmt.Sprintf("Info Name : %s \t Age : %d\n", foo.Name, foo.Age)
}

// String2 :
func (foo *Foo) String2() string {
	return fmt.Sprintf("Name : %s \t Age : %d\n", foo.Name, foo.Age)
}

type DB struct {
	Host string
	Port uint32
	Pw   string
}

type Conf struct {
	OpName *string `json:"jsonOpName" xml:"xmlOpName"`
	Num    *int
	Output *string
	Input  *string
	DB     // an anonymous struct
}

func (conf *Conf) ShowOpName(subname string) string {
	return *(conf.OpName) + " " + subname
}

func (conf *Conf) GetDBConf() DB {
	return conf.DB
}

// 场景1：遍历结构体字段名
func PrintStructField(t *reflect.Type) {
	// can not : ieldnum := (*t).NumField()
	// t.NumField undefined (type *reflect.Type is pointer to interface, not interface)
	fieldnum := (*t).NumField()
	for i := 0; i < fieldnum; i++ {
		fmt.Printf("conf.%v => %v\n", (*t).Field(i).Name, (*t).Field(i).Type)
	}
	fmt.Println()
}

// 场景2：调用结构体方法
func CallMethod(v *reflect.Value, method string, params []interface{}) {
	f := (*v).MethodByName(method)
	fmt.Printf("CallMethod :: params => %v\n", params)
	if f.IsValid() {
		args := make([]reflect.Value, len(params))
		for key, param := range params {
			fmt.Printf("CallMethod :: key => %v, param => %v \n", key, param)
			args[key] = reflect.ValueOf(param)
		}

		// 调用
		ret := f.Call(args)
		fmt.Printf("ret => %v\n", ret)
		if ret[0].Kind() == reflect.String {
			fmt.Printf("%s Called result: %s\n", method, ret[0].String())
		} else {
			fmt.Println("can't call " + method)
		}
	}
	fmt.Println()
}

// 场景3：获取结构体的tag标记
func GetTag(t *reflect.Type, field string, tagName string) error {
	var tagVal string
	var err error

	fieldVal, ok := (*t).FieldByName(field)
	if ok {
		tagVal = fieldVal.Tag.Get(tagName)
		fmt.Printf("get struct[%s] tag[%s]: %s, error:%v\n", field, tagName, tagVal, err)
	} else {
		err = errors.New("no such field name : " + field)
		fmt.Printf("get struct [%s] tag [%s] : %s, error : %v\n", field, tagName, tagVal, err)
		return err
	}
	return nil
}

// 场景4：用字符串反射实例化结构体
type RegStructMap map[string]interface{}

func InitStructByName(name string, regStruct RegStructMap) {
	if regStruct[name] != nil {
		// t := reflect.ValueOf(regStruct[name]).Type()  // also OK
		t := reflect.TypeOf(regStruct[name])
		// func New(typ Type) Value, returns a Value representing a pointer to a new zero value for the specified type
		v := reflect.New(t).Elem()
		fmt.Printf("Init struct :: %#v\n", v)
	}
}

// LearnReflect :
func LearnReflect() {
	var a myInt
	a = 1000
	fmt.Println("a := 1000")
	fmt.Printf("reflect.TypeOf(a) => %v\n", reflect.TypeOf(a))
	fmt.Printf("reflect.ValueOf(a) => %v\n", reflect.ValueOf(a))
	fmt.Printf("reflect.TypeOf(a).String() => %v\n", reflect.TypeOf(a).String())
	fmt.Printf("reflect.TypeOf(a).Name() => %v\n", reflect.TypeOf(a).Name())

	fmt.Println("reflect.TypeOf-------------------------------------------")

	b := "hello word"
	fmt.Println("b := \"hello word\"")
	// reflect.TypeOf() returns reflect.Type
	fmt.Printf("reflect.TypeOf(b) => %v\n", reflect.TypeOf(b))
	// reflect.ValueOf() returns reflect.Value
	fmt.Printf("reflect.ValueOf(b) => %v\n", reflect.ValueOf(b))
	// reflect.TypeOf().Name() returns  reflect.Type or empty string when unmamed.
	fmt.Printf("reflect.TypeOf(b).String() => %v\n", reflect.TypeOf(b).String())
	// the same to reflect.TypeOf().String()
	fmt.Printf("reflect.TypeOf(b).Name() => %v\n", reflect.TypeOf(b).Name())

	fmt.Println("reflect.TypeOf-------------------------------------------")

	foo := Foo{Name: "xyang", Age: 31}
	fmt.Println("foo := Foo{Name: \"xyang\", Age: 31}")
	fmt.Printf("reflect.TypeOf(foo) => %v\n", reflect.TypeOf(foo))
	fmt.Printf("reflect.ValueOf(foo) => %v\n", reflect.ValueOf(foo))
	fmt.Printf("reflect.TypeOf(foo).String() => %v\n", reflect.TypeOf(foo).String())
	fmt.Printf("reflect.TypeOf(foo).Name() => %v\n", reflect.TypeOf(foo).Name())

	fmt.Println("reflect.TypeOf.Field-------------------------------------------")

	// the filed domain
	typ := reflect.TypeOf(foo)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fmt.Printf("%s type is : %s\n", field.Name, field.Type)
	}
	// FieldByIndex needs []int arguments.
	// FieldByIndex returns : StructField(definition see reflect packAge)
	NameStruct := typ.FieldByIndex([]int{0})
	fmt.Printf("foo::reflect.TypeOf(foo).FieldByIndex(0) => %v\n", NameStruct)

	NameStruct1 := typ.Field(0)
	fmt.Printf("foo::reflect.TypeOf(foo).Field(0) => %v\n", NameStruct1)

	// FieldByName returns : StructField, bool
	NameStruct2, info2 := typ.FieldByName("Name")
	fmt.Printf("foo::reflect.TypeOf(foo).FieldByName(Name) => %v, info2 => %v\n", NameStruct2, info2)

	fmt.Println("reflect.TypeOf.Method-------------------------------------------")

	// the method domain
	// NumMethod : returns the number of exported methods in the value's method set
	// Attention : func (foo *Foo) String2() is not the method of foo.
	fmt.Printf("foo::reflect.TypeOf(foo).NumMethod() = > %d\n", typ.NumMethod())
	// typ.Method(0).Name is String, not String2.
	fmt.Printf("foo::reflect.TypeOf(foo).Method().Name = > %v\n", typ.Method(0).Name)
	fmt.Printf("foo::reflect.TypeOf(foo).Method().Type = > %v\n", typ.Method(0).Type)
	fmt.Printf("foo::reflect.TypeOf(foo).Method().Func = > %v\n", typ.Method(0).Func)

	fmt.Println("reflect.TypeOf.Kind-------------------------------------------")

	// the kind domain
	// both TypeOf and Valueof have method Kind()
	// Kind() returns the basic type(the underlying type), while TypeOf returns the statis type.
	fmt.Printf("reflect.TypeOf(a) = > %v\n", reflect.TypeOf(a))                 // main.myInt
	fmt.Printf("reflect.TypeOf(a).Kind() = > %v\n", reflect.TypeOf(a).Kind())   //int
	fmt.Printf("reflect.ValueOf(a).Kind() = > %v\n", reflect.ValueOf(a).Kind()) // int

	fmt.Println("reflect.ValueOf-------------------------------------------")
	// reflect.ValueOf(i interface{}) Value
	// reflect.ValueOf()的返回值类型为reflect.Value,它实现了interface{}参数到reflect.Value的反射
	//              reflect.ValueOf()                        Interface()
	// interface{} ---------------------> reflect.Value -------------------> interface{}
	fmt.Printf("reflect.ValueOf(a) => %v\n", reflect.ValueOf(a))
	fmt.Printf("reflect.ValueOf(a).Interface() => %v\n", reflect.ValueOf(a).Interface())
	// reflect.ValueOf(a).Interface().(int64)) 使用断言获取底层数据值
	// panic: interface conversion: interface {} is main.myInt, not int64
	// fmt.Printf("reflect.ValueOf(a).Interface().(Int64) => %v\n", reflect.ValueOf(a).Interface().(int64))
	fmt.Printf("reflect.ValueOf(a).Interface().(Int64) => %v\n", reflect.ValueOf(a).Interface().(myInt))
	fmt.Printf("reflect.ValueOf(b) => %v\n", reflect.ValueOf(b))
	fmt.Printf("reflect.ValueOf(b).Interface() => %v\n", reflect.ValueOf(b).Interface())
	fmt.Printf("reflect.ValueOf(b).Interface().(string) => %v\n", reflect.ValueOf(b).Interface().(string))
	fmt.Printf("reflect.ValueOf(foo) => %v\n", reflect.ValueOf(foo))
	fmt.Printf("reflect.ValueOf(foo).Interface() => %v\n", reflect.ValueOf(foo).Interface())
	fmt.Printf("reflect.ValueOf(foo).Interface().(Foo) => %v\n", reflect.ValueOf(foo).Interface().(Foo))

	fmt.Println("reflect.ValueOf.Field-------------------------------------------")

	// 和Type的Filed方法不一样，Type.Field()返回的是StructFiled对象，有Name,Type等属性，
	// Value.Field()返回的还是一个Value对象
	// panic, a Myint has no Field.
	// func (v Value) Field(i int) Value
	// Field returns the i'th field of the struct v. It panics if v's Kind is not Struct or i is out of range.
	// fmt.Printf("reflect.ValueOf(a).Field = > %v\n", reflect.ValueOf(a).Field(0))
	fmt.Printf("reflect.ValueOf(foo).Field(0) = > %v\n", reflect.ValueOf(foo).Field(0))
	fmt.Printf("reflect.ValueOf(foo).Field(1) = > %v\n", reflect.ValueOf(foo).Field(1))
	fmt.Printf("reflect.ValueOf(foo).FieldByName(\"Name\") = > %v\n", reflect.ValueOf(foo).FieldByName("Name"))
	fmt.Printf("reflect.ValueOf(foo).Field(\"Age\") = > %v\n", reflect.ValueOf(foo).FieldByName("Age"))

	rv := reflect.ValueOf(foo)
	rt := reflect.TypeOf(foo)
	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Field(i)
		ft := rt.FieldByIndex([]int{i}) // the same to ft:=rt.Field(i)
		// reflect.Value.Interface: cannot return value obtained from unexported field or method
		// reason: Name & Age are private in Foo
		// reflect can not get a private value.
		// chage Foo.name -> Foo.Name, Foo.age -> Foo.Age
		fmt.Printf("reflect.ValueOf().Field(%d).IsValid() => %v\n", i, fv.IsValid())
		fmt.Printf("Foo.%s static type is : %s, underlying type is %v, value is %v\n", ft.Name, fv.Type(), fv.Kind(), fv.Interface())
		// 如果这里用ft.Kind()就会报错，原因是：
		// reflect.ValueOf.Field()返回的是一个reflect.Value，是有Kind方法的
		// reflect.TypeOf.Field()返回的是一个reflect.Type结果体，没有Kind方法
		// fmt.Printf("Foo.%s static type is : %s, underlying type is %v, value is %v\n", ft.Name, ft.Type(), fv.Kind(), fv.Interface())
	}

	fmt.Println("reflect.ValueOf.SetxxxValue-------------------------------------------")
	fmt.Printf("a := 1000, reflect.ValueOf(a).CanSet() => %v\n", reflect.ValueOf(a).CanSet())
	fmt.Printf("a := 1000, reflect.ValueOf(&a).CanSet() => %v\n", reflect.ValueOf(&a).CanSet())
	fmt.Printf("b := \"hello word\", reflect.ValueOf(b).CanSet() => %v\n", reflect.ValueOf(b).CanSet())
	fmt.Printf("foo is a struct, reflect.ValueOf(foo).CanSet() => %v\n", reflect.ValueOf(foo).CanSet())

	// 以上全部是false，说明是不能对它们进行写值的。就算是reflect.ValueOf(&a)，也不能写值。
	// relect.Value是字符s的一个反射对象，是不能直接对它进行赋值操作的。
	// 要对a进行赋值，需要先拿到a的指针对应的reflect.Value, 然后通过Value.Elem()再对应到a，然后才能赋值操作。
	// Elem returns the value that the interface v contains or that the pointer v points to.
	// It panics if v's Kind is not Interface or Ptr. It returns the zero Value if v is nil.
	// true
	fmt.Printf("a := 1000,  reflect.ValueOf(&a).Elem().CanSet() => %v\n", reflect.ValueOf(&a).Elem().CanSet())
	reflect.ValueOf(&a).Elem().SetInt(511)
	fmt.Printf("after set value, a => %d\n", a) // 511

	fmt.Printf("foo is a struct,  reflect.ValueOf(&a).Elem().CanSet() => %v\n", reflect.ValueOf(&foo).Elem().CanSet())
	// reflect: call of reflect.Value.SetInt on struct Value
	// reflect.ValueOf(&foo).Elem().SetInt(511)
	reflect.ValueOf(&foo).Elem().Field(1).SetInt(78) // ok
	fmt.Printf("after set value, foo => %v\n", foo)  // Name : xyang     Age : 78
	// 尝试通过Field得到每个属性，再取地址，再Elem，最后Set
	// name := reflect.ValueOf(foo).Field(0) // panic
	// reflect.ValueOf(&name).Elem().SetString("Devaid")
	// fmt.Printf("after set value, foo => %v\n", foo) // Name : xyang     Age : 78

	fmt.Println("reflect.ValueOf.Method-------------------------------------------")
	// reflect.ValueOf().MethodByName().Call()的形参是[]reflect.Value

	// ret := reflect.ValueOf(foo).MethodByName("String").Call([]reflect.Value{})
	// fmt.Println(ret)
	// 这样赋值调用会报错，说没有实现对应的接口，暂时还没理解，后面再研究。
	fmt.Println(reflect.ValueOf(foo).MethodByName("String").Call([]reflect.Value{}))

	fmt.Println("反射应用场景-------------------------------------------")

	type Acct struct {
		AcctId   int    `json:"acct_id" bson:"acct_id"`
		AcctName string `json:"acct_name" bson:"acct_name"`
	}

	acct := &Acct{AcctId: 23, AcctName: "Li Jie"}
	// by json
	acctInfo, err := json.Marshal(acct)
	// j1 => [123 34 97 99 99 116 95 105 100 34 58 50 51 44 34 97 99 99 116 95 110 97 109 101 34 58 34 76 105 32 74 105 101 34 125]
	// usage string(j) to print
	fmt.Printf("%v %v\n\n", string(acctInfo), err)

	// get tag of struct by reflect
	// refletc.ValueOf()没有tag属性
	t := reflect.TypeOf(acct)
	field := t.Elem().Field(0)
	fmt.Printf("reflect.TypeOf(acct).Elem().Field(0).Tag.Get(\"json\") = > %v\n", field.Tag.Get("json"))
	fmt.Printf("reflect.TypeOf(acct).Elem().Field(0).Tag.Get(\"bson\") = > %v\n", field.Tag.Get("bson"))

	fmt.Println()

	// conf := &Conf{}
	// then tconf := reflect.TypeOf(conf) will panic: reflect: NumField of non-struct type
	conf := Conf{}
	opName := "Create"
	conf.OpName = &opName
	conf.Port = 13306

	tconf := reflect.TypeOf(conf)
	// 如果 func (conf *Conf) ShowOpName(subname string) string {}，则如果调用
	vconf := reflect.ValueOf(&conf)
	// 如果 func (conf Conf) ShowOpName(subname string) string {}，则如果调用，否则方法调用不到
	// vconf := reflect.ValueOf(conf)
	PrintStructField(&tconf)
	CallMethod(&vconf, "ShowOpName", []interface{}{"test"})
	err = GetTag(&tconf, "OpName", "json")
	err = GetTag(&tconf, "OpName", "xml")
	err = GetTag(&tconf, "DbPort", "json")
	err = GetTag(&tconf, "OpName", "bson")

	// 场景4：用字符串反射实例化结构体
	var regStruct = make(RegStructMap)

	regStruct["conf"] = Conf{}
	regStruct["db"] = DB{}

	InitStructByName("db", regStruct)
	InitStructByName("conf", regStruct)

}
