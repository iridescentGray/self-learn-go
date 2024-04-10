package main

import "fmt"

func main() {
	//整数字面量
	fmt.Println(15 == 017) // true

	//下划线字面量
	fmt.Println(6_9 == 69)                  //true
	fmt.Println(0_33_77_22 == 0337722)      //true
	fmt.Println(0x_Bad_Face == 0xBadFace)   //true
	fmt.Println(0x_1F_FFp-16 == 0x1FFFp-16) //true

	//rune字面量
	fmt.Println('a' == 97)           //true
	fmt.Println('a' == '\141')       //true
	fmt.Println('a' == '\x61')       //true
	fmt.Println('a' == '\u0061')     //true
	fmt.Println('a' == '\U00000061') //true
	fmt.Println(0x61 == '\x61')      //true
	fmt.Println('\u4f17' == '众')     //true
}
