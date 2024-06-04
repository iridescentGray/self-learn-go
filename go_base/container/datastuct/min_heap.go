package main

import (
	"container/heap"
	"fmt"
)

// 定义一个类型，表示堆中的元素
type IntHeap []int

// Len 方法返回元素数量
func (h IntHeap) Len() int {
	return len(h)
}

// Less 方法比较两个元素的大小
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap 方法交换两个元素的位置
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push 方法向堆中添加元素
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 方法移除并返回最小元素
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/*
heap 包提供了一种灵活方式来实现堆。
通过实现 heap.Interface 接口，可以创建适合自己需求的堆，例如最小堆、最大堆或其他自定义堆结构
可以实现优先队列、调度算法等
*/
func main() {
	// 小顶堆
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	heap.Push(h, 0)

	fmt.Println(h)

	// 移除并打印堆中的最小元素
	fmt.Printf("Pop最小元素: %d\n", heap.Pop(h))
	fmt.Println(h)

}
