package main

import "fmt"

type counter interface {
	GetCount()int
}

type user struct {
	count int
}

func(u*user)GetCount()int{
	return u.count
}

func p(cnts ...counter){
	for _, cnt := range cnts {
		fmt.Println(cnt)
	}
}



func main(){
	users:=[]*user{
		{2},
		{3},
		{4},
	}
	p(users[0]...)
}
