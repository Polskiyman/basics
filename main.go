package main

import (
	"fmt"
)

func main() {
	var a stack
	a.arr = append(a.arr, 2, 8, 1, 9, 3, 7, 4, 6, 5, 4)
	fmt.Println("__________________________________bubbleSort_____________________________________")
	fmt.Println("before ", a)
	bubbleSort(a.arr)
	fmt.Println("after ", a)

	fmt.Println("__________________________________Stack_____________________________________")

	a.push(22)
	fmt.Println("push 22 ", a)

	err := a.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("pop ", a)

	peek, err := a.Peek()
	if err != nil {
		fmt.Println("peek ", err)
	}
	fmt.Println(peek)
}

func bubbleSort(a []int) {
	len := len(a)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			one := a[j]
			two := a[j+1]
			if one > two {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

type stack struct {
	arr []int
}

func (s *stack) push(v int) {
	s.arr = append(s.arr, v)
}

func (s *stack) Pop() error {
	len := len(s.arr)
	if len == 0 {
		return fmt.Errorf("Pop Error: Stack is empty")
	}
	s.arr = s.arr[:len-1]
	return nil
}

func (s *stack) Peek() (int, error) {
	len := len(s.arr)
	if len == 0 {
		return 0, fmt.Errorf("Pop Error: Stack is empty")
	}

	return s.arr[len-1], nil
}
