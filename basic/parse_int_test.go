package basic

import "testing"

var str = "3276723"

// 综合来看 还是 Atoi 最快
func BenchmarkMethodStrconv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MethodStrconv(str)
	}
}

func BenchmarkMethodParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MethodParse(str)
	}
}

func BenchmarkMethodScanf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MethodScanf(str)
	}
}

func BenchmarkToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToInt(str)
	}
}
