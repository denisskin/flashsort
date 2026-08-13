package flashsort

import (
	"bytes"
	"math/rand"
	"reflect"
	"slices"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestStrings(t *testing.T) {
	ss := []string{
		"Hunter",
		"Isaac",
		"Сэр",
		"Исак",
		"Ньютон",
		"Christopher",
		"Кристофер",
		"Нолан",
		"Bob",
		"Боб",
		"Faith",
		"Alice",
		"Алиса",
		"Gabriel",
		"Пётр",
		"Denis",
		"Денис",
		"Kitty",
		"Ethan",
	}

	Strings(ss)

	require(t, slices.IsSorted(ss))
}

func TestStringsRandom(t *testing.T) {
	a := randStrings(testSize)

	Strings(a)

	require(t, slices.IsSorted(a), "not sorted")
}

func TestBytes(t *testing.T) {
	a := randBytes(testSize)
	b := randBytes(testSize)

	Bytes(a)
	sort.Slice(b, func(i, j int) bool { return bytes.Compare(b[i], b[j]) < 0 })

	assertEqual(t, a, b)
}

func TestBytesWithNullBytes(t *testing.T) {
	data := [][]byte{
		{0, 9}, {0, 8}, {0, 7}, {0, 6}, {0, 5},
		{0, 4}, {0, 3}, {0, 2}, {0, 1}, {0}, nil,
	}
	want := slices.Clone(data)
	slices.SortFunc(want, bytes.Compare)

	Bytes(data)

	assertEqual(t, data, want)
}

func TestBytesWithDistinctLeadingBytes(t *testing.T) {
	data := make([][]byte, 256)
	for i := range data {
		data[i] = []byte{byte(255 - i)}
	}
	want := slices.Clone(data)
	slices.SortFunc(want, bytes.Compare)

	Bytes(data)

	assertEqual(t, data, want)
}

func TestStringsWithLongCommonPrefix(t *testing.T) {
	prefix := "a-very-long-common-prefix-"
	data := make([]string, 300)
	for i := range data {
		data[i] = prefix + strconv.Itoa(len(data)-i)
	}
	want := slices.Clone(data)
	slices.Sort(want)

	Strings(data)

	assertEqual(t, data, want)
}

func TestStringsWithDeeplyNestedBuckets(t *testing.T) {
	data := make([]string, 0, 24)
	for depth := 0; depth < 20; depth++ {
		data = append(data, strings.Repeat("a", depth)+"b")
	}
	data = append(data, strings.Repeat("a", 20)+"a", strings.Repeat("a", 20)+"c")
	slices.Reverse(data)
	want := slices.Clone(data)
	slices.Sort(want)

	Strings(data)

	assertEqual(t, data, want)
}

func TestSlice(t *testing.T) {
	type record struct {
		key string
		id  int
	}
	data := []record{{"charlie", 3}, {"alpha", 1}, {"bravo", 2}}

	Slice(data, func(i, byteIndex int) (byte, bool) {
		if byteIndex >= len(data[i].key) {
			return 0, false
		}
		return data[i].key[byteIndex], true
	})

	require(t, slices.IsSortedFunc(data, func(a, b record) int {
		return strings.Compare(a.key, b.key)
	}))
}

func TestSliceWithNullBytesAndEmptyKeys(t *testing.T) {
	type record struct {
		key []byte
		id  int
	}
	data := []record{
		{[]byte{0, 9}, 9}, {[]byte{0, 8}, 8}, {[]byte{0, 7}, 7},
		{[]byte{0, 6}, 6}, {[]byte{0, 5}, 5}, {[]byte{0, 4}, 4},
		{[]byte{0, 3}, 3}, {[]byte{0, 2}, 2}, {[]byte{0, 1}, 1},
		{[]byte{0}, 0}, {nil, -1},
	}
	want := slices.Clone(data)
	slices.SortFunc(want, func(a, b record) int {
		return bytes.Compare(a.key, b.key)
	})

	Slice(data, func(i, byteIndex int) (byte, bool) {
		if byteIndex >= len(data[i].key) {
			return 0, false
		}
		return data[i].key[byteIndex], true
	})

	assertEqual(t, data, want)
}

const testSize = 100_000

// randStrings returns deterministic pseudo-random hexadecimal strings.
func randStrings(n int) []string {
	r := rand.New(rand.NewSource(int64(n)))
	s := make([]string, n)
	for i := range s {
		s[i] = strconv.FormatUint(r.Uint64(), 16)
	}
	return s
}

// randBytes returns deterministic pseudo-random byte slices.
func randBytes(n int) [][]byte {
	r := rand.New(rand.NewSource(int64(n)))
	s := make([][]byte, n)
	for i := range s {
		s[i] = []byte(strconv.FormatUint(r.Uint64(), 16))
	}
	return s
}

// require fails the test when ok is false.
func require(t *testing.T, ok bool, msg ...any) {
	t.Helper()
	if !ok {
		t.Fatal(msg...)
	}
}

// assertEqual fails the test when a and b are not deeply equal.
func assertEqual(t *testing.T, a, b any) {
	t.Helper()
	require(t, reflect.DeepEqual(a, b), "not equal")
}
