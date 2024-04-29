package main

import "fmt"

type Age int

func (age Age) LargerThan(a Age) bool {
	return age > a
}

// 为自定义的函数类型FilterFunc声明方法。
type FilterFunc func(in int) bool

func (ff FilterFunc) Filte(in int) bool {
	return ff(in)
}

// 为自定义的映射类型StringSet声明方法。
type StringSet map[string]struct{}

func (ss StringSet) Has(key string) bool {
	_, present := ss[key]
	return present
}
func (ss StringSet) Add(key string) {
	ss[key] = struct{}{}
}

type Book struct {
	pages int
}

func (b Book) Pages() int {
	return b.pages
}
func (b *Book) SetPages(pages int) {
	b.pages = pages
}

func main() {
	var hong Age = 1
	var ming Age = 2
	fmt.Println(hong)                  //1
	fmt.Println(ming)                  //2
	fmt.Println(hong.LargerThan(ming)) //false

	var testF FilterFunc = func(in int) bool {
		return in%2 == 0 // 判断是否为偶数
	}
	fmt.Println(FilterFunc.Filte(testF, 1)) //false

	var strMap StringSet = StringSet{}
	fmt.Println(strMap.Has("123")) //false
	strMap.Add("123")
	fmt.Println(strMap.Has("123")) //true

	var mybook Book
	fmt.Println(mybook.Pages()) //0 零值
	mybook.SetPages(1)
	fmt.Println(mybook.Pages()) //1
	(&mybook).SetPages(2)
	fmt.Println(mybook.Pages()) //2

}
