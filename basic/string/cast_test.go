package string

import "testing"

var str = "hello,world"
var bytes =[]byte("hello,world")

func BenchmarkStringToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringToBytes(str)
	}
}

func BenchmarkStringToBytes2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringToBytes2(str)
	}
}


func BenchmarkBytesToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BytesToString(bytes)
	}
}

func BenchmarkBytesToString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BytesToString2(bytes)
	}
}