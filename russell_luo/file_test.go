package russell_luo

import (
	"testing"
)

func BenchmarkAllowRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllowRequest("/benchmark")
	}
}
