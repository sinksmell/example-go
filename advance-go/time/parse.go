package main

import (
	"fmt"
	"time"
)

func main(){
	t,err:=time.Parse("20060102","20200113")
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(t.String())
}


