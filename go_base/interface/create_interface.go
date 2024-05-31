package main

func main() {
	type error interface {
		Error() string
	}
	type ReadWriteCloser = interface {
		Read(buf []byte) (n int, err error)
		Write(buf []byte) (n int, err error)
		error                      // 一个类型名称
		interface{ Close() error } // 一个类型字面表示形式
	}

	// 此接口类型内嵌了一个近似类型。
	// 它的类型集由所有底层类型为[]byte的类型组成
	type AnyByteSlice = interface {
		~[]byte
	}

	// 此接口类型内嵌了一个类型并集。它的类型集包含6个类型：
	// uint、uint8、uint16、uint32、uint64和uintptr。
	type Unsigned interface {
		uint | uint8 | uint16 | uint32 | uint64 | uintptr
	}

}
