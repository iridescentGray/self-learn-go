package main

import "fmt"

func createLazyFunc(lazyVal int) lazyFunc {
	return func() {
		//延迟函数变量
		fmt.Println(lazyVal)
	}
}

type lazyFunc func()

func main() {
	layzSli := make([]lazyFunc, 0)

	for i := 0; i < 10; i++ {
		layzSli = append(layzSli, createLazyFunc(i))
	}

	for _, fc := range layzSli {
		fc()
	}

}
