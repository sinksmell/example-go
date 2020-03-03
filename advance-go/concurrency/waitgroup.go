package concurrency

import (
	"sync"
	"sync/atomic"
)

// 测试不使用 mutex 控制并发会如何

var total struct {
	sync.Mutex
	value int64
}

func UseMutex() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
}

func NotUseMutex() {
	var wg sync.WaitGroup
	wg.Add(2)
	go noMutexWorker(&wg)
	go noMutexWorker(&wg)
	wg.Wait()
}

// 使用原子操作代替锁来保证原子性
func Atomic() {
	var wg sync.WaitGroup
	wg.Add(2)
	go atomicWorker(&wg)
	go atomicWorker(&wg)
	wg.Wait()
}

// 原子操作
func atomicWorker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		atomic.AddInt64(&total.value, int64(i))
	}
}

// 不使用mutex保证 total.value += i 的原子行
func noMutexWorker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		total.value += int64(i)
	}
}

// 使用 mutex 来管理临界区
func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		total.Lock()
		total.value += int64(i)
		total.Unlock()
	}
}
