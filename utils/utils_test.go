package utils

import (
	cryptorand "crypto/rand"
	"math/rand"
	"testing"
)

// go test -benchmem -run=^$  -bench ^BenchmarkRandxUint32  -v -count=3 ./utils

func BenchmarkRandxUint32(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_ = Uint32()
		}
	})
}

func BenchmarkRandxUint32Read(b *testing.B) {
	// b.RunParallel(func(p *testing.PB) {
	// 	for p.Next() {
	for i := 0; i < b.N; i++ {
		p := make([]byte, 16)
		_ = Read(p)
	}
	// })
}

func BenchmarkRandxUint32ReadSingle(b *testing.B) {
	// b.RunParallel(func(p *testing.PB) {
	// 	for p.Next() {
	for i := 0; i < b.N; i++ {
		p := make([]byte, 16)
		_ = ReadSingle(p)
	}
	// })
}

func BenchmarkMathUint32Read(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			p := make([]byte, 10)
			rand.Read(p)
		}
	})
}

func BenchmarkCryptoUint32Read(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			p := make([]byte, 10)
			cryptorand.Read(p)
		}
	})
}
