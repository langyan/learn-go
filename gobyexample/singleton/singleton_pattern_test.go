package singleton

import "testing"

func BenchmarkSingleton(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetConfigInstance()
	}
}