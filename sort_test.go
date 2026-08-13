package flashsort

import (
	"bytes"
	"math/rand"
	"reflect"
	"slices"
	"sort"
	"strconv"
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

func TestStrings_ex(t *testing.T) {

	a := randStrings(N)

	Strings(a)

	require(t, slices.IsSorted(a), "not sorted")
}

func TestBytes(t *testing.T) {

	a := randBytes(N)
	b := randBytes(N)

	Bytes(a)
	sort.Slice(b, func(i, j int) bool { return bytes.Compare(b[i], b[j]) < 0 })

	assertEqual(t, a, b)
}

// --------------------------------------
func randStrings(n int) []string {
	r := rand.New(rand.NewSource(int64(n)))
	s := make([]string, n)
	for i := range s {
		s[i] = strconv.FormatUint(r.Uint64(), 16)
	}
	return s
}

func randBytes(n int) [][]byte {
	r := rand.New(rand.NewSource(int64(n)))
	s := make([][]byte, n)
	for i := range s {
		s[i] = []byte(strconv.FormatUint(r.Uint64(), 16))
	}
	return s
}

func require(t *testing.T, ok bool, msg ...any) {
	if !ok {
		t.Fatal(msg...)
	}
}

func assertEqual(t *testing.T, a, b any) {
	require(t, reflect.DeepEqual(a, b), "not equal")
}
