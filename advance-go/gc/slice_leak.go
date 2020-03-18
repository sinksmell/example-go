package main

import (
	"fmt"
	"runtime"
	//"time"
)

type P struct {
	Age int `json:"age,omitempty"`
}

func getPartOfSlice() []*P {
	var s = make([]*P, 0, 10000)
	for i := 0; i < 10000; i++ {
		var p = &P{1}
		runtime.SetFinalizer(p, func(x *P) { println("gc happen on p", x) })
		s = append(s, p)
	}
	return s[100:101]
}

func main() {
	var k = getPartOfSlice()

	// type1
	// 1. print
	// 2. gc
	fmt.Println(k[0])
	runtime.GC()

	// type 2
	// 1. gc
	// 2 print
	// runtime.GC()
	// fmt.Println(k[0])
	// time.Sleep(time.Hour)

	// 结论, 当持有某个切片的子切片存在引用,这个切片将不会被释放
	// 例如 这个切片长度为 1000,但是只有一个长度为1的子切片在使用中,剩余的部分得不到释放,从而引起内存泄漏
}
