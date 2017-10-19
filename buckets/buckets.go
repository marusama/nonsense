package buckets

import (
	"fmt"
	"sort"
	"strings"
)

type s struct {
	v uint32
	n uint32
}

func (x s) String() string {
	b := make([]byte, x.n)
	for i := 0; i < int(x.n); i++ {
		b[i] = byte(((x.v & (1 << uint(i))) >> uint(i)) + 49)
	}
	return string(b)
}

func fib(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func Print(n uint32, print bool) {
	if n > 32 {
		fmt.Println("overflow")
		return
	}
	x1 := s{v: 0, n: n}
	m := make(map[s]struct{}, fib(int(n)))
	f(x1, m)

	if print {
		result := make([]string, 0, len(m))
		for k := range m {
			result = append(result, k.String())
		}
		sort.Slice(result, func(i, j int) bool {
			l, r := result[i], result[j]
			if len(l) == len(r) {
				return strings.Compare(l, r) > 0
			}
			return len(l) > len(r)
		})

		fmt.Println(fmt.Sprintf("result for %v:", n))
		for _, s := range result {
			fmt.Println(s)
		}
	}
	fmt.Println("len =", len(m))
}

func f(x s, m map[s]struct{}) {
	m[x] = struct{}{}
	l := x.v & 1
	for i := uint32(1); i < x.n; i++ {
		c := (x.v & (1 << i)) >> i
		if c == 0 && l == 0 {
			vv := ((x.v >> i) << (i - 1)) | (1 << (i - 1)) | (x.v & ((1 << (i - 1)) - 1))
			y := s{v: vv, n: x.n - 1}
			if _, ok := m[y]; !ok {
				f(y, m)
			}
		}
		l = c
	}
}
