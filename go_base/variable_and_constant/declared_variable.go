package main

import "fmt"

var go_website = "https://golang.website.org"

const v = 1

func main() {
	//标准变量声明
	fmt.Println("-----var Full------")
	var lang, website string = "Go", "https://golang.org"
	var announceYear int = 2009
	fmt.Println(lang)
	fmt.Println(website)
	fmt.Println(announceYear)
	fmt.Println(go_website)

	//省略
	fmt.Println("-----var omission------")
	var no_tp = 1
	var no_value int
	fmt.Println(no_tp)    //1
	fmt.Println(no_value) //0

	//赋值
	fmt.Println("-----assignment------")
	var y, z float32 = 2, 3
	z, y = y, z
	fmt.Println(y) //3
	fmt.Println(z) //2

	//短变量声明形式
	fmt.Println("-----Short variable------")
	short_lang, short_year := "Go language", 2007
	fmt.Println(short_lang) //Go language
	fmt.Println(short_year) //2007
}
