# flashsort

`flashsort` is an in-place MSD radix sorter for byte-oriented keys in Go. It
provides specialized functions for strings and byte slices, and a generic
function for sorting any slice through a byte accessor.

## Installation

```bash
go get github.com/denisskin/flashsort
```

## Usage

Sort strings or byte slices in ascending lexicographic order:

```go
package main

import (
	"fmt"

	"github.com/denisskin/flashsort"
)

func main() {
	words := []string{"charlie", "alpha", "bravo"}
	flashsort.Strings(words)
	fmt.Println(words) // [alpha bravo charlie]

	keys := [][]byte{{2}, {0, 1}, {0}, nil}
	flashsort.Bytes(keys)
	fmt.Println(keys) // [[] [0] [0 1] [2]]
}
```

Use `Slice` when the key is part of a larger value:

```go
type Record struct {
	Key  string
	Data any
}

records := []Record{{Key: "b"}, {Key: "a"}}
flashsort.Slice(records, func(i, byteIndex int) (byte, bool) {
	key := records[i].Key
	if byteIndex >= len(key) {
		return 0, false
	}
	return key[byteIndex], true
})
```

The accessor receives the element's current index because the sorter rearranges
the slice in place. It must return `false` after the last byte of the key.

## Behavior

- Sorting is in place and is not stable.
- Ordering is lexicographic by raw bytes; UTF-8 strings therefore use the same
  order as Go's standard string comparison.
- Empty and nil keys sort before non-empty keys.
- `Slice` panics when its first argument is not a slice, matching
  `reflect.Swapper`.

## Benchmarks

The following results were measured with 10,000 deterministic pseudo-random
keys on an Apple M2 using Go 1.26.5:

| Benchmark | Time | Speedup |
| --- | ---: | ---: |
| `sort.Strings` | 879 µs/op | baseline |
| `flashsort.Strings` | 423 µs/op | 2.08× |
| `sort.Slice` with `bytes.Compare` | 1,314 µs/op | baseline |
| `flashsort.Bytes` | 446 µs/op | 2.95× |

Run the benchmarks on your system with:

```bash
go test -run '^$' -bench=. -benchmem -benchtime=1s
```

Results depend on the hardware, Go version, input size, and key distribution.

## Development

```bash
go test ./...
go test -bench=. -benchmem
```

## License

MIT. See [LICENSE](LICENSE).
