package str

import (
	"github.com/Tv0ridobro/data-structure/util"
)

// ZFunction returns z function of given string
func ZFunction(s string) []int {
	z := make([]int, len(s))
	l, r := 0, 0
	for i := 1; i < len(z); i++ {
		if i <= r {
			z[i] = util.Min(r-i+1, z[i-l])
		}
		for i+z[i] < len(z) && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l, r = i, i+z[i]-1
		}
	}
	return z
}

// PrefixFunction returns prefix function of given string
func PrefixFunction(s string) []int {
	p := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		k := p[i-1]
		for k > 0 && s[i] != s[k] {
			k = p[k-1]
		}
		if s[i] == s[k] {
			k++
		}
		p[i] = k
	}
	return p
}
