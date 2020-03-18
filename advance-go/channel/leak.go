package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

var (
	res        = []bool{true, false}
	blockedNum = 0
	mu         sync.Mutex
)

func doCount(){
	mu.Lock()
	blockedNum +=1
	fmt.Printf("current blocked num is %d\n",blockedNum)
	mu.Unlock()
}

func main() {
	http.HandleFunc("/ch", func(writer http.ResponseWriter, request *http.Request) {
		MockLeak()
	})
	http.ListenAndServe(":6060", nil)
}

func MockLeak() bool {
	ch := make(chan bool)

	defer close(ch)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				doCount()
			}
		}()
		ch <- doSomething()
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				doCount()
			}
		}()
		ch <- doSomething()
	}()

	return <-ch && <-ch
}

func doSomething() bool {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return res[r.Int()%2]
}
