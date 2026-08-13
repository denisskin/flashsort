package flashsort

func Strings(data []string) {
	n := len(data)
	if n < 2 {
		return
	}
	getByte := func(i int, iByte int) (b byte, ok bool) {
		if s := data[i]; iByte < len(s) {
			return s[iByte], true
		}
		return 0, false
	}
	swap := func(i, j int) {
		data[i], data[j] = data[j], data[i]
	}

	var ok bool
	var b, bp, bHi, bLo int
	var bMax = 256
	var pn, p0, p, pm int
	var c, c0 byte

	var bucketVal [256]int
	var bucketLen [256]int
	var stackPn [stackSize]int
	var stackLv [stackSize]int
	var iStack, lv int

	stackPn[0] = n

start:
	lv = stackLv[iStack] // level
	pn = stackPn[iStack] // pointer to the end of array

	n = pn - p0
	if n < 8 { // use trivial sort
		if n == 2 { // sort 2 element
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
		} else if n > 2 { // use bubble sort
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
		// break iteration
		p0 = pn
		iStack--
		goto sub
	}

	bHi, bLo = 0, bMax

	// skip NULL values
	for ; p0 < pn; p0++ {
		if _, ok = getByte(p0, lv); ok {
			break
		}
	}

	// 1. calc buckets sizes
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
		} else { // move NULL value to the beginning.  p0++
			swap(p, p0)
			p0++
		}
	}

	if bHi == 0 { // if countBuckets == 0
		iStack--
		goto sub
	}

	if bLo == bHi { // if countBuckets == 1
		bucketLen[bLo] = 0
		stackLv[iStack]++
		goto start
	}

	// 2. set scopes of buckets (pointers to value)

	// skip NULL values and values with length = 1
	for b = bLo; bucketLen[b] < 2 && b <= bHi; b++ {
		if bucketLen[b] > 0 {
			bucketVal[b] = p0
			p0++
		}
	}

	if p0 != pn {
		// pre-set size of first sub.bucket for next iteration
		iStack++
		stackPn[iStack] = p0 + bucketLen[b]
		stackLv[iStack] = lv + 1
	}

	for p = p0; b <= bHi; b++ {
		bucketVal[b] = p
		p += bucketLen[b]
	}

	// 3. move values into appropriate buckets
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

	if p0 == pn { // all buckets has length == 1. rise to up
		iStack--
		goto sub
	}

	// 4. start processing of sub buckets.  process first sub bucket
	goto start

sub:
	// continue processing sub buckets
	if iStack < 0 {
		// finish sorting
		return
	}
	pn = stackPn[iStack]
	if p0 >= pn {
		iStack--
		goto sub
	}

	// calculate length of current sub-bucket
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
	stackLv[iStack] = lv + 1
	stackPn[iStack] = p
	goto start
}
