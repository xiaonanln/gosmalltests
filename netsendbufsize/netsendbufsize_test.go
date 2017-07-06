package main

import (
	"net"
	"testing"
)

func init() {
	listener := net.Listen
}

func BenchmarkBuffSize4(b *testing.B) {
	benchmarkBuffSize(b, 4)
}

func benchmarkBuffSize(b *testing.B, bufsize int) {
	for i := 0; i < b.N; i++ {
	}
}
