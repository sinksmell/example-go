package main

import "testing"

func BenchmarkConversion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Conversion()
	}
}

func BenchmarkConversion2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Conversion2()
	}
}