package main

import "fmt"

func infX1(vv interface{}) (flag bool) {
	_, flag = vv.(float64)
	return
}

func main() {
	var ary1 [5]string

	for i, val := range ary1 {
		fmt.Printf("%d:%s\n", i, val)
	}

	var v interface{} = 4

	v = 3

	_, isInt := v.(int)

	fmt.Println(isInt)

	fmt.Println(infX1(10))
	fmt.Println(infX1("Hello"))
	fmt.Println(infX1(3.14))
}
