package panel

import (
	"testing"
)

func BenchmarkPaint(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Paint()
	}
}
