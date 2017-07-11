package profile_getptrbymemberptr

import (
	"testing"
	"unsafe"
)

type T struct {
	a int
	b int
	m M
	c [4096]byte
}

type M struct {
	a int
	b int
}

func (m *M) T() *T {
	dummy := (*T)(unsafe.Pointer(m))
	return (*T)(unsafe.Pointer((uintptr)(unsafe.Pointer(m)) - (uintptr(unsafe.Pointer(&dummy.m)) - uintptr(unsafe.Pointer(dummy)))))
}

func (m *M) T_Bad() *T {
	return (*T)(unsafe.Pointer(uintptr(unsafe.Pointer(m)) - 16))
}

func BenchmarkGetPtrByMemberPtr(b *testing.B) {
	var t T
	for i := 0; i < b.N; i++ {
		if &t != t.m.T() {
			b.Fatal("wrong")
		}
	}
}

func BenchmarkGetPtrByMemberPtr_FastButWrong(b *testing.B) {
	var t T
	for i := 0; i < b.N; i++ {
		if &t != t.m.T_Bad() {
			b.Fatal("wrong")
		}
	}
}
