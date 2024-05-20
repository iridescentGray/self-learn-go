package main

import (
	"fmt"
	"reflect"
)

/*
reflect.ValueOf

从一个非接口类型的值创建一个reflect.Value值,reflect.Value值代表着此非接口值

reflect.Value有很多观察/操纵属主值的方法 (通过不合适方法调用会导致恐慌)
结构体值的非导出字段不能通过反射来修改
*/
func main() {
	n := 123
	p := &n
	fmt.Println("-----base api------")
	valueof_pointer := reflect.ValueOf(p)
	//Value不可修改 不可寻址
	fmt.Println(valueof_pointer.CanSet(), valueof_pointer.CanAddr()) // false false
	valueof_pointer_elem := valueof_pointer.Elem()                   // 取得valueof_pointer的底层指针值引用的值的代表值

	fmt.Println(reflect.ValueOf(valueof_pointer_elem))                         //<int Value>
	fmt.Println(valueof_pointer_elem.CanSet(), valueof_pointer_elem.CanAddr()) // true true
	valueof_pointer_elem.Set(reflect.ValueOf(789))
	fmt.Println(n) // 789

	fmt.Println("-----Non exported fields cannot be set ------")
	var s struct {
		X interface{} // 导出字段
		y interface{} // 非导出字段
	}
	struct_pointer_elem := reflect.ValueOf(&s).Elem()
	fmt.Println(struct_pointer_elem) // {<nil> <nil>}
	struct_pointer_x := struct_pointer_elem.Field(0)
	struct_pointer_y := struct_pointer_elem.Field(1)
	fmt.Println(struct_pointer_x.CanSet(), struct_pointer_x.CanAddr()) // true true
	// 非导出字段不可set
	fmt.Println(struct_pointer_y.CanSet(), struct_pointer_y.CanAddr()) // false true

}
