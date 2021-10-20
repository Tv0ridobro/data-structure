package util

type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 |
		~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 |
		~float64
}

// Equal returns true if given slices contain same elements without checking capacity
// False otherwise
func Equal[T Comparable](f, s []T) bool {
	if len(f) != len(s) {
		return false
	}
	for i := range f {
		if f[i] != s[i] {
			return false
		}
	}
	return true
}
