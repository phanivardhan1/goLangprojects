package main

import (
	"fmt"
)

type Student struct {
	Name string
	Id   int
	Age  int
}

func main() {

	s1 := Student{
		Name: "phani",
		Id:   111,
		Age:  21,
	}

	s2 := Student{
		Name: "phani",
		Id:   121,
		Age:  24,
	}

	s3 := Student{
		Name: "phani",
		Id:   311,
		Age:  22,
	}

	arr := [3]Student{s1, s2, s3}

	s4 := arr[1]
	s5 := &s4
	fmt.Println(s5)

	a := make([]int, 5, 10)

	fmt.Println(a)

}
