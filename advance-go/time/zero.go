package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		t1  time.Time
		now = time.Now()
	)
	fmt.Println(t1.IsZero())
	fmt.Println(now.IsZero())
}
