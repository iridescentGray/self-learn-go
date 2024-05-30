package main

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

/*
context包提供了一个在不同的goroutine之间传递请求范围数据、取消信号和截止时间的方式。
它主要用于管理多个goroutine之间的生命周期和相互通信
*/
func main() {
	//返回一个空的上下文，一般用作根context
	ctx_back := context.Background()
	fmt.Println(ctx_back)

	//用于不知道使用什么context/代码还未完善
	ctx_todo := context.TODO()
	fmt.Println(ctx_todo)

	//创建一个可取消的context和取消函数cancel_func
	ctx_par, cancel_func := context.WithCancel(context.TODO())
	fmt.Println(ctx_par)
	fmt.Println(reflect.TypeOf(cancel_func).Kind()) //func

	//创建一个具有截止时间的context
	ctx_deadline, cancel_deadline := context.WithDeadline(context.TODO(), time.Now().Add(1*time.Hour))
	fmt.Println(ctx_deadline)
	fmt.Println(reflect.TypeOf(cancel_deadline).Kind()) //func

	//创建一个具有截止时间的context
	ctx_timeout, cancel_timeout := context.WithTimeout(context.TODO(), time.Minute)
	fmt.Println(ctx_timeout)
	fmt.Println(reflect.TypeOf(cancel_timeout).Kind()) //func

	//创建一个带有键值对的context，用于传递请求范围内的数据。
	ctx_WithValue := context.WithValue(context.TODO(), "my_key", "inner_value")
	fmt.Println(ctx_WithValue)
	fmt.Println(ctx_WithValue.Value("my_key"))

}
