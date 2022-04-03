package search

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

const E = 10

var length, target int

func init() {
	rand.Seed(time.Now().UnixNano())

	length = int(math.Pow10(E))
	//target = rand.Intn(length)
	target = length
}

func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := zeroToN(length)
		BinarySearch(l, target)
	}
}

func BenchmarkSimpleSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := zeroToN(length)
		SimpleSearch(l, target)
	}
}

func zeroToN(n int) []int {
	l := make([]int, n+1)
	for i := 0; i <= n; i++ {
		l[i] = i
	}
	return l
}
