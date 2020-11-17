package main

import "fmt"

func sum(s ...int) (result int) {
	for _, v := range s {
		result = result + v
	}
	return
}

func map1(s []int) {
	for i, v := range s {
		s[i] = v * v
	}
}

func main() {
	var x []int

	x = make([]int, 10)

	x[1] = 10

	fmt.Println(x)
	fmt.Println(len(x))

	var x2 []int

	x2 = make([]int, 20)

	copy(x2, x)

	fmt.Println(x2)
	fmt.Println(len(x2))
	fmt.Println(cap(x2))

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := a[2:4]
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	s2 := a[2:4:4]
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	s3 := a[2:4:6]
	fmt.Println(len(s3))
	fmt.Println(cap(s3))

	for _, v := range a {
		fmt.Println(v + 1)
	}

	fmt.Println(sum(1, 3, 5, 7, 9))

	z := []int{1, 3, 5}
	fmt.Printf("org:%v\n", z)
	map1(z)
	fmt.Printf("now:%v\n", z)

	m1 := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}
	fmt.Println(m1)
	fmt.Println(len(m1))

	if _, f := m1[1]; f {
		fmt.Println("OK")
	}

	delete(m1, 3)

	fmt.Println(m1)

	var c1 chan int
	c1 = make(chan int, 5) // サイズ５のチャネル
	c1 <- 5
	i := <-c1
	fmt.Println(i)
}
