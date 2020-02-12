package utils

import (
	"testing"
)

// 性能测试
func BenchmarkRandStr8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStr(8)
	}
}

func BenchmarkRandStr10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStr(10)
	}
}
func BenchmarkRandStr16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStr(16)
	}
}
