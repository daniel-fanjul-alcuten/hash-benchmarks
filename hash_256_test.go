package benchmarks

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestNothing(t *testing.T) {
}

func BenchmarkWriter(b *testing.B) {
	buffer, hasher := make([]byte, 1024), sha256.New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hasher.Write(buffer)
	}
}

func BenchmarkToString(b *testing.B) {
	buffer, hasher := make([]byte, 1024), sha256.New()
	hasher.Write(buffer)
	hash := hasher.Sum(nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hex.EncodeToString(hash)
	}
}
