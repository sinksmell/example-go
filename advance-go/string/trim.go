package main

import (
	"fmt"
	"strings"
)

func main(){
	str:="\thello"
	res:=strings.TrimSpace(str)
	fmt.Println(res)
	fmt.Println(len(res))
}