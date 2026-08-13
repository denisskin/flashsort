package flashsort

// Bytes sorts data in ascending lexicographic byte order.
// The sort is performed in place and is not stable.
func Bytes(data [][]byte) {
	n := len(data)
	if n < 2 {
		return
	}
	getByte := func(i int, byteIndex int) (b byte, ok bool) {
		if s := data[i]; byteIndex < len(s) {
			return s[byteIndex], true
		}
		return 0, false
	}
	swap := func(i, j int) {
		data[i], data[j] = data[j], data[i]
	}

	var ok bool
	var b, bp, bHi, bLo int
	const bucketCount = 256
	var pn, p0, p, pm int
	var c, c0 byte

	var bucketVal [256]int
	var bucketLen [256]int
	stackPn := make([]int, 1, 10)
	stackLv := make([]int, 1, 10)
	var iStack, lv int
	stackPn[0] = n

start:
	lv = stackLv[iStack] // level
	pn = stackPn[iStack] // exclusive end of the current partition

	n = pn - p0
	if n < 8 { // use insertion sort for small partitions
		if n == 2 { // sort two elements
			for p = p0 + 1; ; lv++ {
				if c0, ok = getByte(p0, lv); !ok {
					break
				}
				if c, ok = getByte(p, lv); !ok {
					swap(p, p0)
					break
				}
				if c != c0 {
					if c0 > c {
						swap(p, p0)
					}
					break
				}
			}
		} else if n > 2 { // use insertion sort
			for pm = p0 + 1; pm < pn; pm++ {
				for p = pm; p > p0; p-- {
					fSwap := false
					for l := 0; ; l++ {
						if c0, ok = getByte(p-1, l); !ok {
							break
						}
						if c, ok = getByte(p, l); !ok {
							fSwap = true
							break
						}
						if c0 != c {
							fSwap = c0 > c
							break
						}
					}
					if fSwap {
						swap(p, p-1)
					} else {
						break
					}
				}
			}
		}
		// finish the current partition
		p0 = pn
		iStack--
		goto sub
	}

	bHi, bLo = 0, bucketCount

	// skip exhausted keys
	for ; p0 < pn; p0++ {
		if _, ok = getByte(p0, lv); ok {
			break
		}
	}

	// 1. Count bytes in the current position.
	for p = p0; p < pn; p++ {
		if c, ok = getByte(p, lv); ok {
			b = int(c)
			bucketLen[b]++
			if b < bLo {
				bLo = b
			}
			if b > bHi {
				bHi = b
			}
		} else { // move exhausted keys to the beginning
			swap(p, p0)
			p0++
		}
	}

	if bLo == bucketCount { // no keys have another byte
		iStack--
		goto sub
	}

	if bLo == bHi { // all remaining keys have the same byte
		bucketLen[bLo] = 0
		stackLv[iStack]++
		goto start
	}

	// 2. Calculate bucket boundaries.

	// Skip exhausted keys and single-element buckets.
	for b = bLo; b <= bHi && bucketLen[b] < 2; b++ {
		if bucketLen[b] > 0 {
			bucketVal[b] = p0
			p0++
		}
	}

	if p0 != pn {
		// Save the first non-trivial bucket for the next iteration.
		iStack++
		stackPn = append(stackPn[:iStack], p0+bucketLen[b])
		stackLv = append(stackLv[:iStack], lv+1)
	}

	for p = p0; b <= bHi; b++ {
		bucketVal[b] = p
		p += bucketLen[b]
	}

	// 3. Move values into their buckets.
	for b = bLo; b < bHi; b++ {
		for bucketLen[b] > 0 {
			for {
				c, _ = getByte(bucketVal[b], lv)
				if bp = int(c); bp != b {
					swap(bucketVal[b], bucketVal[bp])
					bucketLen[bp]--
					bucketVal[bp]++
				} else {
					break
				}
			}
			bucketLen[b]--
			bucketVal[b]++
		}
	}
	bucketLen[bHi] = 0

	if p0 == pn { // all buckets contain at most one element
		iStack--
		goto sub
	}

	// 4. Process the first non-trivial child bucket.
	goto start

sub:
	// Continue with the remaining child buckets.
	if iStack < 0 {
		// sorting is complete
		return
	}
	pn = stackPn[iStack]
	if p0 >= pn {
		iStack--
		goto sub
	}

	// Find the end of the current child bucket.
	lv = stackLv[iStack]
	c, _ = getByte(p0, lv)
	for p = p0 + 1; p < pn; p++ {
		c0, _ = getByte(p, lv)
		if c != c0 {
			break
		}
		c = c0
	}

	iStack++
	stackLv = append(stackLv[:iStack], lv+1)
	stackPn = append(stackPn[:iStack], p)
	goto start
}
