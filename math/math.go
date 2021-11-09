package math

import (
	"constraints"
	"math"
	"reflect"
)

// NearestPowerOf2 returns nearest power of 2 (k <= (1 << ans) < 2k)
// For non-positive numbers returns 0
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

// Log2 returns floor log2 value of given number
// For non-positive numbers returns -1
func Log2(i int) int {
	if i <= 0 {
		return -1
	}
	ans := 0
	for i != 1 {
		i /= 2
		ans++
	}
	return ans
}

// Min returns min of 2 elements
func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// Max returns max of 2 elements
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// GCD returns gcd of 2 numbers
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func SmallestValue[T constraints.Ordered]() T {
	var a T
	v := reflect.ValueOf(&a).Elem()
	switch v.Kind() {
	case reflect.Int:
		v.SetInt(math.MinInt)
	case reflect.Int8:
		v.SetInt(math.MinInt8)
	case reflect.Int16:
		v.SetInt(math.MinInt16)
	case reflect.Int32:
		v.SetInt(math.MinInt32)
	case reflect.Int64:
		v.SetInt(math.MinInt64)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.SetUint(0)
	case reflect.Float32, reflect.Float64:
		v.SetFloat(math.Inf(-1))
	case reflect.String:
		v.SetString("")
	}
	return a
}

func Comparator[T constraints.Ordered]() func(T, T) int {
	return func(a T, b T) int {
		if a < b {
			return -1
		} else if a == b {
			return 0
		} else {
			return 1
		}
	}
}
