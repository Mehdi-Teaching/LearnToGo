package main

import (
	"math/rand"
	"testing"
)

// IMPORTANT:
/*
	1. File name must end with [_test]. in this case benchmark_test.go
	2. Benchmark functions must start with [Benchmark]. for example BenchmarkRandInt
	3. Benchmark functions must have [b *testing.B] as parameter.
	4. run tests using command [go test -bench [PACKAGE]]
*/

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}
