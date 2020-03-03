package main

import "fmt"

func main() {
	arr := make([]int, 10)
	fmt.Println(arr)
	// [low,high)
	s1:=arr[:len(arr)]
	fmt.Println(s1)
	// panic should  low <= high
	s2:=arr[len(arr)+1:len(arr)]
	fmt.Println(s2)
}
