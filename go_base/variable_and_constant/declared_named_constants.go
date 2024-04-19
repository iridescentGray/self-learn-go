package main

import "fmt"

const π = 3.1416
const Pi = π
const (
	No         = !Yes
	Yes        = true
	MaxDegrees = 360
	Unit       = "弧度"
)

const X float32 = 3.14
const X_F = float32(3.14)
const A, B = int64(-3), int64(5)

const MaxUint uint = (1 << 64) - 1
const MaxUint2 = ^uint(0)

const MaxInt = int(^uint(0) >> 1)

func main() {
	//类型不明确
	fmt.Println(Yes)
	fmt.Println(Unit)
	fmt.Println(MaxDegrees)

	//使用具名常量
	const DoublePi, HalfPi, Unit2 = π * 2, π * 0.5, "度"
	fmt.Println(DoublePi)
	fmt.Println(HalfPi)
	fmt.Println(Unit2)

	//类型明确
	fmt.Println(X)
	fmt.Println(X_F)
	fmt.Println(A)
	fmt.Println(B)

	//uint的最大值
	fmt.Println(MaxUint)
	fmt.Println(MaxUint2)

	//int 最大值
	fmt.Println(MaxInt)

	//检查当前操作系统是32位的还是64位
	const NativeWordBits = 32 << (^uint(0) >> 63) // 64 or 32
	fmt.Println(NativeWordBits)

	//自动补全
	const (
		Completion_A float32 = 3.14
		Completion_B
		Completion_C, Completion_D = "Go", "language"
		Completion_E, Completion_F
	)
	fmt.Println(Completion_B)
	fmt.Println(Completion_F)

	//自动补全
	const (
		Failed    = iota - 1 // == -1
		Unknown              // == 0
		Succeeded            // == 1
	)
	fmt.Println(Failed)
	fmt.Println(Unknown)
	fmt.Println(Succeeded)

	const (
		Readable   = 1 << iota // == 1
		Writable               // == 2
		Executable             // == 4
	)
	fmt.Println(Readable)
	fmt.Println(Writable)
	fmt.Println(Executable)
}
