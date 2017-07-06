package main

import (
	"math/rand"
	"sync"
	"testing"
)

var (
	pool = sync.Pool{
		New: func() interface{} {
			return &Message{}
		},
	}
)

type Message struct {
	data [8192]byte
}

func BenchmarkPoolGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := pool.Get()
		if rand.Float32() < 0.98 {
			pool.Put(v)
		}
	}
}

func BenchmarkMalloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = &Message{}
	}
}
