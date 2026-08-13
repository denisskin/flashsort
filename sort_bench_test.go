package flashsort

import (
	"bytes"
	"sort"
	"testing"
)

const N = 1e6

func BenchmarkStdSortStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := randStrings(N)
		b.StartTimer()

		sort.Strings(s)
	}
}

func BenchmarkStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := randStrings(N)
		b.StartTimer()

		Strings(s)
	}
}

func BenchmarkStdSortBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := randBytes(N)
		b.StartTimer()

		sort.Slice(s, func(i, j int) bool { return bytes.Compare(s[i], s[j]) == -1 })
	}
}

func BenchmarkBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := randBytes(N)
		b.StartTimer()

		Bytes(s)
	}
}
