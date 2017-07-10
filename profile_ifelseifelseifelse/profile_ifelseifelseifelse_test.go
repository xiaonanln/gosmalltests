package profile_ifelseifelseifelse

import (
	"testing"

	"github.com/xiaonanln/goworld/gwlog"
)

const (
	IFELSECOUNT = 20
)

var (
	r     = IFELSECOUNT - 1
	dummy = 0
)

func BenchmarkIfElseIfElseIfElse(b *testing.B) {

	for i := 0; i < b.N; i++ {
		if r == 0 {
			dummy += 1
		} else if r == 1 {
			dummy += 1
		} else if r == 2 {
			dummy += 1
		} else if r == 3 {
			dummy += 1
		} else if r == 4 {
			dummy += 1
		} else if r == 5 {
			dummy += 1
		} else if r == 6 {
			dummy += 1
		} else if r == 7 {
			dummy += 1
		} else if r == 8 {
			dummy += 1
		} else if r == 9 {
			dummy += 1
		} else if r == 10 {
			dummy += 1
		} else if r == 11 {
			dummy += 1
		} else if r == 12 {
			dummy += 1
		} else if r == 13 {
			dummy += 1
		} else if r == 14 {
			dummy += 1
		} else if r == 15 {
			dummy += 1
		} else if r == 16 {
			dummy += 1
		} else if r == 17 {
			dummy += 1
		} else if r == 18 {
			dummy += 1
		} else if r == 19 {
			dummy += 1
		} else {
			gwlog.Panicf("should not goes here")
		}
	}

}

func BenchmarkDispatchByFuncArray(b *testing.B) {

	dispatchers := []func(){}
	for i := 0; i < IFELSECOUNT; i++ {
		i := i
		dispatchers = append(dispatchers, func() {
			dummy += i
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dispatchers[r]()
	}
}

func BenchmarkDispatchByFuncMap(b *testing.B) {

	dispatchers := map[int]func(){}
	for i := 0; i < IFELSECOUNT; i++ {
		i := i
		dispatchers[i] = func() {
			dummy += i
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dispatchers[r]()
	}
}

func BenchmarkDispatchByDynamicFuncMap(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		map[int]func(){
			0: func() {
				dummy += 1
			},
			1: func() {
				dummy += 1
			},
			2: func() {
				dummy += 1
			},
			3: func() {
				dummy += 1
			},
			4: func() {
				dummy += 1
			},
			5: func() {
				dummy += 1
			},
			6: func() {
				dummy += 1
			},
			7: func() {
				dummy += 1
			},
			8: func() {
				dummy += 1
			},
			9: func() {
				dummy += 1
			},
			10: func() {
				dummy += 1
			},
			11: func() {
				dummy += 1
			},
			12: func() {
				dummy += 1
			},
			13: func() {
				dummy += 1
			},
			14: func() {
				dummy += 1
			},
			15: func() {
				dummy += 1
			},
			16: func() {
				dummy += 1
			},
			17: func() {
				dummy += 1
			},
			18: func() {
				dummy += 1
			},
			19: func() {
				dummy += 1
			},
		}[r]()
	}
}
