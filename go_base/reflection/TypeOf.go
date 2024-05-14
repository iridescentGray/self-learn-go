package main

import (
	"fmt"
	"reflect"
)

/*
reflect.TypeOf
从一个任何非接口类型的值创建一个reflect.Type值，此reflect.Type值表示着此非接口值的类型

reflect.Type是接口类型
它的方法能够查看到reflect.Type值所表示的Go类型的各种信息

# Elem用于获取元素类型/指针类型的基类型
*/
type MyT []interface{ m() }

func (MyT) m() {}

type F func(string, int) bool

func (f F) m(s string) bool {
	return f(s, 32)
}
func (f F) M() {}

type I interface {
	m(s string) bool
	M()
}

func main() {
	//基础api
	type A = [16]int16
	var a <-chan map[A][]byte
	fmt.Println(reflect.TypeOf(a).Kind()) // map[[16]int16][]uint8
	fmt.Println(reflect.TypeOf(a).Elem()) // chan

	fmt.Println(reflect.TypeOf(a).ChanDir())            // <-chan
	fmt.Println(reflect.TypeOf(a).Elem().Kind())        // map
	fmt.Println(reflect.TypeOf(a).Elem().Key())         // [16]int16
	fmt.Println(reflect.TypeOf(a).Elem().Key().Kind())  // array
	fmt.Println(reflect.TypeOf(a).Elem().Elem().Kind()) //slice

	// 判断可比较
	fmt.Println("-----Comparable------")
	type C = map[int]int
	var b int = 1
	c := C{}
	fmt.Println(reflect.TypeOf(b).Comparable()) // true
	fmt.Println(reflect.TypeOf(c).Comparable()) // false  map不可比较

	//判断可转换
	fmt.Println("-----ConvertibleTo------")
	var d rune = 1
	var e int32 = 1
	var f float64 = 1

	fmt.Println(reflect.TypeOf(d).ConvertibleTo(reflect.TypeOf(e))) // true
	fmt.Println(reflect.TypeOf(e).ConvertibleTo(reflect.TypeOf(f))) // true

	//判断实现
	fmt.Println("-----Implements------")
	Myt_type := reflect.TypeOf(MyT{})
	empty_interface_type := reflect.TypeOf(new(interface{}))

	fmt.Println(Myt_type)                    //main.MyT
	fmt.Println(empty_interface_type)        //*interface {}
	fmt.Println(Myt_type.Kind())             //slice
	fmt.Println(empty_interface_type.Kind()) // ptr
	fmt.Println(Myt_type.Elem())             // interface { main.m() }
	fmt.Println(empty_interface_type.Elem()) // interface {}

	fmt.Println(empty_interface_type.Implements(empty_interface_type.Elem()))        // true
	fmt.Println(empty_interface_type.Elem().Implements(empty_interface_type.Elem())) // true
	fmt.Println(Myt_type.Implements(empty_interface_type.Elem()))                    // true
	fmt.Println(empty_interface_type.Implements(Myt_type.Elem()))                    // false

	//获取字段/入参/返回值
	fmt.Println("-----Field------")
	var x struct {
		F F
		i I
	}
	struct_type := reflect.TypeOf(x)
	fmt.Println(struct_type.Kind())        // struct
	fmt.Println(struct_type.NumField())    // 2    返回一个结构体类型的所有字段（包括非导出字段）的数目
	fmt.Println(struct_type.Field(0).Name) // F
	fmt.Println(struct_type.Field(1).Name) // i
	//      包路径（PkgPath）是非导出字段（或者方法）的内在属性
	fmt.Println(struct_type.Field(0).PkgPath) //  空
	fmt.Println(struct_type.Field(1).PkgPath) // main

	fmt.Println("-----Method------")
	struct_F_type, struct_i_type := struct_type.Field(0).Type, struct_type.Field(1).Type
	fmt.Println(struct_F_type.Kind())                                 // func
	fmt.Println(struct_F_type.NumMethod(), struct_i_type.NumMethod()) // 1 2   返回一个类型的所有导出的方法
	fmt.Println(struct_F_type.Method(0).Name)                         // M
	fmt.Println(struct_i_type.Method(1).Name)                         // m

	fmt.Println(struct_F_type.MethodByName("m")) //{...}  false   不能用来获取一个类型的非导出方法
	fmt.Println(struct_i_type.MethodByName("m")) //{...}  true

	fmt.Println("-----Entering/Return value------")
	fmt.Println(struct_F_type.IsVariadic())                    // false 是否为变长参数
	fmt.Println(struct_F_type.NumIn(), struct_F_type.NumOut()) // 2 1   入参/返回值数量
	t0, t1, t2 := struct_F_type.In(0), struct_F_type.In(1), struct_F_type.Out(0)
	fmt.Println(t0.Kind(), t1.Kind(), t2.Kind()) // 下一行打印出：string int bool

}
