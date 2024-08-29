package main

import (
	"fmt"

	"github.com/duke-git/lancet/v2/strutil"
)

func main() {
	fmt.Println(strutil.Wrap("foo", ""))    //foo
	fmt.Println(strutil.Wrap("foo", "*"))   //*foo*
	fmt.Println(strutil.Wrap("'foo'", "'")) //''foo''
	fmt.Println(strutil.Wrap("", "*"))      //

	fmt.Println(strutil.HasSuffixAny("foo bar", []string{"bar", "xyz", "hello"})) //true
	fmt.Println(strutil.HasSuffixAny("foo bar", []string{"oom", "world"}))        //false

}
