package flashsort

import (
	"bytes"
	"sort"
	"testing"
)

const benchmarkSize = 10_000

// BenchmarkStdSortStrings measures the standard library string sorter.
func BenchmarkStdSortStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := randStrings(benchmarkSize)
		b.StartTimer()

		sort.Strings(s)
	}
}

// BenchmarkStrings measures Strings.
func BenchmarkStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := randStrings(benchmarkSize)
		b.StartTimer()

		Strings(s)
	}
}

// BenchmarkStdSortBytes measures the standard library byte-slice sorter.
func BenchmarkStdSortBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := randBytes(benchmarkSize)
		b.StartTimer()

		sort.Slice(s, func(i, j int) bool { return bytes.Compare(s[i], s[j]) < 0 })
	}
}

// BenchmarkBytes measures Bytes.
func BenchmarkBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := randBytes(benchmarkSize)
		b.StartTimer()

		Bytes(s)
	}
}
