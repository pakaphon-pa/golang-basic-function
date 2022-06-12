package concurrency

import "testing"

func BenchmarkSequentialProcessing(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ConcurrencyProcessing()
	}
}
