package main

import "fmt"

type T int

func (t T) M(n int) T {
	fmt.Println(n)
	return t
}

func main() {
	var t T
	// t.M(1)将在M(2)调用被推入延迟调用队列时(now)被估值
	defer t.M(1).M(2)
	t.M(3).M(4)
	//1 3 4 2
}
