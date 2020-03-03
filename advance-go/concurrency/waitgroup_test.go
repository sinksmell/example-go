package concurrency

import "testing"

func TestUseMutex(t *testing.T) {
	UseMutex()
}

func TestNotUseMutex(t *testing.T) {
	NotUseMutex()
}

func BenchmarkUseMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseMutex()
	}
}

func BenchmarkNotUseMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NotUseMutex()
	}
}

func BenchmarkAtomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Atomic()
	}
}