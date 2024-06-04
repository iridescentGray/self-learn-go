package main

import (
	"fmt"
)

func main() {

	fmt.Println("----------------slice---------------")
	var heads_array = []*[4]byte{
		&[4]byte{'P', 'N', 'G', ' '},
		&[4]byte{'G', 'I', 'F', ' '},
		&[4]byte{'J', 'P', 'E', 'G'},
	}
	fmt.Println(heads_array)

	var heads_array_simple = []*[4]byte{
		{'P', 'N', 'G', ' '},
		{'G', 'I', 'F', ' '},
		{'J', 'P', 'E', 'G'},
	}
	fmt.Println(heads_array_simple)

	fmt.Println("----------------array---------------")
	type language struct {
		name string
		year int
	}

	var struct_array = [...]language{
		language{"C", 1972},
		language{"Python", 1991},
		language{"Go", 2009},
	}
	fmt.Println(struct_array)

	var struct_array_simple = [...]language{
		{"C", 1972},
		{"Python", 1991},
		{"Go", 2009},
	}
	fmt.Println(struct_array_simple)

	fmt.Println("----------------map---------------")
	var map_array = map[string]map[string]int{
		"k1": map[string]int{"C": 1972},
		"k2": map[string]int{"Python": 1991},
		"k3": map[string]int{"Go": 2009},
	}
	fmt.Println(map_array)

	var map_array_simple = map[string]map[string]int{
		"k1": {"C": 1972},
		"k2": {"Python": 1991},
		"k3": {"Go": 2009},
	}
	fmt.Println(map_array_simple)

}
