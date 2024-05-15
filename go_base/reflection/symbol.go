package main

import (
	"fmt"
	"reflect"
)

type F func(string, int) bool

func (f F) m(s string) bool {
	return f(s, 32)
}
func (f F) M() {}

type I interface {
	m(s string) bool
	M()
}

/*
 */
func main() {
	//获取字段/入参/返回值
	fmt.Println("-----Field------")
	var x struct {
		F F `optional:"yes"`
		i I
	}
	struct_type := reflect.TypeOf(x)
	fmt.Println(struct_type.Kind())        // struct
	fmt.Println(struct_type.NumField())    // 2    返回一个结构体类型的所有字段（包括非导出字段）的数目
	fmt.Println(struct_type.Field(0).Name) // F
	fmt.Println(struct_type.Field(1).Name) // i
	// 包路径（PkgPath）是非导出字段（或者方法）的内在属性
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

	//获取Tag
	fmt.Println("-----Tag------")
	struct_F_tag := struct_type.Field(0).Tag
	fmt.Println(struct_F_tag)                    //optional:"yes"
	fmt.Println(struct_F_tag.Lookup("optional")) //  yes true
	fmt.Println(struct_F_tag.Lookup("test"))     // false

}
