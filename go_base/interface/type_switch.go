package main

import "fmt"

func main() {
	values := []interface{}{
		456, "abc", true, 0.33, map[int]bool{}, nil,
	}
	for _, x := range values {
		switch v := x.(type) {
		case string: // 一个类型名
			fmt.Println("string:", v)
		case int, float64: // 多个类型名
			fmt.Println("number:", v)
		case nil:
			fmt.Println(v)
		default:
			fmt.Println("others:", v)
		}
	}

}
