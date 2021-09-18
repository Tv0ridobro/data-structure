package util

func NearestPowerOf2(k int) int {
	if (k & (k - 1)) == 0 {
		return k
	} else {
		k |= k >> 1
		k |= k >> 2
		k |= k >> 4
		k |= k >> 8
		k |= k >> 16
		return (k - (k >> 1)) << 1
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
