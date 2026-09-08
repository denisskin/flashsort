# flashsort

<p align="center">
  <img src="assets/repository-icon.png" alt="flashsort icon" width="160">
</p>

<p align="center">
  <strong>High-performance byte-oriented sorting for Go.</strong><br>
  Sort strings, byte slices, and custom keys up to 3× faster than the standard
  comparison-based approach.
</p>

`flashsort` is an in-place MSD radix sorter built for workloads where sorting
speed matters. Instead of repeatedly comparing entire values, it partitions
keys byte by byte, making it especially effective for large collections of
strings and binary keys.

## Performance

On the included benchmark, `flashsort` sorts strings **2.08× faster** and byte
slices **2.95× faster** than the equivalent standard-library approach.

Results for 10,000 deterministic pseudo-random keys on an Apple M2 with
Go 1.26.5:

| Data | Standard library | `flashsort` | Speedup |
| --- | ---: | ---: | ---: |
| Strings | 879 µs/op | **423 µs/op** | **2.08×** |
| Byte slices | 1,314 µs/op | **446 µs/op** | **2.95×** |

The `flashsort` implementations reported **0 heap allocations per operation**
in this benchmark.

Run it on your machine:

```bash
go test -run '^$' -bench=. -benchmem -benchtime=1s
```

Benchmark results depend on the hardware, Go version, input size, key length,
and key distribution.

## Highlights

- **Strong benchmark performance** — 2.08× faster for strings and 2.95× faster
  for byte slices in the benchmark above.
- **Byte-oriented sorting** — works naturally with strings, identifiers,
  hashes, serialized keys, paths, and protocol data.
- **In-place operation** — values are rearranged directly in the input slice.
- **Low allocation overhead** — the specialized string and byte-slice
  functions complete the included benchmark without heap allocations.
- **Custom-key support** — structs and domain types can be sorted without
  building a separate key slice.
- **No external dependencies** — the package relies only on the Go standard
  library.

## Installation

```bash
go get github.com/denisskin/flashsort
```

## Usage

### Strings and byte slices

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

### UTF-8 strings

`Strings` works with UTF-8 text without any special configuration:

```go
names := []string{"张伟", "Денис", "Alice", "Ζήνων", "Борис"}

flashsort.Strings(names)

fmt.Println(names) // [Alice Ζήνων Борис Денис 张伟]
```

The result follows Go's lexicographic string order. It is suitable for
deterministic sorting of UTF-8 keys, but it is not locale-aware alphabetical
collation.

### Custom keys

Use `Slice` when the sort key is part of a larger value:

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
the slice in place. Return `false` when `byteIndex` reaches the end of the key.

## Details

- Values are sorted in ascending lexicographic byte order.
- UTF-8 strings follow the same ordering as Go's standard string comparison.
- Empty and nil keys are placed before non-empty keys.
- Sorting is performed directly on the supplied slice.
- `Slice` panics when its first argument is not a slice, matching
  `reflect.Swapper`.

## Development

```bash
go test ./...
go test -bench=. -benchmem
```

## License

MIT. See [LICENSE](LICENSE).
