package main

import (
	"context"
	"fmt"
	"time"
)

// 由于 main goroutine 可能会先于 handle goroutine 退出,所以 handle 打印的信息可能会看不到
func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel() // 广播取消事件

	go handle(ctx, time.Second*2)

	// select 等待 ctx timeout
	select {
	case <-ctx.Done():
		fmt.Println("[main] ctx is done ", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	timer := time.NewTimer(duration)
	select {
	case <-ctx.Done():
		fmt.Println("[handle] ctx is done ", ctx.Err())
	case <-timer.C:
		fmt.Println("process request with ", duration)
	}

}
